package exercises

import (
	"testing"
)

func TestBinarySystemAdd(t *testing.T) {
	a := []uint8{1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1}
	b := []uint8{1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0}
	c := BinarySystemAdd(a, b)
	d := []uint8{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}
	for k := range c {
		if c[k] != d[k] {
			t.Errorf("%v与正确答案%v不符", c, d)
			break
		}
	}

}
