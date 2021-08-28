package flatten

import (
	"reflect"
)

func Flatten(acc *[]interface{}, val interface{}) {
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			v := rv.Index(i).Interface()
			Flatten(acc, v)
		}
	} else {
		*acc = append(*acc, val)
	}
}
