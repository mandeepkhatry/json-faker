package main

import (
	"faker/generate"
	"faker/matching"
	"fmt"
	"log"
)

func main() {
	log.Println("[test]")

	var prop = map[string]interface{}{
		"type": "object",
		"patternProperties": map[string]interface{}{
			"^S_": map[string]interface{}{"type": "string"},
			"^I_": map[string]interface{}{"type": "integer"},
			"^K_": map[string]interface{}{
				"type": "string",
				"enum": []string{"Street", "Avenue", "Boulevard"},
			},
		},
		// "propertyNames": map[string]interface{}{
		// 	"pattern": "^[A-Za-z_][A-Za-z0-9_]*$",
		// },
		// "properties": map[string]interface{}{
		// 	"num": map[string]interface{}{
		// 		"type": "number",
		// 	},
		// 	"street_num": map[string]interface{}{
		// 		"type": "number",
		// 	},
		// 	"name": map[string]interface{}{
		// "patternProperties": map[string]interface{}{
		// 	"^S_": map[string]interface{}{"type": "string"},
		// 	"^I_": map[string]interface{}{"type": "integer", "multipleOf": 10, "minimum": 500},
		// 	"^K_": map[string]interface{}{
		// 		"type": "string",
		// 	},
		// 	"array": map[string]interface{}{
		// 		"type": "array",
		// 		"items": []map[string]interface{}{
		// 			map[string]interface{}{
		// 				"type": "number",
		// 			},
		// 			map[string]interface{}{
		// 				"type": "number",
		// 			},
		// 			map[string]interface{}{
		// 				"type": "string",
		// 			},
		// 		},
		// 		"enum": []string{"Street", "Avenue", "Boulevard"},
		// 	},
		// },
		// "propertyNames": map[string]interface{}{
		// 	"pattern": "^[A-Za-z_][A-Za-z0-9_]*$",
		// },
		// "properties": map[string]interface{}{
		// 	"num": map[string]interface{}{
		// 		"type": "number",
		// 	},
		// 	"street_num": map[string]interface{}{
		// 		"type": "number",
		// 	},
		// 	"name": map[string]interface{}{
		// 		"type": "string",
		// 	},
		// 	"array": map[string]interface{}{
		// 		"type": "array",
		// 		"items": []map[string]interface{}{
		// 			map[string]interface{}{
		// 				"type": "number",
		// 			},
		// 			map[string]interface{}{
		// 				"type": "number",
		// 			},
		// 			map[string]interface{}{
		// 				"type": "string",
		// 			},
		// 		},
		// 	},
		// },
	}

	//Wrap DataMatching for data matching of json schema
	x := matching.DataMatching(generate.GenerateObject(prop))
	fmt.Println(x)
}