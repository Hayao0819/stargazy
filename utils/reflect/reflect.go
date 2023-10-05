package reflectutils

import "reflect"

func GetFields(value any) []*reflect.StructField {
	rv := reflect.ValueOf(value)
	rt := rv.Type()
	arr := []*reflect.StructField{}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		arr = append(arr, &field)
	}
	return arr
}
