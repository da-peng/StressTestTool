package utils

import (
	"encoding/json"
	"reflect"
)

// StructToMap 结构体 实例 转为Map
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

// StructToJSON 结构体 实例 转为Json
func StructToJSON(obj interface{}) []byte {

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		Errors(err.Error())
		return nil
	}

	return jsonBytes
}

// MapToJSON Map数据结构转为JSON
func MapToJSON(mapInstances map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(mapInstances)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return jsonBytes
}

// JSONToStruct JSON数据结构转为结构体
func JSONToStruct(jsonStr string, obj interface{}) interface{} {

	err := json.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return obj
}

// JSONToMap JSON转为Map数据结构
func JSONToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		Errors(err.Error())
		return nil
	}
	return mapResult
}
