package utils

import (
	"faker/def"
	"strings"
)

func GetReference(ref string, definations map[string]interface{}) (map[string]interface{}, error) {
	refSplit := strings.Split(ref, "/")

	lookups := make([]string, 0)

	if refSplit[0] == "#" && refSplit[1] == "definitions" {
		lookups = append(lookups, refSplit[2:]...)
	} else {
		lookups = append(lookups, refSplit...)
	}

	nextDefination := definations

	for _, eachLookUpKey := range lookups {

		if _, present := nextDefination[eachLookUpKey]; present {
			nextDefination = nextDefination[eachLookUpKey].(map[string]interface{})
		} else {
			return make(map[string]interface{}), def.ErrInvalidReference
		}
	}

	return nextDefination, nil

}
