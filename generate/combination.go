package generate

import (
	"faker/validation"
)

func ExecuteAllOf(properties map[string]interface{}, additionalData map[string]interface{}) map[string]interface{} {

	generatedObject := make(map[string]interface{}, 0)

	if _, present := properties["if"]; present {

		isValidated, _ := validation.ValidateDataWithRespectToSchema(properties["if"].(map[string]interface{}), additionalData)

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
	} else {
		return GenerateObject(properties)
	}

	return generatedObject

}
