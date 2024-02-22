package mugshot_go

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"io"
)

type Option struct {
	Endpoint string
}

type MugshotClient struct {
	ApiKey string
	Option Option
}

func ClientDefault(apikey string) *MugshotClient {
	return &MugshotClient{
		ApiKey: apikey,
		Option: Option{
			Endpoint: "https://v1.mugshot.dev",
		},
	}
}

func Client(apikey string, option Option) *MugshotClient {
	return &MugshotClient{
		ApiKey: apikey,
		Option: option,
	}
}

func (c *MugshotClient) AddFace(imageFile io.Reader, metadata map[string]interface{}) (*AddFaceResponse, error) {
	url := c.Option.Endpoint + "/face/add"

	resp, err := resty.New().R().
		SetFileReader("image", "image.jpg", imageFile).
		SetFormData(map[string]string{"metadata": c.mapToJSON(metadata)}).
		SetHeader("Authorization", c.ApiKey).
		SetHeader("User-Agent", "Mugshot-SDK/1.0.0").
		Post(url)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("HTTP error! Status: " + resp.Status())
	}

	var data AddFaceResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *MugshotClient) SearchFace(imageFile io.Reader) (*SearchFaceResponse, error) {
	url := c.Option.Endpoint + "/face/find"

	resp, err := resty.New().R().
		SetFileReader("image", "image.jpg", imageFile).
		SetHeader("Authorization", c.ApiKey).
		SetHeader("User-Agent", "Mugshot-SDK/1.0.0").
		Post(url)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("HTTP error! Status: " + resp.Status())
	}

	var data SearchFaceResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *MugshotClient) SearchFaceFirst(imageFile io.Reader) (*SearchFaceResponse, error) {
	url := c.Option.Endpoint + "/face/find"

	resp, err := resty.New().R().
		SetFileReader("image", "image.jpg", imageFile).
		SetHeader("Authorization", c.ApiKey).
		SetHeader("User-Agent", "Mugshot-SDK/1.0.0").
		Post(url)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("HTTP error! Status: " + resp.Status())
	}

	var data SearchFaceResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *MugshotClient) MatchFace(imageFile io.Reader) (*MatchFaceResponse, error) {
	url := c.Option.Endpoint + "/face/find/match"
	resp, err := resty.New().R().
		SetFileReader("image", "image.jpg", imageFile).
		SetHeader("Authorization", c.ApiKey).
		SetHeader("User-Agent", "Mugshot-SDK/1.0.0").
		Post(url)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("HTTP error! Status: " + resp.Status())
	}

	var data MatchFaceResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *MugshotClient) DeleteFace(faceId string) (*DeleteFaceResponse, error) {
	url := c.Option.Endpoint + "/face/delete"
	resp, err := resty.New().R().
		SetFormData(map[string]string{"face_id": faceId}).
		SetHeader("Authorization", c.ApiKey).
		SetHeader("User-Agent", "Mugshot-SDK/1.0.0").
		Post(url)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("HTTP error! Status: " + resp.Status())
	}

	var data DeleteFaceResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *MugshotClient) mapToJSON(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
