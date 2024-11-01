package internal

import (
	"encoding/json"
	"log"
	"os"

	"github.com/haatos/goshipit/internal/model"
)

var ComponentCodeMap model.ComponentCodeMap
var ComponentExampleCodeMap model.ComponentExampleCodeMap

func init() {
	getComponentCodeMap()
	getComponentExampleCodeMap()
}

func getComponentCodeMap() {
	b, err := os.ReadFile("generated/component_code_map.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, &ComponentCodeMap); err != nil {
		log.Fatal(err)
	}

	for k := range ComponentCodeMap {
		for i := range ComponentCodeMap[k] {
			ComponentCodeMap[k][i].Label = SnakeCaseToCapitalized(ComponentCodeMap[k][i].Name)
		}
	}
}

func getComponentExampleCodeMap() {
	b, err := os.ReadFile("generated/component_example_code_map.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(b, &ComponentExampleCodeMap); err != nil {
		log.Fatal(err)
	}

	for k := range ComponentExampleCodeMap {
		for i := range ComponentExampleCodeMap[k] {
			ComponentExampleCodeMap[k][i].Label = SnakeCaseToCapitalized(ComponentExampleCodeMap[k][i].Name)
		}
	}
}

func SnakeCaseToCapitalized(s string) string {
	b := []byte(s)
	for i := range b {
		if i == 0 || (i > 0 && b[i-1] == ' ') {
			b[i] = b[i] - ('a' - 'A')
		}
		if b[i] == '_' {
			b[i] = ' '
		}
	}
	return string(b)
}
