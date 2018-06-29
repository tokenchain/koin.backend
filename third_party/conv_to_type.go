package third

import (
	"reflect"
	"strconv"
)

// SetFieldTo set adequately with the type of the field the value
// 'toSet'.
func SetFieldTo(field reflect.Value, toSet string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(toSet)
	case reflect.Int:
		if i, err := strconv.ParseInt(toSet, 10, 32); err == nil {
			field.SetInt(i)
		}
	case reflect.Int32:
		if i, err := strconv.ParseInt(toSet, 10, 32); err == nil {
			field.SetInt(i)
		}
	case reflect.Int16:
		if i, err := strconv.ParseInt(toSet, 10, 16); err == nil {
			field.SetInt(i)
		}
	case reflect.Int64:
		if i, err := strconv.ParseInt(toSet, 10, 64); err == nil {
			field.SetInt(i)
		}
	case reflect.Int8:
		if i, err := strconv.ParseInt(toSet, 10, 8); err == nil {
			field.SetInt(i)
		}
	case reflect.Uint:
		if i, err := strconv.ParseUint(toSet, 10, 64); err == nil {
			field.SetUint(i)
		}
	case reflect.Uint64:
		if i, err := strconv.ParseUint(toSet, 10, 64); err == nil {
			field.SetUint(i)
		}
	case reflect.Uint32:
		if i, err := strconv.ParseUint(toSet, 10, 32); err == nil {
			field.SetUint(i)
		}
	case reflect.Uint16:
		if i, err := strconv.ParseUint(toSet, 10, 16); err == nil {
			field.SetUint(i)
		}
	case reflect.Uint8:
		if i, err := strconv.ParseUint(toSet, 10, 8); err == nil {
			field.SetUint(i)
		}
	case reflect.Bool:
		if i, err := strconv.ParseBool(toSet); err == nil {
			field.SetBool(i)
		}
	case reflect.Float64:
		if i, err := strconv.ParseFloat(toSet, 64); err == nil {
			field.SetFloat(i)
		}
	case reflect.Float32:
		if i, err := strconv.ParseFloat(toSet, 32); err == nil {
			field.SetFloat(i)
		}
	}
}