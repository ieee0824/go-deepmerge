package deepmerge

import (
	"reflect"
	"errors"
)

var (
	TypeNotMatchErr = errors.New("type not match")
)

func convertSlice(i interface{}) []interface{} {
	ret := []interface{}{}


	switch i.(type) {
	case []interface{}:
		return i.([]interface{})
	case []string:
		for _, v := range i.([]string) {
			ret = append(ret, v)
		}
		return ret
	case []int:
		for _, v := range i.([]int) {
			ret = append(ret, v)
		}
		return ret
	case []float64:
		for _, v := range i.([]float64) {
			ret = append(ret, v)
		}
		return ret
	case []float32:
		for _, v := range i.([]float32) {
			ret = append(ret, v)
		}
		return ret
	case []byte:
		return append(ret, i)
	}
	return nil
}

func Merge(src, dst interface{}) (interface{}, error) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType.Kind() != dstType.Kind() {
		return nil, TypeNotMatchErr
	}

	switch srcType.Kind() {
	case reflect.Map:
		srcMap := src.(map[string]interface{})
		for k, dstVal := range dst.(map[string]interface{}) {
			srcVal, ok := srcMap[k]
			if !ok {
				srcMap[k] = dstVal
			} else {
				mergedVal, err := Merge(srcVal, dstVal)
				if err != nil {
					return nil, err
				}
				srcMap[k] = mergedVal
			}
		}
		return src, nil
	case reflect.Slice:
		//return append(src.([]interface{}), dst.([]interface{})...), nil
		srcSlice := convertSlice(src)
		dstSlice := convertSlice(dst)
		return append(srcSlice, dstSlice...), nil
	default:
		return src, nil
	}
}
