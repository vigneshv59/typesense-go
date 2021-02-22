package typesense

import (
	"context"

	"github.com/vigneshv59/typesense-go/typesense/api"
)

// CollectionsInterface is a type for Collections API operations
type CollectionsInterface interface {
	Create(schema *api.CollectionSchema) (*api.Collection, error)
	Retrieve() ([]*api.Collection, error)
}

// collections is internal implementation of CollectionsInterface
type collections struct {
	apiClient APIClientInterface
}

func (c *collections) Create(schema *api.CollectionSchema) (*api.Collection, error) {
	response, err := c.apiClient.CreateCollectionWithResponse(context.Background(),
		api.CreateCollectionJSONRequestBody(*schema))
	if err != nil {
		return nil, err
	}
	if response.JSON201 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON201, nil
}

func (c *collections) Retrieve() ([]*api.Collection, error) {
	response, err := c.apiClient.GetCollectionsWithResponse(context.Background())
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return *response.JSON200, nil
}
