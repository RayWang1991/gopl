package format

import (
	"reflect"
	"strconv"
)

// Any formats any value as a string

func Any(v interface{}) string {
	return formatAtom(reflect.ValueOf(v))
}

func cf(r, i float64, bitSize int) string {
	if r == 0 && i == 0 {
		return "0"
	} else if r == 0 {
		return strconv.FormatFloat(i, 'g', 6, bitSize) + "i"
	} else if i == 0 {
		return strconv.FormatFloat(r, 'g', 6, bitSize)
	} else {
		return strconv.FormatFloat(r, 'g', 6, bitSize) + "+" +
			strconv.FormatFloat(i, 'g', 6, bitSize) + "i"
	}
}

// formatAtom formats a value without digging its internal structure
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'g', 6, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', 6, 64)
	case reflect.Complex64:
		c := v.Complex()
		r, i := real(c), imag(c)
		return cf(r, i, 32)
	case reflect.Complex128:
		c := v.Complex()
		r, i := real(c), imag(c)
		return cf(r, i, 64)
	case reflect.Array, reflect.Struct, reflect.Interface:
		return v.Type().String() + "value"
	case reflect.Map, reflect.Chan, reflect.Slice, reflect.Ptr, reflect.Func:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return "Unknown"
	}
}
