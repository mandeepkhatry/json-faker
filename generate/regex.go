package generate

import (
	"github.com/lucasjones/reggen"
)

func GenerateRegex(regex string) string {
	str, _ := reggen.Generate(regex, 1)
	return str
}
