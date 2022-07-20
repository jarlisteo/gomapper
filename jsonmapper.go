package gomapper

import (
	"encoding/json"
	"strings"
)

type JsonMap map[string]string

type GoMapper struct {
	Ref  JsonMap
	Data string
	Base string
}

func New(ref JsonMap, data string) *GoMapper {
	return &GoMapper{Ref: ref, Data: data}
}

func (mapper *GoMapper) Map(object any) {
	var FinalMap []any

	mapper.setBase()

	//IF IS ARRAY DO MAPPING FOR EACH SUBGROUP
	if jsonIsArray(mapper.Data) {
		var iterableMap []any
		_ = json.Unmarshal([]byte(mapper.Data), &iterableMap)
		for _, subMap := range iterableMap {
			_ = MapValues(&subMap, mapper.Ref)
			FinalMap = append(FinalMap, subMap)
		}
	}

	//CONVERT TO INPUT OBJECT
	str, _ := json.Marshal(FinalMap)
	_ = json.Unmarshal(str, &object)
}

func (mapper *GoMapper) GoMainMap() {
	var mainObjects []any
	var baseMap map[string]any
	_ = json.Unmarshal([]byte(mapper.Data), &baseMap)
	if jsonIsArray(mapper.Data) {
		mainObjects = append(mainObjects, baseMap[mapper.Base])
	} else {
		mainObjects = baseMap[mapper.Base].([]any)
	}
	str, _ := json.Marshal(mainObjects)
	mapper.Data = string(str)
}

func MapValues(Map *any, ref JsonMap) error {
	mapCopy := *Map
	newMap := make(map[string]any)
	for value, key := range ref {
		newMap[value] = getValue(mapCopy, key)
	}
	*Map = newMap
	return nil
}

func getValue(i any, v string) any {
	stringArray := strings.Split(v, ".")
	result := i.(map[string]any)[stringArray[0]]
	switch result.(type) {
	case map[string]any:
		newString := append(stringArray[:0], stringArray[1:]...)
		return getValue(result, strings.Join(newString, "."))
	}
	return result
}

func (mapper *GoMapper) setBase() {
	if base, ok := mapper.Ref["base"]; ok {
		mapper.Base = base
		mapper.GoMainMap()
	}
}
