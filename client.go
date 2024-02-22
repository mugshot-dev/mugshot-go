package mugshot_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

type Option struct {
	Endpoint string
}
type MugshotClient struct {
	ApiKey string
	Option Option
}

func ClientDefault(apikey string) *MugshotClient {
	return &MugshotClient{apikey, Option{
		Endpoint: "https://v1.mugshot.dev",
	}}
}

func Client(apikey string, option Option) *MugshotClient {
	return &MugshotClient{
		ApiKey: apikey,
		Option: option,
	}
}

func (c *MugshotClient) AddFace(imageFile io.Reader, metadata map[string]interface{}) (*AddFaceResponse, error) {
	url := c.Option.Endpoint + "/face/add"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	imagePart, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(imagePart, imageFile); err != nil {
		return nil, err
	}

	metadataPart, err := writer.CreateFormField("metadata")
	if err != nil {
		return nil, err
	}
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}
	metadataPart.Write(metadataJSON)

	writer.Close()

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", c.ApiKey)
	request.Header.Set("User-Agent", "Mugshot-SDK/1.0.0")
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP error! Status: " + response.Status)
	}

	var data AddFaceResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
