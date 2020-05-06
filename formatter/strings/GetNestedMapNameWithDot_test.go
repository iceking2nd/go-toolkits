package strings

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetNestedMapNameWithDot(t *testing.T) {
	var (
		mapTables = []struct {
			in       string
			expected []string
		}{
			{`{"root": {"children1": {"children2": ""}}}`, []string{"root.children1.children2"}},
			{`{"root": {"children": ""}}`, []string{"root.children"}},
			{`{"root": ""}`, []string{"root"}},
		}
	)

	for _, mt := range mapTables {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(mt.in), &data)
		actual := GetNestedMapNameWithDot(data, "")
		if !reflect.DeepEqual(actual, mt.expected) {
			t.Errorf("GetNestedMapNameWithDot(%v) = %s, expected %s", mt.in, actual, mt.expected)
		}
	}
}
