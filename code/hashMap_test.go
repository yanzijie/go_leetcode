package code

import (
	"reflect"
	"testing"
)

// 1.两数之和
func TestTwoSumMap(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := TwoSumMap(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
