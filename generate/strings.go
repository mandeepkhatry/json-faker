package generate

import (
	"faker/def"
	"math/rand"
	"time"
)

func GenerateString(properties map[string]interface{}) string {

	if _, present := properties["enum"]; present {
		if len(properties["enum"].([]string)) != 0 {
			return properties["enum"].([]string)[0]
		}

	}

	if _, present := properties["constant"]; present {
		return properties["constant"].(string)

	}

	rand.Seed(time.Now().UnixNano())

	var minLength, maxLength int

	if _, present := properties["maxLength"]; !present {
		maxLength = def.DummyDataRange["string"].(map[string]int)["maxLength"]
	} else {
		maxLength = int(properties["maxLength"].(float64))
	}

	if _, present := properties["minLength"]; !present {
		minLength = def.DummyDataRange["string"].(map[string]int)["minLength"]
	} else {
		minLength = int(properties["minLength"].(float64))
	}

	if _, present := properties["enum"]; present {
		enum := properties["enum"].([]string)
		if len(enum) > 0 {
			return enum[0]
		}
		return "test"
	}

	if _, present := properties["format"]; present {
		return GenerateStringFormat(properties["format"].(string))
	}

	if _, present := properties["pattern"]; present {
		return GenerateRegex(properties["pattern"].(string))
	}

	stringLength := rand.Intn(maxLength-minLength+1) + minLength

	b := make([]byte, stringLength)
	for i := range b {
		b[i] = def.CharSet[rand.Intn(len(def.CharSet))]
	}
	return string(b)
}
