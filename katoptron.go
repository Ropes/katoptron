package katoptron

import (
	"fmt"
	"reflect"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func Any(v interface{}) string {
	return formatAtom(reflect.ValueOf(v))
}

// formatAtom formats any reflect Value into a printable string
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid reflect value"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'E', -1, 10)
	case reflect.Complex64, reflect.Complex128:
		//TODO: Figure out complex variables
		return "PUNTing on Complex variables"
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	log.Infof("Katoptron: %s (%T)", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		log.Warnf("Invalid Reflection found %s", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, key), v.MapIndex(key))
		}
	case reflect.Interface:
		if v.IsNil() {
			log.Warnf("%s == nil", path)
		} else {
			log.Infof("%s.type = %s", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		log.Infof("%s == %v", path, v)
	}
}
