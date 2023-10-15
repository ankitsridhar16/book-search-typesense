package main

import (
	"fmt"
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

	var config map[string]interface{}
	configErr := yaml.Unmarshal(fileContent, &config)
	if configErr != nil {
		log.Fatal(configErr)
	}
	fmt.Println(config["postgres_url"])
}
