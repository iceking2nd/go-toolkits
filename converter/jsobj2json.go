package converter

import (
	"encoding/json"

	"github.com/dop251/goja"
)

func JSObj2JSON(obj string) ([]byte, error) {
	vm := goja.New()

	v, err := vm.RunString("(" + obj + ")")
	if err != nil {
		return nil, err
	}

	exported := v.Export()

	return json.Marshal(exported)
}
