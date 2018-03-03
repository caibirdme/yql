package lambda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var testData = []struct {
		expr   string
		dst    []int
		expect []int
	}{
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
	}
	ass := assert.New(t)
	for _, tc := range testData {
		mstate := Filter(tc.expr)
		if !ass.NoError(mstate.err) {
			t.FailNow()
		}
		result := mstate.Call(tc.dst)
		if !ass.NoError(result.err) {
			t.FailNow()
		}
		inf, err := result.Interface()
		if !ass.NoError(err) {
			t.FailNow()
		}
		ans, ok := inf.([]int)
		if !ass.True(ok) {
			t.FailNow()
		}
		ass.Equal(tc.expect, ans, "expr=%s", tc.expr)
	}
}
