package utils

import (
	"reflect"
	"strings"
)

/* 对象转化为 map
params:
    obj: 需要转化的对象
    all: 是否所有的字段都需要转化
        true: 所有字段都转化
        false: 如果为nil, "", 0 的字符串不转
*/
func Obj2Map(obj interface{}, all bool) map[string]interface{} {
	v := reflect.ValueOf(obj)

	typ := v.Type()
	switch typ.Kind() {
	case reflect.Ptr:
		return struct2Map(v.Elem().Interface(), all)
	case reflect.Struct:
		return struct2Map(obj, all)
	}

	return make(map[string]interface{})

}

func struct2Map(obj interface{}, all bool) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		// 获取key
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = strings.ToLower(t.Field(i).Name)
		}
		// 获取value
		value := v.Field(i).Interface()
		if !all { // 不是所有字段都需要需要判断每个类型, 并且确定是否需要
			if !haveValue(value) {
				continue
			}
		}

		data[key] = value
	}
	return data
}

func haveValue(value interface{}) bool {
	v := reflect.ValueOf(value)
	typ := v.Type()
	switch typ.Kind() {
	case reflect.String:
		return !(v.Len() == 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return !(v.Int() == 0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return !(v.Uint() == 0)
	case reflect.Float32, reflect.Float64:
		return !(v.Float() == 0)
	case reflect.Interface, reflect.Ptr:
		return !v.IsNil()
	}

	return true
}
