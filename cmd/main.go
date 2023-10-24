package main

import (
	"ankitsridhar16/book-search-typesense/internal/postgres"
	"ankitsridhar16/book-search-typesense/internal/typesense"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {
	// Load yaml config
	fileContent, fileErr := ioutil.ReadFile("../config.yaml")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	var config map[string]string
	configErr := yaml.Unmarshal(fileContent, &config)
	if configErr != nil {
		log.Fatal(configErr)
	}

	// Setup Postgres connection
	pgDB, dbErr := postgres.Init(config["postgres_url"])
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer pgDB.Close()

	// Setup TypeSense connection
	_ = typesense.NewClient(config["ts_server_url"], config["ts_api_key"])

	// Fetch data from DB
	_, dbErr = pgDB.FetchDataFromDB()
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}
