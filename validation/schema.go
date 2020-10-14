package validation

import (
	"faker/def"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateDataWithRespectToSchema(schema map[string]interface{}, data map[string]interface{}) (bool, error) {
	loader := gojsonschema.NewGoLoader(schema)
	docLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(loader, docLoader)
	if err != nil {
		return false, err
	}

	if !result.Valid() {
		return false, def.ErrInvalidDataAsPerSchema
	}

	return true, nil

}
