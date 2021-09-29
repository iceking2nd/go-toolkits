package slice

import (
	"reflect"
)

func Diff(s1 interface{}, s2 interface{}) (interface{}, interface{}) {
	if reflect.ValueOf(s1).Kind() != reflect.Slice || reflect.ValueOf(s2).Kind() != reflect.Slice {
		return nil, nil
	}
	if reflect.TypeOf(s1) != reflect.TypeOf(s2) {
		return nil, nil
	}
	ds1 := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(s1).Elem()), 0, 0)
	ds2 := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(s1).Elem()), 0, 0)
	t1 := make(map[interface{}]interface{})
	t2 := make(map[interface{}]interface{})
	for i := 0; i < reflect.ValueOf(s1).Len(); i++ {
		t1[reflect.ValueOf(s1).Index(i).Interface()] = reflect.ValueOf(s1).Index(i).Interface()
	}

	reflectionValue2 := reflect.New(ds2.Type())
	reflectionValue2.Elem().Set(ds2)
	slicePtr2 := reflect.ValueOf(reflectionValue2.Interface())
	sliceValuePtr2 := slicePtr2.Elem()
	for i := 0; i < reflect.ValueOf(s2).Len(); i++ {
		t2[reflect.ValueOf(s2).Index(i).Interface()] = reflect.ValueOf(s2).Index(i).Interface()
		if _, in := t1[reflect.ValueOf(s2).Index(i).Interface()]; !in {
			sliceValuePtr2.Set(reflect.Append(sliceValuePtr2, reflect.ValueOf(s2).Index(i)))
		}
	}

	reflectionValue1 := reflect.New(ds1.Type())
	reflectionValue1.Elem().Set(ds1)
	slicePtr1 := reflect.ValueOf(reflectionValue1.Interface())
	sliceValuePtr1 := slicePtr1.Elem()
	for i := 0; i < reflect.ValueOf(s1).Len(); i++ {
		if _, in := t2[reflect.ValueOf(s1).Index(i).Interface()]; !in {
			sliceValuePtr1.Set(reflect.Append(sliceValuePtr1, reflect.ValueOf(s1).Index(i)))
		}
	}
	return sliceValuePtr1.Interface(), sliceValuePtr2.Interface()
}
