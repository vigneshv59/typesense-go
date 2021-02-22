package typesense

import (
	"context"

	"github.com/vigneshv59/typesense-go/typesense/api"
)

type KeyInterface interface {
	Retrieve() (*api.ApiKey, error)
	Delete() (*api.ApiKey, error)
}

type key struct {
	apiClient APIClientInterface
	keyID     int64
}

func (k *key) Retrieve() (*api.ApiKey, error) {
	response, err := k.apiClient.GetKeyWithResponse(context.Background(), k.keyID)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200, nil
}

func (k *key) Delete() (*api.ApiKey, error) {
	response, err := k.apiClient.DeleteKeyWithResponse(context.Background(), k.keyID)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200, nil
}
