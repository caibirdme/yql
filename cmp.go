package yql

import (
	"math"
	"sort"
	"strconv"
)

const (
	opEqual       = "="
	opNotEqual    = "!="
	opLarger      = ">"
	opLargerEqual = ">="
	opLess        = "<"
	opLessEqual   = "<="
	opInter       = "∩"
	opNotInter    = "!∩"
	opIn          = "in"
	opNotIn       = "!in"
)

const (
	epsilon = float64(1e-10)
)

func cmpInt(actual, expect int64, op string) bool {
	switch op {
	case opEqual,opIn:
		return actual == expect
	case opNotEqual,opNotIn:
		return actual != expect
	case opLarger:
		return actual > expect
	case opLargerEqual:
		return actual >= expect
	case opLess:
		return actual < expect
	case opLessEqual:
		return actual <= expect
	default:
		return false
	}
}

func cmpFloat(actual, expect float64, op string) bool {
	switch op {
	case opEqual:
		return floatEqual(actual, expect)
	case opNotEqual:
		return !floatEqual(actual, expect)
	case opLarger:
		return actual > expect
	case opLargerEqual:
		return !(actual < expect)
	case opLess:
		return actual < expect
	case opLessEqual:
		return floatLessEqual(actual, expect)
	default:
		return false
	}
}

func cmpStr(actual, expect string, op string) bool {
	switch op {
	case opEqual,opIn:
		return actual == expect
	case opNotEqual,opNotIn:
		return actual != expect
	case opLarger:
		return actual > expect
	case opLargerEqual:
		return actual >= expect
	case opLess:
		return actual < expect
	case opLessEqual:
		return actual <= expect
	default:
		return false
	}
}

func cmpBool(actual, expect bool, op string) bool {
	switch op {
	case opEqual:
		return actual == expect
	case opNotEqual:
		return actual != expect
	default:
		return false
	}
}

func compareSet(actual interface{}, expect []string, op string) bool {
	switch op {
	case opEqual, opNotEqual, opInter, opNotInter, opIn, opNotIn:
	default:
		return false
	}
	switch actualArr := actual.(type) {
	case int:
		return cmpIntSet([]int64{int64(actualArr)}, expect, op)
	case int64:
		return cmpIntSet([]int64{actualArr}, expect, op)
	case float64:
		return cmpFloatSet([]float64{actualArr}, expect, op)
	case string:
		return cmpStringSet([]string{actualArr}, expect, op)
	case []int:
		return cmpIntSet(intArr2i64Arr(actualArr), expect, op)
	case []int64:
		return cmpIntSet(actualArr, expect, op)
	case []float64:
		return cmpFloatSet(actualArr, expect, op)
	case []string:
		return cmpStringSet(actualArr, expect, op)
	default:
		return false
	}
}

func intArr2i64Arr(arr []int) []int64 {
	length := len(arr)
	if length == 0 {
		return nil
	}
	res := make([]int64, 0, length)
	for _, v := range arr {
		res = append(res, int64(v))
	}
	return res
}

var intSetCmpFunc = map[string]func([]int64, []int64) bool{
	opInter:    intSetInter,
	opNotInter: intSetNotInter,
	opIn:       intSetBelong,
	opNotIn:    intSetNotBelong,
}

func cmpIntSet(actualVals []int64, expectVals []string, op string) bool {
	expectArr := make([]int64, 0, len(expectVals))
	for _, expect := range expectVals {
		v, err := strconv.ParseInt(removeQuote(expect), 10, 64)
		if nil != err {
			return false
		}
		expectArr = append(expectArr, v)
	}
	cmp, ok := intSetCmpFunc[op]
	if ok {
		return cmp(actualVals, expectArr)
	}
	return false
}

var floatSetCmpFunc = map[string]func([]float64, []float64) bool{
	opInter:    floatSetInter,
	opNotInter: floatSetNotInter,
	opIn:       floatSetBelong,
	opNotIn:    floatSetNotBelong,
}

func cmpFloatSet(actualVals []float64, expectVals []string, op string) bool {
	expectArr := make([]float64, 0, len(expectVals))
	for _, expect := range expectVals {
		v, err := strconv.ParseFloat(removeQuote(expect), 64)
		if nil != err {
			return false
		}
		expectArr = append(expectArr, v)
	}
	cmp, ok := floatSetCmpFunc[op]
	if ok {
		return cmp(actualVals, expectArr)
	}
	return false
}

var stringSetCmpFunc = map[string]func([]string, []string) bool{
	opInter:    strSetInter,
	opNotInter: strSetNotInter,
	opIn:       strSetBelong,
	opNotIn:    strSetNotBelong,
}

func cmpStringSet(actual []string, expect []string, op string) bool {
	cmp, ok := stringSetCmpFunc[op]
	if !ok {
		return false
	}
	expectVals := make([]string, 0, len(expect))
	for _, v := range expect {
		expectVals = append(expectVals, removeQuote(v))
	}
	return cmp(actual, expectVals)
}

func intSetBelong(actualVals []int64, expectVals []int64) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || len(actualVals) > length {
		return false
	}
	expectArr := i64Arr(expectVals)
	sort.Sort(expectArr)
	for _, actual := range actualVals {
		t := sort.Search(length, func(i int) bool {
			return actual <= expectArr[i]
		})
		if t >= length || actual != expectArr[t] {
			return false
		}
	}
	return true
}

func intSetNotBelong(actualVals []int64, expectVals []int64) bool {
	return !intSetBelong(actualVals, expectVals)
}

func intSetInter(actualVals []int64, expectVals []int64) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || length == 0 {
		return false
	}
	expectArr := i64Arr(expectVals)
	sort.Sort(expectArr)
	for _, actual := range actualVals {
		t := sort.Search(length, func(i int) bool {
			return actual <= expectArr[i]
		})
		if t < length && actual == expectArr[t] {
			return true
		}
	}
	return false
}

func intSetNotInter(actualVals []int64, expectVals []int64) bool {
	return !intSetInter(actualVals, expectVals)
}

type i64Arr []int64

func (arr i64Arr) Len() int {
	return len(arr)
}

func (arr i64Arr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr i64Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func floatSetBelong(actualVals []float64, expectVals []float64) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || len(actualVals) > length {
		return false
	}
	sort.Float64s(expectVals)
	for _, actual := range actualVals {
		t := sort.Search(length, func(i int) bool {
			return floatLessEqual(actual, expectVals[i])
		})
		if t >= length || !floatEqual(actual, expectVals[t]) {
			return false
		}
	}
	return true
}

func floatSetNotBelong(actualVals []float64, expectVals []float64) bool {
	return !floatSetBelong(actualVals, expectVals)
}

func floatSetInter(actualVals []float64, expectVals []float64) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || length == 0 {
		return false
	}
	sort.Float64s(expectVals)
	for _, actual := range actualVals {
		t := sort.Search(length, func(i int) bool {
			return floatLessEqual(actual, expectVals[i])
		})
		if t < length && floatEqual(actual, expectVals[t]) {
			return true
		}
	}
	return false
}

func floatSetNotInter(actualVals []float64, expectVals []float64) bool {
	return !floatSetInter(actualVals, expectVals)
}

func strSetBelong(actualVals []string, expectVals []string) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || len(actualVals) > length {
		return false
	}
	sort.Strings(expectVals)
	for _, actual := range actualVals {
		t := sort.SearchStrings(expectVals, actual)
		if t >= length || actual != expectVals[t] {
			return false
		}
	}
	return true
}

func strSetNotBelong(actualVals []string, expectVals []string) bool {
	return !strSetBelong(actualVals, expectVals)
}

func strSetInter(actualVals []string, expectVals []string) bool {
	length := len(expectVals)
	if len(actualVals) == 0 || length == 0 {
		return false
	}
	sort.Strings(expectVals)
	for _, actual := range actualVals {
		t := sort.SearchStrings(expectVals, actual)
		if t < length && actual == expectVals[t] {
			return true
		}
	}
	return false
}

func strSetNotInter(actualVals []string, expectVals []string) bool {
	return !strSetInter(actualVals, expectVals)
}

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func floatLessEqual(a, b float64) bool {
	return a < b || floatEqual(a, b)
}
