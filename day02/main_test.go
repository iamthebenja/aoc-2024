package main

import "testing"

func TestCheckReport2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{74, 76, 78, 79, 76}, true},
		{[]int{55, 58, 59, 57, 60, 61, 68}, false},
		{[]int{7, 10, 8, 10, 11}, false},
		{[]int{29, 28, 27, 25, 26, 22, 20}, true},
		{[]int{29, 28, 27, 25, 26, 25, 22, 20}, true},
		{[]int{89, 91, 92, 95, 93, 94}, true},
		{[]int{89, 87, 86, 83, 85, 84}, true},
		{[]int{48, 46, 47, 49, 51, 54, 56}, true},
		{[]int{1, 1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 3, 4, 5, 5}, true},
		{[]int{5, 1, 2, 3, 4, 5}, true},
		{[]int{1, 4, 3, 2, 1}, true},
		{[]int{1, 6, 7, 8, 9}, true},
		{[]int{1, 2, 3, 4, 3}, true},
		{[]int{9, 8, 7, 6, 7}, true},
		{[]int{8, 9, 10, 11}, true},
		{[]int{57, 57, 55, 58, 54}, false},
		{[]int{1, 2, 5, 3, 4, 7}, true},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := CheckReport2(test.input)
		if result != test.expected {
			t.Errorf("%v failed, expected %v, got %v", test.input, test.expected, result)
		}
	}
}
