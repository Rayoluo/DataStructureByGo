package util

import "reflect"

// 比较两种空接口类型变量，如果a小于b则返回-1；如果a大于b则返回1；如果a等于b则返回0
func Compare(a, b interface{}) int {
	atype := reflect.TypeOf(a).String()
	btype := reflect.TypeOf(b).String()
	if atype != btype {
		panic("can't compare variables with different types")
	}
	switch a.(type) {
	case int:
		if a.(int) < b.(int) {
			return -1
		} else if a.(int) > b.(int) {
			return 1
		} else {
			return 0
		}
	case string:
		if a.(string) < b.(string) {
			return -1
		} else if a.(string) > b.(string) {
			return 1
		} else {
			return 0
		}
	case float64:
		if a.(float64) < b.(float64) {
			return -1
		} else if a.(float64) > b.(float64) {
			return 1
		} else {
			return 0
		}
	default:
		panic("unsupported type parameters")
	}

}
