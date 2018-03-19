package lambda

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestFilter_Int(t *testing.T) {
	var testData = []struct {
		expr   string
		dst    []int
		expect []int
	}{
		{
			expr:   `(p) =>  p % 2 == 1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 3, 5, 7},
		},
		{
			expr:   `(v) =>  (v&1) == 1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 3, 5, 7},
		},
		{
			expr:   `(v) =>  v&1 == 0`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{2, 4, 6},
		},
		{
			expr:   `(v) =>  v+v<= 6 `,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 2, 3},
		},
		{
			expr:   `(v) =>  (v<<2)>>1 == 8`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{4},
		},
		{
			expr:   `(v) =>  v%2 == 0`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{2, 4, 6},
		},
		{
			expr:   `(v) =>  v > 5`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{6, 7},
		},
		{
			expr:   `(v) =>  v > 5+1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{7},
		},
		{
			expr:   `(v) =>  v*2 == v+3`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{3},
		},
		{
			expr:   `(v) =>  v > 1+2+3/(0+1)`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{7},
		},
		{
			expr:   `(v) =>  v <= 1+2+3/(0+1)`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.expr)
		ass.NoError(mstate.err, "%s", tc.expr)
		r := mstate.Call(tc.dst)
		inf, err := r.Interface()
		ass.NoError(err, "%s", tc.expr)
		ans, ok := inf.([]int)
		ass.True(ok)
		ass.Equal(tc.expect, ans, "expr=%s", tc.expr)
	}
}

func TestFilter_Float64(t *testing.T) {
	var testData = []struct {
		expr   string
		dst    []float64
		expect []float64
	}{
		{
			expr:   `(v) =>  v*2 == v+3`,
			dst:    []float64{1, 2, 3, 4, 5, 6, 7},
			expect: []float64{3},
		},
		{
			expr:   `(v) =>  v*2 >= 5`,
			dst:    []float64{1, 2, 3, 4, 5, 6, 7},
			expect: []float64{3, 4, 5, 6, 7},
		},
		{
			expr:   `(v) =>  v*2 >= 5+1+1`,
			dst:    []float64{1, 2, 3, 4, 5, 6, 7},
			expect: []float64{4, 5, 6, 7},
		},
		{
			expr:   `(v) =>  v*2 >= 5+1+1*((2+1*10000)*0+1)`,
			dst:    []float64{1, 2, 3, 4, 5, 6, 7},
			expect: []float64{4, 5, 6, 7},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.expr)
		ass.NoError(mstate.err, "%s", tc.expr)
		r := mstate.Call(tc.dst)
		inf, err := r.Interface()
		ass.NoError(err, "%s", tc.expr)
		ans, ok := inf.([]float64)
		ass.True(ok)
		ass.Equal(tc.expect, ans, "expr=%s", tc.expr)
	}
}

type Student struct {
	Age  int
	Name string
}

func TestFilter_Struct(t *testing.T) {
	var students = []Student{
		Student{
			Name: "deen",
			Age:  24,
		},
		Student{
			Name: "bob",
			Age:  22,
		},
		Student{
			Name: "alice",
			Age:  23,
		},
		Student{
			Name: "tom",
			Age:  25,
		},
		Student{
			Name: "jerry",
			Age:  20,
		},
	}
	var testData = []struct {
		expr   string
		expect []int
	}{
		{
			expr:   `(v) => v.Age+2 > 24+1`,
			expect: []int{0, 3},
		},
		{
			expr:   `(v) => v.Age >= 23`,
			expect: []int{0, 2, 3},
		},
		{
			expr:   `(v) => v.Age < 23`,
			expect: []int{1, 4},
		},
		{
			expr:   `(v) => v.Age < 23 || v.Name == "tom"`,
			expect: []int{1, 3, 4},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.expr)
		ass.NoError(mstate.err, "%s", tc.expr)
		r := mstate.Call(students)
		inf, err := r.Interface()
		ass.NoError(err, "%s", tc.expr)
		ans, ok := inf.([]Student)
		ass.True(ok)
		var expectArr []Student
		for _, idx := range tc.expect {
			expectArr = append(expectArr, students[idx])
		}
		ass.Equal(expectArr, ans, "expr=%s", tc.expr)
	}
}

func TestFilter_Pointer(t *testing.T) {
	var students = []*Student{
		&Student{
			Name: "deen",
			Age:  24,
		},
		&Student{
			Name: "bob",
			Age:  22,
		},
		&Student{
			Name: "alice",
			Age:  23,
		},
		&Student{
			Name: "tom",
			Age:  25,
		},
		&Student{
			Name: "jerry",
			Age:  20,
		},
	}
	var testData = []struct {
		expr   string
		expect []int
	}{
		{
			expr:   `(v) => v.Age+2 > 24+1`,
			expect: []int{0, 3},
		},
		{
			expr:   `(v) => v.Age >= 23`,
			expect: []int{0, 2, 3},
		},
		{
			expr:   `(v) => v.Age < 23`,
			expect: []int{1, 4},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.expr)
		ass.NoError(mstate.err, "%s", tc.expr)
		r := mstate.Call(students)
		inf, err := r.Interface()
		ass.NoError(err, "%s", tc.expr)
		ans, ok := inf.([]*Student)
		ass.True(ok)
		var expectArr []*Student
		for _, idx := range tc.expect {
			expectArr = append(expectArr, students[idx])
		}
		ass.Equal(expectArr, ans, "expr=%s", tc.expr)
	}
}

func TestMap_Int(t *testing.T) {
	var testData = []struct {
		expr   string
		dst    []int
		expect []int
	}{
		{
			expr:   `(p) =>  p +1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{2, 3, 4, 5, 6, 7, 8},
		},
		{
			expr:   `(v) =>  v & 1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 0, 1, 0, 1, 0, 1},
		},
		{
			expr:   `(v) =>  v | 1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{1, 3, 3, 5, 5, 7, 7},
		},
		{
			expr:   `(v) =>  (v - 1)*3`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{0, 3, 6, 9, 12, 15, 18},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Map(tc.expr)
		ass.NoError(mstate.err, "%s", tc.expr)
		r := mstate.Call(tc.dst)
		inf, err := r.Interface()
		ass.NoError(err, "%s", tc.expr)
		ans, ok := inf.([]int)
		ass.True(ok)
		ass.Equal(tc.expect, ans, "expr=%s", tc.expr)
	}
}

func TestChainable(t *testing.T) {
	var testData = []struct {
		m      string
		f      string
		dst    []int
		expect []int
	}{
		{
			m:      `(p) =>  p +1`,
			f:      `(v) => v&1==1`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{2, 4, 6, 8},
		},
		{
			m:      `(v) =>  v << 2`,
			f:      `(v) => v > 3 && v <= 7`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{16, 20, 24, 28},
		},
		{
			m:      `(v) =>  (v - 1)*3`,
			f:      `(v) => v<4`,
			dst:    []int{1, 2, 3, 4, 5, 6, 7},
			expect: []int{0, 3, 6},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.f).Map(tc.m)
		ass.NoError(mstate.err, "filter:%s\r\nmap:%s\r\n", tc.f, tc.m)
		r := mstate.Call(tc.dst)
		inf, err := r.Interface()
		ass.NoError(err)
		ass.NoError(mstate.err, "filter:%s\r\nmap:%s\r\n", tc.f, tc.m)
		ans, ok := inf.([]int)
		ass.True(ok)
		ass.Equal(tc.expect, ans, "filter:%s\r\nmap:%s\r\n", tc.f, tc.m)
	}
}

func TestThreeChain(t *testing.T) {
	scores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//                                   syntax error
	r := Filter(`(v) => v % 2 == 0`).Map(`v => v*v`).Call(scores)
	s, err := r.Interface()
	ass := assert.New(t)
	ass.Error(err)
	dst := []int{1, 2, 3, 4, 5, 6, 7}
	r = Filter(`(v) => v > 3 && v <= 7`).Map(`(v) =>  v << 2`).Filter(`(v) => v % 8 == 0`).Call(dst)
	s, err = r.Interface()
	ass.NoError(err)
	ass.Equal([]int{16, 24}, s)
}
