package tsp_test

import (
	"fmt"
	"testing"

	"github.com/nakario/tsp"
	"github.com/stretchr/testify/assert"
)

func TestPointArea(t *testing.T) {
	for _, v := range []struct{
		input tsp.Point
		expected tsp.Area
	}{
		{
			input: tsp.Point{X: 0, Y: 0},
			expected: tsp.AreaOnLine,
		},
		{
			input: tsp.Point{X: 1, Y: 1},
			expected: tsp.AreaOnLine,
		},
		{
			input: tsp.Point{X: -2, Y: 2},
			expected: tsp.AreaOnLine,
		},
		{
			input: tsp.Point{X: 0, Y: -3},
			expected: tsp.AreaOnLine,
		},
		{
			input: tsp.Point{X: 2, Y: 1},
			expected: tsp.AreaENE,
		},
		{
			input: tsp.Point{X: 1, Y: 2},
			expected: tsp.AreaNNE,
		},
		{
			input: tsp.Point{X: -1, Y: 2},
			expected: tsp.AreaNNW,
		},
		{
			input: tsp.Point{X: -2, Y: 1},
			expected: tsp.AreaWNW,
		},
		{
			input: tsp.Point{X: -2, Y: -1},
			expected: tsp.AreaWSW,
		},
		{
			input: tsp.Point{X: -1, Y: -2},
			expected: tsp.AreaSSW,
		},
		{
			input: tsp.Point{X: 1, Y: -2},
			expected: tsp.AreaSSE,
		},
		{
			input: tsp.Point{X: 2, Y: -1},
			expected: tsp.AreaESE,
		},
	} {
		t.Run(fmt.Sprintf("input: %v", v.input), func(t *testing.T) {
			actual := v.input.Area()
			assert.Equal(t, v.expected, actual)
		})
	}
}
