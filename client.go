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

// ClientDefault returns a new MugshotClient with the provided API key.
// It takes a string parameter apikey and returns a pointer to MugshotClient.
func ClientDefault(apikey string) *MugshotClient {
	return &MugshotClient{
		ApiKey: apikey,
		Option: Option{
			Endpoint: "https://v1.mugshot.dev",
		},
	}
}

// Client initializes and returns a MugshotClient with the provided API key and option.
//
// Parameters:
//
//	apikey string - the API key for authentication
//	option Option - the option for the Mugshot client
//
// Returns:
//
//	*MugshotClient - the initialized Mugshot client
func Client(apikey string, option Option) *MugshotClient {
	return &MugshotClient{
		ApiKey: apikey,
		Option: option,
	}
}

// AddFace adds a face to the MugshotClient.
//
// It takes an image file and metadata as parameters and returns an AddFaceResponse
// pointer and an error.
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

// SearchFace searches for a face in the given image file.
//
// imageFile: io.Reader containing the image file.
// Returns *SearchFaceResponse and error.
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

// SearchFaceFirst searches for a face in an image file.
//
// It takes an io.Reader imageFile as a parameter and returns a *SearchFaceResponse
// and an error.
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

// MatchFace sends a face image to the Mugshot API for matching and returns the match result.
//
// imageFile io.Reader
// *MatchFaceResponse, error
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

// DeleteFace deletes a face with the given face ID.
// It returns a pointer to DeleteFaceResponse and an error.
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

// mapToJSON maps the given data to JSON format.
//
// data: a map of string to interface.
// string: the JSON representation of the given data.
func (c *MugshotClient) mapToJSON(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
