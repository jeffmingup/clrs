package exercises

import (
	"testing"
)

func TestSearchSum(t *testing.T) {
	arr := []int{5, 7, 3, 2, 7, 55, 33, 55, 13, 56, 758}
	if !SearchSum(arr, 8) {
		t.Fail()
	}
}
