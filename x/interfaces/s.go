package interfaces

import (
	"math"
	"reflect"
)

func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}
	_v := reflect.ValueOf(v)
	switch _v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if _v.IsNil() {
			return true
		}
	}
	return isZero(_v)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !IsZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Slice:
		return v.Len() == 0
	case reflect.Ptr:
		return IsZero(v.Elem())
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !IsZero(v.Field(i)) {
				return false
			}
		}
		return true
	//case reflect.Invalid:
	//	return true
	default: // reflect.Invalid, Complex64, reflect.Complex128;
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}
