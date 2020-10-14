package generate

import (
	"faker/def"
)

func GenerateStringFormat(stringType string) string {
	return def.StringFormat[stringType]

}
