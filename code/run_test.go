package code

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findDifferentElements(t *testing.T) {
	testCases := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Arrays with different elements",
			arr1:     []int{1, 2, 3, 4, 5, 6, 7},
			arr2:     []int{1, 2, 3, 5, 6},
			expected: []int{4, 7},
		},
		{
			name:     "Arrays with duplicate elements",
			arr1:     []int{1, 2, 3, 4, 5, 6, 7, 10},
			arr2:     []int{1, 12, 3, 5, 6, 8, 8},
			expected: []int{2, 4, 7, 10, 12, 8, 8},
		},
		{
			name:     "Arrays with duplicate elements",
			arr1:     []int{1, 2, 2, 3, 4, 5},
			arr2:     []int{2, 3, 3, 5, 6},
			expected: []int{1, 4, 6},
		},
		{
			name:     "One empty array",
			arr1:     []int{1, 2, 3},
			arr2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Two empty arrays",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "Identical arrays",
			arr1:     []int{1, 2, 3},
			arr2:     []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "Identical arrays",
			arr1:     []int{1, 2, 3},
			arr2:     []int{1, 3},
			expected: []int{2},
		},
	}

	var result []int
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result = FindDifferentElements(testCase.arr1, testCase.arr2)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v but got %v", testCase.expected, result)
			}
		})
	}

	fmt.Println("before delete result:", result)
	lr := RemoveDuplicatesInt(result)
	fmt.Println("after delete result:", lr)
}

func Test_RemoveDuplicatesTwo(t *testing.T) {
	testCases := []struct {
		name     string
		arr      []string
		expected []string
	}{
		{
			name:     "22222",
			arr:      []string{"aa", "bb", "cc", "bb", "aa", "cc"},
			expected: []string{"aa", "bb", "cc"},
		},
		{
			name:     "22222",
			arr:      []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RemoveDuplicatesString(tc.arr)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{0}, []int{1}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2,
			8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7,
			4, 0, 0, 6}, []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7,
			3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}},
	}

	for _, test := range tests {
		result := PlusOne(test.input)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
		}
	}
}

func TestClimbStairs(t *testing.T) {
	ClimbStairs(4)
}
