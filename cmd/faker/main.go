package main

import (
	"encoding/json"
	"faker/faker"
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

	fmt.Println("Fake Data : ", faker.Faker(prop).Build())

	/*
	  To match words lets say : username and confirm_username
	  where format of similar field is matchingWord_{field}

	  Different regions have different way of using the mactching words.

	  Wrap DataMatching function over faker build and pass array of matching words.
	*/
	matchingWords := []string{"confirm"}

	fmt.Println("Passing data matching : ", matching.DataMatching(faker.Faker(prop).Build(), matchingWords))

}
