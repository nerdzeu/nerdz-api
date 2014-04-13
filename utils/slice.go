package utils

import (
	"fmt"
	"reflect"
)

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
			Type := values.Index(i).Type()
			elem := reflect.Indirect(reflect.New(Type))

			switch Type.Kind() {
			case reflect.Struct:
				for j := 0; j < Type.NumField(); j++ {
					// Copy each field
					elem.Field(j).Set(values.Index(i).Field(j))
				}
			default:
				elem.Set(values.Index(i))
			}
			reversedSlice.Index(k).Set(elem)
			k++
		}

		return reversedSlice.Interface()
	default:
		panic(fmt.Sprintf("Type not allowed: %v", reflect.TypeOf(slice).Kind()))
	}
}
