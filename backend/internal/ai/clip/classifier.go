package clip

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"
	"os"
	"sort"

	ort "github.com/yalue/onnxruntime_go"
	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
)

const (
	// CLIP image encoder input size.
	clipImageSize = 224

	// Embedding dimension produced by CLIP.
	embedDim = 512
)

var (
	// CLIP normalization values used during image preprocessing.
	mean = [3]float32{0.48145466, 0.4578275, 0.40821073}
	std  = [3]float32{0.26862954, 0.26130258, 0.27577711}
)

// ClassificationResult stores the best prediction for each category.
type ClassificationResult struct {
	Garment TopResult `json:"garment"`
	Style   TopResult `json:"style"`
	Color   TopResult `json:"color"`
}

// TopResult represents the highest ranked label and confidence score.
type TopResult struct {
	Label      string  `json:"label"`
	Confidence float64 `json:"confidence"`
}

// combinedLabel represents one prompt combination.
type combinedLabel struct {
	color   string
	style   string
	garment string
}

// CLIPClassifier contains model sessions and precomputed text embeddings.
type CLIPClassifier struct {
	imageSession *ort.DynamicAdvancedSession
	textSession  *ort.DynamicAdvancedSession
	combinations []combinedLabel
	combEmbeds   [][]float32
}

// NewCLIPClassifier loads CLIP models and precomputes text embeddings
// for every garment, style and color combination.
func NewCLIPClassifier(modelDir string) (*CLIPClassifier, error) {

	// Configure and initialize ONNX Runtime.
	ort.SetSharedLibraryPath(sharedLibPath())
	if err := ort.InitializeEnvironment(); err != nil {
		return nil, fmt.Errorf("onnxruntime init: %w", err)
	}

	imageModelPath := modelDir + "/clip_image_encoder.onnx"
	textModelPath := modelDir + "/clip_text_encoder.onnx"

	// Ensure required model files exist.
	for _, p := range []string{imageModelPath, textModelPath} {
		if _, err := fileExists(p); err != nil {
			return nil, fmt.Errorf(
				"model file not found %q — run ./scripts/download_models.sh first: %w",
				p,
				err,
			)
		}
	}

	// Create image encoder session.
	imgSess, err := ort.NewDynamicAdvancedSession(
		imageModelPath,
		[]string{"pixel_values"},
		[]string{"image_embeds"},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("load image encoder: %w", err)
	}

	// Create text encoder session.
	txtSess, err := ort.NewDynamicAdvancedSession(
		textModelPath,
		[]string{"input_ids"},
		[]string{"text_embeds"},
		nil,
	)
	if err != nil {
		_ = imgSess.Destroy()
		return nil, fmt.Errorf("load text encoder: %w", err)
	}

	c := &CLIPClassifier{
		imageSession: imgSess,
		textSession:  txtSess,
	}

	// Generate all possible label combinations.
	// Total: 11 colors × 3 styles × 16 garments = 528 prompts.
	for _, color := range Colors {
		for _, style := range Styles {
			for _, garment := range Garments {
				c.combinations = append(c.combinations,
					combinedLabel{color, style, garment})
			}
		}
	}

	fmt.Printf(
		"Pre-computing embeddings for %d combinations...\n",
		len(c.combinations),
	)

	c.combEmbeds = make([][]float32, len(c.combinations))

	// Precompute text embeddings once to avoid recomputation at runtime.
	for i, comb := range c.combinations {

		// Example prompt:
		// "a photo of a yellow casual tshirt"
		prompt := fmt.Sprintf(
			"a photo of a %s %s %s",
			comb.color,
			comb.style,
			comb.garment,
		)

		tokenIDs := clipTokenize(prompt)

		inputTensor, err := ort.NewTensor(
			ort.NewShape(1, 77),
			tokenIDs,
		)
		if err != nil {
			return nil, err
		}

		outputTensor, err := ort.NewEmptyTensor[float32](
			ort.NewShape(1, embedDim),
		)
		if err != nil {
			_ = inputTensor.Destroy()
			return nil, err
		}

		if err := c.textSession.Run(
			[]ort.Value{ort.Value(inputTensor)},
			[]ort.Value{ort.Value(outputTensor)},
		); err != nil {
			_ = inputTensor.Destroy()
			_ = outputTensor.Destroy()
			return nil, fmt.Errorf(
				"text encoder for %q: %w",
				prompt,
				err,
			)
		}

		emb := make([]float32, embedDim)
		copy(emb, outputTensor.GetData())

		normalize(emb)
		c.combEmbeds[i] = emb

		_ = inputTensor.Destroy()
		_ = outputTensor.Destroy()
	}

	fmt.Printf("%d combinations ready\n", len(c.combinations))

	return c, nil
}

// Classify predicts garment, style and color for an input image.
func (c *CLIPClassifier) Classify(
	r io.Reader,
) (*ClassificationResult, error) {

	img, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	// Resize and normalize image for CLIP input.
	pixels := preprocessImage(img)

	inputTensor, err := ort.NewTensor(
		ort.NewShape(1, 3, clipImageSize, clipImageSize),
		pixels,
	)
	if err != nil {
		return nil, fmt.Errorf("create image tensor: %w", err)
	}
	defer func() {
		_ = inputTensor.Destroy()
	}()

	outputTensor, err := ort.NewEmptyTensor[float32](
		ort.NewShape(1, embedDim),
	)
	if err != nil {
		return nil, fmt.Errorf("create output tensor: %w", err)
	}
	defer func() {
		_ = outputTensor.Destroy()
	}()

	// Run image encoder inference.
	if err := c.imageSession.Run(
		[]ort.Value{ort.Value(inputTensor)},
		[]ort.Value{ort.Value(outputTensor)},
	); err != nil {
		return nil, fmt.Errorf("image encoder inference: %w", err)
	}

	imageEmbed := outputTensor.GetData()
	normalize(imageEmbed)

	// Compute cosine similarity against all prompt embeddings.
	scores := make([]float64, len(c.combinations))
	for i, emb := range c.combEmbeds {
		scores[i] = cosineSimilarity(imageEmbed, emb)
	}

	// Aggregate scores by category.
	colorScores := make(map[string]float64)
	styleScores := make(map[string]float64)
	garmentScores := make(map[string]float64)

	for i, s := range scores {
		comb := c.combinations[i]

		colorScores[comb.color] += s
		styleScores[comb.style] += s
		garmentScores[comb.garment] += s
	}

	// Select highest scoring prediction.
	topColor := rankMap(colorScores, 1)[0]
	topStyle := rankMap(styleScores, 1)[0]
	topGarment := rankMap(garmentScores, 1)[0]

	return &ClassificationResult{
		Color:   topColor,
		Style:   topStyle,
		Garment: topGarment,
	}, nil
}

func rankMap(
	scores map[string]float64,
	topK int,
) []TopResult {

	type pair struct {
		label string
		score float64
	}

	var total float64
	for _, s := range scores {
		total += s
	}

	var pairs []pair
	for k, v := range scores {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(
		pairs,
		func(i, j int) bool {
			return pairs[i].score > pairs[j].score
		},
	)

	if topK > len(pairs) {
		topK = len(pairs)
	}

	results := make([]TopResult, topK)

	for i, p := range pairs[:topK] {
		results[i] = TopResult{
			Label:      p.label,
			Confidence: math.Round(p.score/total*10000) / 100,
		}
	}

	return results
}

// Close releases ONNX Runtime resources.
func (c *CLIPClassifier) Close() {
	if c.imageSession != nil {
		_ = c.imageSession.Destroy()
	}
	if c.textSession != nil {
		_ = c.textSession.Destroy()
	}
}

// preprocessImage resizes and normalizes an image for CLIP.
func preprocessImage(img image.Image) []float32 {

	dst := image.NewRGBA(
		image.Rect(0, 0, clipImageSize, clipImageSize),
	)

	draw.BiLinear.Scale(
		dst,
		dst.Bounds(),
		img,
		img.Bounds(),
		draw.Over,
		nil,
	)

	pixels := make([]float32,
		3*clipImageSize*clipImageSize,
	)

	for y := 0; y < clipImageSize; y++ {
		for x := 0; x < clipImageSize; x++ {

			r, g, b, _ := dst.At(x, y).RGBA()
			idx := y*clipImageSize + x

			pixels[0*clipImageSize*clipImageSize+idx] =
				(float32(r>>8)/255.0 - mean[0]) / std[0]

			pixels[1*clipImageSize*clipImageSize+idx] =
				(float32(g>>8)/255.0 - mean[1]) / std[1]

			pixels[2*clipImageSize*clipImageSize+idx] =
				(float32(b>>8)/255.0 - mean[2]) / std[2]
		}
	}

	return pixels
}

// normalize converts a vector to unit length.
func normalize(v []float32) {
	var sum float64

	for _, x := range v {
		sum += float64(x) * float64(x)
	}

	norm := float32(math.Sqrt(sum))
	if norm == 0 {
		return
	}

	for i := range v {
		v[i] /= norm
	}
}

// cosineSimilarity computes cosine similarity between two embeddings.
func cosineSimilarity(a, b []float32) float64 {
	var s float64

	for i := range a {
		s += float64(a[i]) * float64(b[i])
	}

	return s
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}
