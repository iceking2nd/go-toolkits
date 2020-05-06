package strings

import (
	"fmt"
	"reflect"
)

func GetNestedMapNameWithDot(m map[string]interface{}, prefix string) []string {
	var nestedName = []string{}
	for k, v := range m {

		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			nestedName = append(nestedName, GetNestedMapNameWithDot(v.(map[string]interface{}), fmt.Sprintf("%s%s.", prefix, k))...)
		default:
			nestedName = append(nestedName, prefix+k)
		}
	}
	return nestedName
}
