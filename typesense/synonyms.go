package typesense

import (
	"context"

	"github.com/vigneshv59/typesense-go/typesense/api"
)

// SynonymsInterface is a type for Search Synonyms API operations
type SynonymsInterface interface {
	// Create or update a synonym
	Upsert(synonymID string, synonymSchema *api.SearchSynonymSchema) (*api.SearchSynonym, error)
	// List all collection synonyms
	Retrieve() ([]*api.SearchSynonym, error)
}

// synonyms is internal implementation of SynonymsInterface
type synonyms struct {
	apiClient      APIClientInterface
	collectionName string
}

func (s *synonyms) Upsert(synonymID string, synonymSchema *api.SearchSynonymSchema) (*api.SearchSynonym, error) {
	response, err := s.apiClient.UpsertSearchSynonymWithResponse(context.Background(),
		s.collectionName, synonymID, api.UpsertSearchSynonymJSONRequestBody(*synonymSchema))
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200, nil
}

func (s *synonyms) Retrieve() ([]*api.SearchSynonym, error) {
	response, err := s.apiClient.GetSearchSynonymsWithResponse(context.Background(), s.collectionName)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200.Synonyms, nil
}
