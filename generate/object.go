package generate

import (
	"faker/validation"
	"strconv"
)

func GenerateObject(properties map[string]interface{}) map[string]interface{} {

	if _, present := properties["propertyNames"]; present {
		pattern := properties["propertyNames"].(map[string]interface{})["pattern"].(string)
		return map[string]interface{}{
			GenerateRegex(pattern): "value",
		}
	}

	generatedObject := make(map[string]interface{}, 0)

	if _, present := properties["patternProperties"]; present {

		patternProperties := properties["patternProperties"].(map[string]interface{})

		for field, fieldProperties := range patternProperties {
			fieldType := fieldProperties.(map[string]interface{})["type"].(string)
			patternFieldName := GenerateRegex(field)

			if fieldType == "array" {
				generatedObject[patternFieldName] = GenerateArray(fieldProperties.(map[string]interface{}))
			} else if fieldType == "object" {
				generatedObject[patternFieldName] = GenerateObject(fieldProperties.(map[string]interface{}))
			} else {

				generatedObject[patternFieldName] = FieldToGenerator[fieldType](fieldProperties.(map[string]interface{}))
			}

		}

	}

	if _, present := properties["properties"]; present {
		objectProperties := properties["properties"].(map[string]interface{})

		for field, fieldProperties := range objectProperties {
			fieldType := fieldProperties.(map[string]interface{})["type"].(string)

			if fieldType == "array" {
				generatedObject[field] = GenerateArray(fieldProperties.(map[string]interface{}))
			} else if fieldType == "object" {
				generatedObject[field] = GenerateObject(fieldProperties.(map[string]interface{}))
			} else {
				generatedObject[field] = FieldToGenerator[fieldType](fieldProperties.(map[string]interface{}))
			}

		}
	}

	if _, present := properties["allOf"]; present {
		for _, eachAllOfProp := range properties["allOf"].([]interface{}) {
			for k, v := range ExecuteAllOf(eachAllOfProp.(map[string]interface{}), generatedObject) {
				generatedObject[k] = v
			}
		}
	}

	if _, present := properties["oneOf"]; present {
		for _, eachAllOfProp := range properties["oneOf"].([]interface{}) {
			generatedObject = GenerateObject(eachAllOfProp.(map[string]interface{}))
			break

		}
	}

	if _, present := properties["anyOf"]; present {
		for _, eachAllOfProp := range properties["anyOf"].([]interface{}) {
			generatedObject = GenerateObject(eachAllOfProp.(map[string]interface{}))
			break
		}
	}

	if _, present := properties["if"]; present {

		isValidated, _ := validation.ValidateDataWithRespectToSchema(properties["if"].(map[string]interface{}), generatedObject)

		if isValidated {
			if _, present := properties["then"]; present {
				thenObject := GenerateObject(properties["then"].(map[string]interface{}))

				for k, v := range thenObject {
					generatedObject[k] = v
				}
			}
		} else {
			if _, present := properties["else"]; present {
				elseObject := GenerateObject(properties["else"].(map[string]interface{}))
				for k, v := range elseObject {
					generatedObject[k] = v
				}
			}

		}
	}

	if len(generatedObject) == 0 {
		if _, minPropertiesPresent := properties["minProperties"]; minPropertiesPresent {
			minPropertiesPresent := int(properties["minProperties"].(float64))

			for i := 0; i < minPropertiesPresent; i++ {
				generatedObject["test"+strconv.Itoa(i)] = i
			}

			return generatedObject
		}

		generatedObject["k1"] = "v1"
		generatedObject["k2"] = "v2"
	}

	return generatedObject

}
