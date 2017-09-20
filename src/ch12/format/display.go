package format

import (
	"reflect"
	"fmt"
)

func Display(name string, v interface{}) {
	display(name, reflect.ValueOf(v))
}

func display(path string, v reflect.Value) {
	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(k)), v.MapIndex(k))
		}
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i ++ {
			display(fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name), v.Field(i))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("*%s", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(fmt.Sprintf("%s.value", path), v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
