package generate

import (
	"faker/def"
	"math/rand"
	"time"
)

func GenerateInteger(properties map[string]interface{}) int {

	if _, present := properties["enum"]; present {
		if len(properties["enum"].([]int)) != 0 {
			return properties["enum"].([]int)[0]
		}
	}

	if _, present := properties["constant"]; present {
		return int(properties["constant"].(float64))

	}

	rand.Seed(time.Now().UnixNano())

	var min, max int

	if _, present := properties["exclusiveMaximum"]; present {
		max = int(properties["exclusiveMaximum"].(float64)) - 1
	}

	if _, present := properties["exclusiveMinimum"]; present {
		min = int(properties["exclusiveMinimum"].(float64)) + 1
	}

	if _, present := properties["maximum"]; !present {
		max = def.DummyDataRange["integer"].(map[string]int)["maximum"]
	} else {
		max = int(properties["maximum"].(float64))
	}

	if _, present := properties["minimum"]; !present {
		min = def.DummyDataRange["integer"].(map[string]int)["minimum"]
	} else {
		min = int(properties["minimum"].(float64))
	}

	if _, present := properties["enum"]; present {
		enum := properties["enum"].([]int)
		if len(enum) > 0 {
			return enum[0]
		}
		return 0
	}

	if _, present := properties["multipleOf"]; present {
		multipleOf := int(properties["multipleOf"].(float64))
		i := 0
		for {
			eachMultiple := multipleOf * i
			if eachMultiple >= min && eachMultiple <= max {
				return multipleOf * i
			} else if eachMultiple > max {
				return 0
			}
			i++

		}
	}

	return rand.Intn(max-min+1) + min
}

func GenerateFloat(properties map[string]interface{}) float64 {

	if _, present := properties["enum"]; present {
		if len(properties["enum"].([]float64)) != 0 {
			return properties["enum"].([]float64)[0]
		}

	}

	if _, present := properties["constant"]; present {
		return properties["constant"].(float64)

	}

	rand.Seed(time.Now().UnixNano())

	var min, max float64

	if _, present := properties["exclusiveMaximum"]; present {
		max = properties["exclusiveMaximum"].(float64) - 0.5
	}

	if _, present := properties["exclusiveMinimum"]; present {
		min = properties["exclusiveMinimum"].(float64) + 0.5
	}

	if _, present := properties["maximum"]; !present {
		max = def.DummyDataRange["number"].(map[string]float64)["maximum"]
	} else {
		max = properties["maximum"].(float64)
	}

	if _, present := properties["minimum"]; !present {
		min = def.DummyDataRange["number"].(map[string]float64)["minimum"]
	} else {
		min = properties["minimum"].(float64)
	}

	if _, present := properties["enum"]; present {
		enum := properties["enum"].([]float64)
		if len(enum) > 0 {
			return enum[0]
		}
		return 0.0
	}

	if _, present := properties["multipleOf"]; present {
		multipleOf := properties["multipleOf"].(float64)
		i := 0.0
		for {
			eachMultiple := multipleOf * i
			if eachMultiple >= min && eachMultiple <= max {
				return multipleOf * i
			} else if eachMultiple > max {
				return 0.0
			}

		}
	}

	return rand.Float64()*(max-min) + min
}
