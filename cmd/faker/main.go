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
			"type" : "string",
			"enum": ["United States of America", "Canada"]
		  }
		},
		"if": {
		  "properties": { "country": { "const": "United States of America" } }
		},
		"then": {
		  "properties": { "postal_code": { "type": "string"} }
		},
		"else": {
		  "properties": { "postal_code": { "type" : "number"} }
		}
	  }
	  
	`
	var prop map[string]interface{}

	json.Unmarshal([]byte(jsonSchema), &prop)

	// var prop = map[string]interface{}{
	// 	"type": "object",
	// 	"patternProperties": map[string]interface{}{
	// 		"^S_": map[string]interface{}{"type": "string"},
	// 		"^I_": map[string]interface{}{"type": "integer"},
	// 		"^K_": map[string]interface{}{
	// 			"type": "string",
	// 			"enum": []string{"Street", "Avenue", "Boulevard"},
	// 		},
	// 	},
	// 	// "propertyNames": map[string]interface{}{
	// 	// 	"pattern": "^[A-Za-z_][A-Za-z0-9_]*$",
	// 	// },
	// 	// "properties": map[string]interface{}{
	// 	// 	"num": map[string]interface{}{
	// 	// 		"type": "number",
	// 	// 	},
	// 	// 	"street_num": map[string]interface{}{
	// 	// 		"type": "number",
	// 	// 	},
	// 	// 	"name": map[string]interface{}{
	// 	// "patternProperties": map[string]interface{}{
	// 	// 	"^S_": map[string]interface{}{"type": "string"},
	// 	// 	"^I_": map[string]interface{}{"type": "integer", "multipleOf": 10, "minimum": 500},
	// 	// 	"^K_": map[string]interface{}{
	// 	// 		"type": "string",
	// 	// 	},
	// 	// 	"array": map[string]interface{}{
	// 	// 		"type": "array",
	// 	// 		"items": []map[string]interface{}{
	// 	// 			map[string]interface{}{
	// 	// 				"type": "number",
	// 	// 			},
	// 	// 			map[string]interface{}{
	// 	// 				"type": "number",
	// 	// 			},
	// 	// 			map[string]interface{}{
	// 	// 				"type": "string",
	// 	// 			},
	// 	// 		},
	// 	// 		"enum": []string{"Street", "Avenue", "Boulevard"},
	// 	// 	},
	// 	// },
	// 	// "propertyNames": map[string]interface{}{
	// 	// 	"pattern": "^[A-Za-z_][A-Za-z0-9_]*$",
	// 	// },
	// 	// "properties": map[string]interface{}{
	// 	// 	"num": map[string]interface{}{
	// 	// 		"type": "number",
	// 	// 	},
	// 	// 	"street_num": map[string]interface{}{
	// 	// 		"type": "number",
	// 	// 	},
	// 	// 	"name": map[string]interface{}{
	// 	// 		"type": "string",
	// 	// 	},
	// 	// 	"array": map[string]interface{}{
	// 	// 		"type": "array",
	// 	// 		"items": []map[string]interface{}{
	// 	// 			map[string]interface{}{
	// 	// 				"type": "number",
	// 	// 			},
	// 	// 			map[string]interface{}{
	// 	// 				"type": "number",
	// 	// 			},
	// 	// 			map[string]interface{}{
	// 	// 				"type": "string",
	// 	// 			},
	// 	// 		},
	// 	// 	},
	// 	// },
	// }

	//Application specific matching words
	matchingWords := []string{}
	//Wrap DataMatching for data matching of json schema
	x := matching.DataMatching(generate.GenerateObject(prop), matchingWords)
	fmt.Println(x)

}
