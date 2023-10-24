package typesense

import (
	"github.com/typesense/typesense-go/typesense"
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
