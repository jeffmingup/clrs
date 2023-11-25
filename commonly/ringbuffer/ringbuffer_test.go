package ringbuffer

import (
	"math"
	"testing"
)

func TestRingBuffer_Push(t *testing.T) {
	rb := NewRingBuffer()
	for i := 0; i < defaultSize; i++ {
		rb.Push(i)
		if rb.head != uint(i+1) {
			t.Errorf("RingBuffer.Push() failed, expected head = %v, got head = %v", i+1, rb.head)
		}
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RingBuffer.Push() failed, expected panic for full buffer")
		}
	}()
	rb.Push(defaultSize)
}

func TestRingBuffer_Pop(t *testing.T) {
	rb := NewRingBuffer()
	for i := 0; i < defaultSize; i++ {
		rb.Push(i)
	}
	for i := 0; i < defaultSize; i++ {
		data := rb.Pop()
		if data.(int) != i {
			t.Errorf("RingBuffer.Pop() failed, expected data = %v, got data = %v", i, data)
		}
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RingBuffer.Pop() failed, expected panic for empty buffer")
		}
	}()
	rb.Pop()
}

func TestRingBuffer_full(t *testing.T) {
	rb := NewRingBuffer()
	for i := 0; i < defaultSize; i++ {
		rb.Push(i)
	}
	if !rb.full() {
		t.Errorf("RingBuffer.full() failed, expected full buffer")
	}
	rb.Pop()
	if rb.full() {
		t.Errorf("RingBuffer.full() failed, expected non-full buffer")
	}
}

func TestRingBuffer_empty(t *testing.T) {
	rb := NewRingBuffer()
	if !rb.empty() {
		t.Errorf("RingBuffer.empty() failed, expected empty buffer")
	}
	rb.Push(1)
	if rb.empty() {
		t.Errorf("RingBuffer.empty() failed, expected non-empty buffer")
	}
}

func TestUintOverflow(t *testing.T) {
	var in uint = math.MaxUint64
	var out uint = math.MaxUint64 - 1
	if in-out != 1 {
		t.Errorf("Uint overflow failed, expected %v, got %v", 1, in-out)
	}
	in++
	if in-out != 2 {
		t.Errorf("Uint overflow failed, expected %v, got %v", 2, in-out)
	}
	out++
	if in-out != 1 {
		t.Errorf("Uint overflow failed, expected %v, got %v", 1, in-out)
	}
}

var (
	Len  = 8
	Mask = Len - 1
	In   = 8 - 5
)

// % len
func BenchmarkModLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = In % Len
	}
}

// & Mask
func BenchmarkAndMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = In & Mask
	}
}
