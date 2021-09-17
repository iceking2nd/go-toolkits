package slice

import (
	"reflect"
	"strings"
)

func Contains(s interface{}, i interface{}) bool {
	if reflect.ValueOf(s).Kind().String() == "slice" {
		if reflect.ValueOf(i).Kind().String() != "slice" {
			if strings.Trim(reflect.TypeOf(s).String(), "[]") == reflect.TypeOf(i).String() {
				var t = map[interface{}]interface{}{}
				for x := 0; x < reflect.ValueOf(s).Len(); x++ {
					t[reflect.ValueOf(s).Index(x).Interface()] = nil
				}
				_, ok := t[reflect.ValueOf(i).Interface()]
				return ok
			}
		}
	}
	return false
}
