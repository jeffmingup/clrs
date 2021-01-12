package exercises

import (
	"log"
	"testing"
)

var fib = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for _, v := range fib {
		ret := f()
		log.Println(ret)
		if ret != v {
			t.Errorf("%v不正确", v)
		}
	}
}
