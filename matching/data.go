package matching

import (
	"strings"
)

func DataMatching(data map[string]interface{}, matchingWords []string) map[string]interface{} {

	matchWords := make(map[string]bool)

	for _, v := range matchingWords {
		matchWords[v] = true
	}

	if len(matchWords) == 0 {
		return data
	}

	adjustedData := make(map[string]interface{})

	for k, v := range data {
		keySplits := strings.Split(k, "_")

		if _, present := matchWords[keySplits[0]]; present {
			adjustedData[k] = data[strings.Join(keySplits[1:], "_")]
		} else {
			adjustedData[k] = v
		}
	}

	return adjustedData
}
