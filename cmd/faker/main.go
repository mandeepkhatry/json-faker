package main

import (
	"encoding/json"
	"faker/generate"
	"faker/matching"
	"fmt"
	"log"
)

func main() {
	log.Println("[test]")

	jsonSchema := `
	{
		"type": "object",
		"properties": {
		  "street_address": {
			"type": "string"
		  },
		  "country": {
			"type":"string",
			"enum": ["United States of America", "Canada", "Netherlands"]
		  }
		},
		"allOf": [
		  {
			"if": {
				"properties": { "country": { "const": "United States of America" } }
			  },
			  "then": {
				"properties": { "postal_code": { "type": "string"} }
			  }
		  },
		  {
			"if": {
				"properties": { "country": { "const": "Canada" } }
			  },
			  "then": {
				"properties": { "postal_code": { "type": "number"} }
			  }
		  },
		  {
			"if": {
				"properties": { "country": { "const": "Netherlands" } }
			  },
			  "then": {
				"properties": { "postal_code": { "type": "integer"} }
			  }
		  }
		]
	  }

	`

	var prop map[string]interface{}
	json.Unmarshal([]byte(jsonSchema), &prop)

	//Application specific matching words
	matchingWords := []string{}
	//Wrap DataMatching for data matching of json schema
	x := matching.DataMatching(generate.GenerateObject(prop), matchingWords)
	fmt.Println(x)

}
