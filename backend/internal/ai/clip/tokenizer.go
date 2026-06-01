package clip

import (
	"strings"
	"unicode"
)

const (
	clipSOT    = 49406
	clipEOT    = 49407
	clipMaxLen = 77
)

var clipVocab = buildBasicVocab()

// clipTokenize converts text into CLIP token IDs.
func clipTokenize(text string) []int64 {

	tokens := []int64{clipSOT}

	text = strings.ToLower(strings.TrimSpace(text))

	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsSpace(r) ||
			r == '-' ||
			r == '_'
	})

	for _, word := range words {

		if id, ok := clipVocab[word]; ok {
			tokens = append(tokens, int64(id))
			continue
		}

		// Fallback to character-level tokenization.
		for _, r := range word {
			if id, ok := clipVocab[string(r)]; ok {
				tokens = append(tokens, int64(id))
			} else {
				tokens = append(tokens, 267)
			}
		}
	}

	tokens = append(tokens, clipEOT)

	result := make([]int64, clipMaxLen)
	copy(result, tokens)

	return result
}

// buildBasicVocab creates a lightweight CLIP vocabulary.
func buildBasicVocab() map[string]int {

	vocab := make(map[string]int)

	// ASCII characters used as the base CLIP vocabulary.
	for i := 33; i < 127; i++ {
		vocab[string(rune(i))] = i + 256
	}

	// Token IDs extracted from CLIP's original vocab.json.
	words := map[string]int{

		"a":     320,
		"photo": 1125,
		"of":    539,

		"tshirt":   14907,
		"long":     1538,
		"sleeve":   10536,
		"hoodie":   13444,
		"sweater":  11455,
		"shorts":   9680,
		"skirt":    12386,
		"sandals":  20731,
		"jeans":    10157,
		"sneakers": 17397,
		"boots":    7319,
		"coat":     7356,
		"jacket":   6164,
		"top":      1253,
		"shirt":    2523,
		"shoes":    4079,
		"heels":    10909,

		"black":  1449,
		"white":  1579,
		"gray":   7048,
		"blue":   1746,
		"red":    736,
		"green":  1901,
		"yellow": 4481,
		"brown":  2866,
		"pink":   3360,
		"purple": 5496,
		"orange": 4287,

		"casual": 10129,
		"sporty": 33346,
		"formal": 12978,
	}

	for k, v := range words {
		vocab[k] = v
	}

	return vocab
}
