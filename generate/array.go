package generate

import (
	"faker/def"
	"fmt"
)

func GenerateArray(properties map[string]interface{}) interface{} {

	if _, typePresent := properties["items"]; !typePresent {
		var maxItems int

		if _, present := properties["maxItems"]; !present {
			maxItems = def.DummyDataRange["array"].(map[string]int)["maximum"]
		} else {
			maxItems = int(properties["maxItems"].(float64))
		}

		// if _, present := properties["minItems"]; !present {
		// 	minItems = 0
		// } else {
		// 	minItems = properties["minItems"].(int)
		// }

		numberArray := make([]float64, 0)

		var arrayProperties = map[string]interface{}{
			"maximum": def.DummyDataRange["number"].(map[string]int)["maximum"],
			"minimum": def.DummyDataRange["number"].(map[string]int)["minimum"],
		}

		for i := 0; i < maxItems; i++ {
			numberArray = append(numberArray, GenerateFloat(arrayProperties))
		}

		return numberArray
	}

	arrayItemType := fmt.Sprintf("%T", properties["items"])

	if arrayItemType == "[]map[string]interface {}" {

		resultingArray := make([]interface{}, 0)

		for _, eachTypeProperties := range properties["items"].([]map[string]interface{}) {

			eachArrayType := eachTypeProperties["type"].(string)

			if eachArrayType == "array" {
				resultingArray = append(resultingArray, GenerateArray(eachTypeProperties))
			} else if eachArrayType == "object" {
				resultingArray = append(resultingArray, GenerateObject(eachTypeProperties))
			} else {
				resultingArray = append(resultingArray, FieldToGenerator[eachArrayType](eachTypeProperties))
			}

		}
		return resultingArray

		//Mix Type Array

	} else if arrayItemType == "map[string]interface {}" {

		//Single Type Arrays

		itemType := properties["items"].(map[string]interface{})["type"].(string)

		if itemType == "array" {
			return GenerateArray(properties["items"].(map[string]interface{}))
		} else if itemType == "object" {
			return []interface{}{GenerateObject(properties["items"].(map[string]interface{}))}
		}
	}

	return []string{}

}
