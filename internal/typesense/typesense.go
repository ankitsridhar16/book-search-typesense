package typesense

import (
	"ankitsridhar16/book-search-typesense/internal/postgres"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"strconv"
)

type TSClient struct {
	*typesense.Client
}

// NewClient initialize TypeSense client
func NewClient(tsConnSerr, apiKey string) *TSClient {
	tsClient := typesense.NewClient(
		typesense.WithServer(tsConnSerr),
		typesense.WithAPIKey(apiKey))
	return &TSClient{tsClient}
}

// CreateCollection creates a new collection if it doesn't exist based on collectionName
func (tsClient TSClient) CreateCollection(collectionName string) error {
	schema := &api.CollectionSchema{
		Name: collectionName,
		Fields: []api.Field{
			{
				Name: "id",
				Type: "string",
			},
			{
				Name: "title",
				Type: "string",
			},
			{
				Name: "publicationYear",
				Type: "int32",
			},
			{
				Name: "averageRating",
				Type: "float",
			},
			{
				Name: "imageURL",
				Type: "string",
			},
			{
				Name: "ratingsCount",
				Type: "int32",
			},
		},
	}

	_, collErr := tsClient.Collections().Create(schema)
	if collErr != nil {
		return collErr
	}
	return nil
}

// IndexData indexes the data to typesense based on collection name
func (tsClient TSClient) IndexData(collName string, books []postgres.Book) error {
	for _, book := range books {
		doc := map[string]interface{}{
			"id":              strconv.Itoa(book.ID),
			"title":           book.Title,
			"publicationYear": book.PublicationYear,
			"averageRating":   book.AverageRating,
			"imageURL":        book.ImageURL,
			"ratingsCount":    book.RatingsCount,
		}
		doc, docErr := tsClient.Collection(collName).Documents().Create(doc)
		if docErr != nil {
			return docErr
		}
	}

	return nil
}
