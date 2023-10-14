package main

import (
	"encoding/json"
	"fmt"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"os"
	"time"
)

type Book struct {
	Title           string   `json:"title"`
	Authors         []string `json:"authors"`
	PublicationYear int      `json:"publication_year"`
	RatingsCount    int      `json:"ratings_count"`
	AverageRating   float64  `json:"average_rating"`
}

func main() {
	// Initialize Typesense client
	tsClient := typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"),
	)

	// Define the schema for the collection
	schema := &api.CollectionSchema{
		Name: "books",
		Fields: []api.Field{
			{
				Name: "title",
				Type: "string",
			},
			{
				Name: "authors",
				Type: "string[]",
			},
			{
				Name: "publication_year",
				Type: "int32",
			},
			{
				Name: "ratings_count",
				Type: "int32",
			},
			{
				Name: "average_rating",
				Type: "float",
			},
		},
	}

	fmt.Println("Populating index in Typesense")

	// Delete the existing collection if it exists
	_, err := tsClient.Collection("books").Delete()
	if err == nil {
		fmt.Println("Deleting existing collection: books")
	}

	fmt.Println("Creating schema: ")
	fmt.Println(schema)

	// Create the collection with the specified schema
	_, err = tsClient.Collections().Create(schema)
	if err != nil {
		fmt.Println("Error creating collection: ", err)
		return
	}

	fmt.Println("Adding records: ")
	start := time.Now()
	books := loadBooksFromJSON("./data/books.json")

	var batchSize *int
	size := 40
	batchSize = &size
	params := &api.ImportDocumentsParams{
		BatchSize: batchSize,
	}
	_, err = tsClient.Collection("books").Documents().Import(books, params)
	if err != nil {
		fmt.Println("Error importing documents: ", err)
		return
	}

	fmt.Println("Done indexing.")
	elapsed := time.Since(start)
	fmt.Printf("Total indexing time is %s\n", elapsed)
}

func loadBooksFromJSON(filename string) []interface{} {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file: ", err)
		return nil
	}
	defer file.Close()

	var books []Book
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&books)
	if err != nil {
		fmt.Println("Error decoding JSON data: ", err)
		return nil
	}

	var documents []interface{}
	for _, book := range books {
		documents = append(documents, book)
	}
	return documents
}
