package gomapper

import (
	"encoding/json"
	"strings"
)

type JsonMap map[string]string

func Map(refObject any, refMap JsonMap, input string) {
	if invalidInput(input) {
		return
	}

	if refMapHaveBase(refMap) {
		input = goMainMap(input, refMap["base"])
	}

	if jsonIsArray(input) {
		input = mapArray(input, refMap)
	} else {
		input = mapSingle(input, refMap)
	}

	//OUT FINAL OBJECT
	_ = json.Unmarshal([]byte(input), &refObject)
}

func invalidInput(input string) bool {
	if input == "" || input == "[]" {
		return true
	}
	return false
}

func refMapHaveBase(refMap JsonMap) bool {
	if _, ok := refMap["base"]; ok {
		return true
	}
	return false
}

func mapArray(input string, refMap JsonMap) string {
	var finalMap []any
	var iterableMap []any
	_ = json.Unmarshal([]byte(input), &iterableMap)
	for _, subMap := range iterableMap {
		mapValues(subMap.(map[string]any), refMap)
		finalMap = append(finalMap, subMap)
	}
	str, _ := json.Marshal(finalMap)
	return string(str)
}

func mapSingle(input string, refMap JsonMap) string {
	var FinalMap map[string]any
	_ = json.Unmarshal([]byte(input), &FinalMap)
	mapValues(FinalMap, refMap)
	str, _ := json.Marshal(FinalMap)
	return string(str)
}

func goMainMap(jsonInput string, base string) string {
	var str []byte
	var baseMap map[string]any
	_ = json.Unmarshal([]byte(jsonInput), &baseMap)
	if jsonIsArray(jsonInput) {
		var mainObjects []any
		mainObjects = append(mainObjects, baseMap[base])
		str, _ = json.Marshal(mainObjects)
	} else {
		str, _ = json.Marshal(baseMap[base])
	}
	return string(str)
}

func mapValues(inputMap map[string]any, ref JsonMap) {
	for value, key := range ref {
		inputMap[value] = getValue(inputMap, key)
	}
}

func getValue(inputMap map[string]any, v string) any {
	stringArray := strings.Split(v, ".")
	result := inputMap[stringArray[0]]
	switch r := result.(type) {
	case map[string]any:
		newString := append(stringArray[:0], stringArray[1:]...)
		return getValue(r, strings.Join(newString, "."))
	}
	return result
}
