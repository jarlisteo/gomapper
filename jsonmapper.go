package jsonmapper

import (
	"encoding/json"
	"strings"
)

var Data, Base string

type JsonMap map[string]string

func checkIsArray(i interface{}) bool {
	i2 := i
	var isSingle bool
	json.Unmarshal([]byte(Data), &i2)
	switch i.(type) {
	case []interface{}:
		isSingle = false
	case map[string]interface{}:
		isSingle = true
	}
	return isSingle
}

func getMainArrayMap() []interface{} {
	var interfaceArray []interface{}
	var baseMap map[string]interface{}
	json.Unmarshal([]byte(Data), &baseMap)
	if checkIsArray(baseMap[Base]) {
		interfaceArray = append(interfaceArray, baseMap[Base])
	} else {
		interfaceArray = baseMap[Base].([]interface{})
	}
	return interfaceArray
}

func getValue(i interface{}, v string) any {
	stringArray := strings.Split(v, ".")
	result := i.(map[string]interface{})[stringArray[0]]
	switch result.(type) {
	case map[string]interface{}:
		newString := append(stringArray[:0], stringArray[1:]...)
		return getValue(result, strings.Join(newString, "."))
	}
	return result
}

func Map[T any](data string, Map JsonMap) []T {
	Data = data
	Base = Map["base"]
	var Mapped []T
	for _, i := range getMainArrayMap() {
		newMap := make(map[string]any)
		for k, v := range Map {
			newMap[k] = getValue(i, v)
		}
		newObject := new(T)
		mapStringed, _ := json.Marshal(newMap)
		json.Unmarshal([]byte(mapStringed), &newObject)
		Mapped = append(Mapped, *newObject)
	}
	return Mapped
}
