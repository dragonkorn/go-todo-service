package utils

import (
	"encoding/json"
	"reflect"
)

// Struct  --> Map[string]interface{} Then MapStructure.Decode()  (Stuct A ---> Stuct B)
func Struct2Map(m interface{}) map[string]interface{} {
	mappedData := make(map[string]interface{})
	// fmt.Println(m)
	v := reflect.ValueOf(m)
	// fmt.Println("==============================")
	// fmt.Println(v)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		getField := v.Field(i)
		if getField.Kind() == reflect.Ptr && getField.IsNil() {
			mappedData[typeOfS.Field(i).Name] = nil
		} else {
			if getField.CanInterface() {
				mappedData[typeOfS.Field(i).Name] = GetPointerValue(getField.Interface())
			}
		}
	}

	return mappedData
}

func Json2Map(jsonByte []byte) (returnMap map[string]interface{}) {
	json.Unmarshal(jsonByte, &returnMap)
	return returnMap
}

func GetPointerValue(i interface{}) interface{} {

	valOf := reflect.ValueOf(i)
	var fieldValDeref reflect.Value

	if valOf.Kind() == reflect.Ptr {
		fieldValDeref = valOf.Elem()
	} else {
		fieldValDeref = valOf
	}
	return fieldValDeref.Interface()
}