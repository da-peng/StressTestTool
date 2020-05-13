package utils

import (
	"encoding/json"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
		return nil
	}
	return data
}

func StructToJson(obj interface{}) []byte {

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		Errors(err.Error())
		return nil
	}

	return jsonBytes
}

func MapToJson(mapInstances map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(mapInstances)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return jsonBytes
}

func JsonToStruct(jsonStr string, obj interface{}) interface{} {

	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return obj
}

func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return mapResult

}
