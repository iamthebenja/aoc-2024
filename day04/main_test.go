package main

import "testing"

func TestCountXMAS(t *testing.T) {
	var tests = []struct {
		input    [][]rune
		expected int
	}{
		{[][]rune{
			{'X', 'X', 'M', 'S'},
		}, 0},
		{[][]rune{
			{'X', 'M', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'S', 'X', 'M', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'X', 'A', 'A', 'S'},
			{'M', 'M', 'A', 'S'},
			{'A', 'M', 'S', 'S'},
			{'S', 'M', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'X', 'A', 'A', 'S'},
			{'S', 'M', 'A', 'S'},
			{'A', 'M', 'A', 'S'},
			{'S', 'M', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'X', 'A', 'A', 'X'},
			{'S', 'S', 'M', 'S'},
			{'A', 'A', 'A', 'S'},
			{'S', 'M', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'X', 'A', 'A', 'S'},
			{'S', 'S', 'A', 'S'},
			{'A', 'M', 'A', 'S'},
			{'X', 'S', 'A', 'S'},
		}, 1},
		{[][]rune{
			{'S', 'A', 'A', 'S'},
			{'S', 'A', 'A', 'S'},
			{'A', 'M', 'M', 'S'},
			{'M', 'S', 'A', 'X'},
		}, 1},
		{[][]rune{
			{'S', 'A', 'A', 'S'},
			{'S', 'M', 'A', 'A'},
			{'A', 'M', 'M', 'M'},
			{'S', 'A', 'M', 'X'},
		}, 2},
		{[][]rune{
			{'S', 'A', 'A', 'S'},
			{'S', 'A', 'A', 'S'},
			{'S', 'X', 'A', 'S'},
			{'S', 'M', 'A', 'A'},
			{'A', 'A', 'M', 'M'},
			{'S', 'S', 'M', 'X'},
		}, 2},
	}

	for _, test := range tests {
		result := CountXMAS(test.input)
		if result != test.expected {
			t.Errorf("%v failed, expected %v, got %v", test.input, test.expected, result)
		}
	}
}
