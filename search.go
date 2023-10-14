package main

import (
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"log"
	"time"
)

func main() {
	// Initialize TS client
	tsClient := typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"))

	start := time.Now()
	searchParameters := &api.SearchCollectionParams{
		Q:       "gone ",
		QueryBy: "title",
	}

	res, err := tsClient.Collection("books").Documents().Search(searchParameters)
	if err != nil {
		log.Fatal(err)
	}
	// Display the search results
	fmt.Printf("Total matches: %d\n", res.Found)
	elapsed := time.Since(start)
	fmt.Printf("searchingtime is %s\n", elapsed)

}
