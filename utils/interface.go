package utils

import "reflect"

func InterfaceIsNil(data interface{}) bool {
	if data == nil || (reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) {
		return true
	}
	return false
}
