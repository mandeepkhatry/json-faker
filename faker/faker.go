package faker

import (
	"faker/def"
	"faker/generate"
	"faker/utils"
	"faker/validation"
	"fmt"
	"strconv"
)

type FakerItem struct {
	Definations interface{}
	Schema      interface{}
}

func Faker(schema interface{}) *FakerItem {
	if definations, present := schema.(map[string]interface{})["definitions"]; present {
		return &FakerItem{
			Definations: definations,
			Schema:      schema.(map[string]interface{}),
		}
	}
	return &FakerItem{
		Schema: schema.(map[string]interface{}),
	}
}

func (f *FakerItem) Build() map[string]interface{} {
	return f.GenerateObject(f.Schema.(map[string]interface{}))
}

func (f *FakerItem) GenerateObject(properties map[string]interface{}) map[string]interface{} {

	if _, present := properties["propertyNames"]; present {
		pattern := properties["propertyNames"].(map[string]interface{})["pattern"].(string)
		return map[string]interface{}{
			generate.GenerateRegex(pattern): "value",
		}
	}

	generatedObject := make(map[string]interface{}, 0)

	if _, present := properties["patternProperties"]; present {

		patternProperties := properties["patternProperties"].(map[string]interface{})

		for field, fieldProperties := range patternProperties {

			fProperties := fieldProperties

			if ref, isReferenced := fieldProperties.(map[string]interface{})["$ref"]; isReferenced {
				fProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			fieldType := fProperties.(map[string]interface{})["type"].(string)
			patternFieldName := generate.GenerateRegex(field)

			if fieldType == "array" {
				generatedObject[patternFieldName] = f.GenerateArray(fProperties.(map[string]interface{}))
			} else if fieldType == "object" {
				generatedObject[patternFieldName] = f.GenerateObject(fProperties.(map[string]interface{}))
			} else {

				generatedObject[patternFieldName] = generate.FieldToGenerator[fieldType](fProperties.(map[string]interface{}))
			}

		}

	}

	if _, present := properties["properties"]; present {
		objectProperties := properties["properties"].(map[string]interface{})

		for field, fieldProperties := range objectProperties {

			fProperties := fieldProperties

			if ref, isReferenced := fieldProperties.(map[string]interface{})["$ref"]; isReferenced {
				fProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			fieldType := fProperties.(map[string]interface{})["type"].(string)

			if fieldType == "array" {
				generatedObject[field] = f.GenerateArray(fProperties.(map[string]interface{}))
			} else if fieldType == "object" {
				generatedObject[field] = f.GenerateObject(fProperties.(map[string]interface{}))
			} else {
				generatedObject[field] = generate.FieldToGenerator[fieldType](fProperties.(map[string]interface{}))
			}

		}
	}

	if _, present := properties["allOf"]; present {
		for _, eachAllOfProp := range properties["allOf"].([]interface{}) {

			allOfProperties := eachAllOfProp

			if ref, isReferenced := eachAllOfProp.(map[string]interface{})["$ref"]; isReferenced {
				allOfProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			allOfGeneratedObject := make(map[string]interface{})

			if _, present := allOfProperties.(map[string]interface{})["if"]; present {

				ifProperties := allOfProperties.(map[string]interface{})["if"]

				if ref, isReferenced := allOfProperties.(map[string]interface{})["if"].(map[string]interface{})["$ref"]; isReferenced {
					ifProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
				}

				isValidated, _ := validation.ValidateDataWithRespectToSchema(ifProperties.(map[string]interface{}), generatedObject)

				if isValidated {
					if _, present := allOfProperties.(map[string]interface{})["then"]; present {

						thenProperties := allOfProperties.(map[string]interface{})["then"]

						if ref, isReferenced := allOfProperties.(map[string]interface{})["then"].(map[string]interface{})["$ref"]; isReferenced {
							thenProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
						}

						thenObject := f.GenerateObject(thenProperties.(map[string]interface{}))

						for k, v := range thenObject {
							allOfGeneratedObject[k] = v
						}
					}
				} else {
					if _, present := allOfProperties.(map[string]interface{})["else"]; present {
						elseProperties := allOfProperties.(map[string]interface{})["else"]

						if ref, isReferenced := allOfProperties.(map[string]interface{})["else"].(map[string]interface{})["$ref"]; isReferenced {
							elseProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
						}
						elseObject := f.GenerateObject(elseProperties.(map[string]interface{}))
						for k, v := range elseObject {
							allOfGeneratedObject[k] = v
						}
					}

				}
			} else {
				allOfGeneratedObject = f.GenerateObject(allOfProperties.(map[string]interface{}))
			}

			for k, v := range allOfGeneratedObject {
				generatedObject[k] = v
			}

		}
	}

	if _, present := properties["oneOf"]; present {
		for _, eachOneOfProp := range properties["oneOf"].([]interface{}) {

			oneOfProperties := eachOneOfProp

			if ref, isReferenced := eachOneOfProp.(map[string]interface{})["$ref"]; isReferenced {
				oneOfProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			generatedObject = f.GenerateObject(oneOfProperties.(map[string]interface{}))
			break

		}
	}

	if _, present := properties["anyOf"]; present {
		for _, eachAnyProp := range properties["anyOf"].([]interface{}) {

			anyOfProperties := eachAnyProp

			if ref, isReferenced := eachAnyProp.(map[string]interface{})["$ref"]; isReferenced {
				anyOfProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			generatedObject = f.GenerateObject(anyOfProperties.(map[string]interface{}))
			break
		}
	}

	if _, present := properties["if"]; present {

		ifProperties := properties["if"]

		if ref, isReferenced := properties["if"].(map[string]interface{})["$ref"]; isReferenced {
			ifProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
		}

		isValidated, _ := validation.ValidateDataWithRespectToSchema(ifProperties.(map[string]interface{}), generatedObject)

		if isValidated {
			if _, present := properties["then"]; present {

				thenProperties := properties["then"]

				if ref, isReferenced := properties["then"].(map[string]interface{})["$ref"]; isReferenced {
					thenProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
				}

				thenObject := f.GenerateObject(thenProperties.(map[string]interface{}))

				for k, v := range thenObject {
					generatedObject[k] = v
				}
			}
		} else {
			if _, present := properties["else"]; present {
				elseProperties := properties["else"]

				if ref, isReferenced := properties["else"].(map[string]interface{})["$ref"]; isReferenced {
					elseProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
				}
				elseObject := f.GenerateObject(elseProperties.(map[string]interface{}))
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

func (f *FakerItem) GenerateArray(properties map[string]interface{}) interface{} {

	minItems := def.DummyDataRange["array"].(map[string]int)["minItems"]

	if min, present := properties["minItems"]; present {
		minItems = int(min.(float64))
	}

	arrayItemType := fmt.Sprintf("%T", properties["items"])

	if arrayItemType == "[]interface {}" {

		resultingArray := make([]interface{}, 0)

		for _, eachTypeProperties := range properties["items"].([]interface{}) {

			eachProperties := eachTypeProperties

			if ref, isReferenced := eachTypeProperties.(map[string]interface{})["$ref"]; isReferenced {
				eachProperties, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
			}

			eachArrayType := eachProperties.(map[string]interface{})["type"].(string)

			if eachArrayType == "array" {
				resultingArray = append(resultingArray, f.GenerateArray(eachProperties.(map[string]interface{})))
			} else if eachArrayType == "object" {
				resultingArray = append(resultingArray, f.GenerateObject(eachProperties.(map[string]interface{})))
			} else {
				resultingArray = append(resultingArray, generate.FieldToGenerator[eachArrayType](eachProperties.(map[string]interface{})))
			}

		}
		return resultingArray

		//Mix Type Array

	} else if arrayItemType == "map[string]interface {}" {
		resultingArray := make([]interface{}, 0)

		prop := properties["items"]

		if ref, isReferenced := properties["items"].(map[string]interface{})["$ref"]; isReferenced {
			prop, _ = utils.GetReference(ref.(string), f.Definations.(map[string]interface{}))
		}

		//Single Type Arrays

		itemType := prop.(map[string]interface{})["type"].(string)

		if itemType == "array" {
			for i := 0; i < minItems; i++ {
				resultingArray = append(resultingArray, f.GenerateArray(prop.(map[string]interface{})))
			}

		} else if itemType == "object" {
			for i := 0; i < minItems; i++ {
				resultingArray = append(resultingArray, f.GenerateObject(prop.(map[string]interface{})))
			}
		} else {
			for i := 0; i < minItems; i++ {
				resultingArray = append(resultingArray, generate.FieldToGenerator[itemType](prop.(map[string]interface{})))
			}
		}

		return resultingArray
	}

	return []string{}

}
