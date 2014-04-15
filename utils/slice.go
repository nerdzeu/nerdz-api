package utils

import (
	"fmt"
	"reflect"
)

func copyElement(Type reflect.Type, source reflect.Value) reflect.Value {

	elem := reflect.Indirect(reflect.New(Type))

	switch Type.Kind() {
	case reflect.Struct:
		for j := 0; j < Type.NumField(); j++ {
			elem.Field(j).Set(source.Field(j))
		}
	default:
		elem.Set(source)
	}

	return elem
}

func ReverseSlice(slice interface{}) interface{} {

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Ptr:

		values := reflect.Indirect(reflect.ValueOf(slice))
		if values.Len() == 0 {
			return slice
		}

		reversedSlice := reflect.MakeSlice(reflect.SliceOf(reflect.Indirect(values.Index(0)).Type()), values.Len(), values.Len())

		k := 0
		for i := values.Len() - 1; i >= 0; i-- {
			reversedSlice.Index(k).Set(copyElement(values.Index(i).Type(), values.Index(i)))
			k++
		}

		return reversedSlice.Interface()
	default:
		panic(fmt.Sprintf("Type not allowed: %v", reflect.TypeOf(slice).Kind()))
	}
}
