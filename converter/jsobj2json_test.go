package converter

import (
	"slices"
	"testing"
)

func TestJSObj2JSON(t *testing.T) {
	obj := `{a:'a', b:10, c:false, d:3.14, e:null}`
	objJSON := `{"a":"a","b":10,"c":false,"d":3.14,"e":null}`
	jsonBytes, err := JSObj2JSON(obj)
	if err != nil {
		t.Errorf("%v", err)
	}
	if slices.Compare(jsonBytes, []byte(objJSON)) != 0 {
		t.Errorf(`"%s" was not expect value`, jsonBytes)
	}
}
