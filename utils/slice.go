package utils

import (
	"fmt"
	"reflect"
)

// copyElement returns the reflect.Value, (deep) copy of source
func copyElement(source reflect.Value) reflect.Value {

	Type := source.Type()
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

// ReverseSlice reverse slice if slice is of type reflect.Slice or reflrect.Ptr (to silce)
// panics if slice type is different
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
			reversedSlice.Index(k).Set(copyElement(values.Index(i)))
			k++
		}

		return reversedSlice.Interface()
	default:
		panic(fmt.Sprintf("Type not allowed: %v", reflect.TypeOf(slice).Kind()))
	}
}
