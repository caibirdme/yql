package yql

import (
	"fmt"
	"reflect"
)

var supportFuncHandler = map[string]func(interface{}) interface{}{
	"count": countRunner,
	"sum":   sumRunner,
	"avg":   avgRunner,
	"max":   maxRunner,
	"min":   minRunner,
}

func errFuncNotSupport(funcName string) error {
	return fmt.Errorf("%s isn't supported", funcName)
}

func funcRuner(actual interface{}, funcs []string) (interface{}, error) {
	if 0 == len(funcs) {
		return actual, nil
	}
	for _, fName := range funcs {
		fn, ok := supportFuncHandler[fName]
		if !ok {
			return nil, errFuncNotSupport(fName)
		}
		actual = fn(actual)
	}
	return actual, nil
}

func countRunner(in interface{}) interface{} {
	arr := reflect.ValueOf(in)
	if arr.Type().Kind() == reflect.Slice {
		return arr.Len()
	}
	return 1
}

func sumRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []int:
		return sumInt(t)
	case []int64:
		return sumInt64(t)
	case []float64:
		return sumFloat64(t)
	default:
		return in
	}
}

func sumInt(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}

func sumInt64(arr []int64) int64 {
	var res int64
	for _, v := range arr {
		res += v
	}
	return res
}

func sumFloat64(arr []float64) float64 {
	var res float64
	for _, v := range arr {
		res += v
	}
	return res
}

func maxRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []int:
		return maxInt(t)
	case []int64:
		return maxInt64(t)
	case []float64:
		return maxFloat64(t)
	default:
		return in
	}
}

func maxInt(arr []int) int {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t > res {
			res = t
		}
	}
	return res
}

func maxInt64(arr []int64) int64 {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t > res {
			res = t
		}
	}
	return res
}

func maxFloat64(arr []float64) float64 {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t > res {
			res = t
		}
	}
	return res
}

func minRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []int:
		return minInt(t)
	case []int64:
		return minInt64(t)
	case []float64:
		return minFloat64(t)
	default:
		return in
	}
}

func minInt(arr []int) int {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t < res {
			res = t
		}
	}
	return res
}

func minInt64(arr []int64) int64 {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t < res {
			res = t
		}
	}
	return res
}

func minFloat64(arr []float64) float64 {
	var res = arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		t := arr[i]
		if t < res {
			res = t
		}
	}
	return res
}

func avgRunner(in interface{}) interface{} {
	switch t := in.(type) {
	case []int:
		return avgInt(t)
	case []int64:
		return avgInt64(t)
	case []float64:
		return avgFloat64(t)
	default:
		return in
	}
}

func avgInt(arr []int) float64 {
	var total int
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func avgInt64(arr []int64) float64 {
	var total int64
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func avgFloat64(arr []float64) float64 {
	var total float64
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}
