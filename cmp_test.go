package yql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSetBelong(t *testing.T) {
	var testData = []struct {
		ina, inb []int64
		out      bool
	}{
		{
			ina: []int64{3, 1, 2},
			inb: []int64{5, 1, 4, 3, 2},
			out: true,
		},
		{
			ina: []int64{3, 1, 2},
			inb: []int64{5, 1, 4, 3, 0},
			out: false,
		},
		{
			ina: []int64{3, 1, 2, 6, 8, 10, 11},
			inb: []int64{5, 1, 4, 3, 0},
			out: false,
		},
		{
			ina: []int64{},
			inb: []int64{1},
			out: false,
		},
		{
			ina: []int64{1},
			inb: []int64{1},
			out: true,
		},
		{
			ina: []int64{},
			inb: []int64{},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, intSetBelong(tc.ina, tc.inb), "ina=%+v||inb=%+v", tc.ina, tc.inb)
	}
}

func TestIntSetInter(t *testing.T) {
	var testData = []struct {
		ina, inb []int64
		out      bool
	}{
		{
			ina: []int64{3, 1, 2},
			inb: []int64{5, 1, 4, 3, 2},
			out: true,
		},
		{
			ina: []int64{3, 1, 2},
			inb: []int64{5, 1, 4, 3, 0},
			out: true,
		},
		{
			ina: []int64{3, 1, 2, 6, 8, 10, 11},
			inb: []int64{5, 1, 4, 3, 0},
			out: true,
		},
		{
			ina: []int64{},
			inb: []int64{1},
			out: false,
		},
		{
			ina: []int64{1},
			inb: []int64{1},
			out: true,
		},
		{
			ina: []int64{},
			inb: []int64{},
			out: false,
		},
		{
			ina: []int64{1, 3, 5, 7, 9},
			inb: []int64{2, 4, 6, 8, 10, 5},
			out: true,
		},
		{
			ina: []int64{1, 3, 5, 7, 9},
			inb: []int64{2, 4, 6, 8, 10, 12},
			out: false,
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, intSetInter(tc.ina, tc.inb), "ina=%+v||inb=%+v", tc.ina, tc.inb)
	}
}
