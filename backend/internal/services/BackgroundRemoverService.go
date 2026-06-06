package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type RemoveBGService struct {
	apiKey string
}

func NewRemoveBGService(apiKey string) *RemoveBGService {
	return &RemoveBGService{
		apiKey: apiKey,
	}
}

func (s *RemoveBGService) RemoveBackground(
	imageBytes []byte,
) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(
		"image_file",
		"image.jpg",
	)
	if err != nil {
		return nil, err
	}

	_, err = part.Write(imageBytes)
	if err != nil {
		return nil, err
	}

	_ = writer.WriteField("size", "auto")

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.remove.bg/v1.0/removebg",
		body,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Api-Key", s.apiKey)
	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {

		msg, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf(
			"remove.bg error: %s",
			string(msg),
		)
	}

	return io.ReadAll(resp.Body)
}
