package structure

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *SkipList
	}{
		{
			name: "newtest",
			want: &SkipList{
				Head: &Element{
					forward: make([]*Element, maxLevel),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipList_Get(t *testing.T) {
	sl := New()
	for i := 0; i < 100; i++ {
		sl.Put(float64(i), i)
		if sl.Get(float64(i)).Value.(int) != i {
			t.Errorf("SkipList.Get() = %v, want %v", sl.Get(float64(i)).Value.(int), i)
		}
	}

	if sl.Get(1000) != nil {
		t.Errorf("SkipList.Get() = %v, want %v", sl.Get(1000), nil)
	}
}

func TestSkipList_Put(t *testing.T) {
	sl := New()
	for i := 0; i < 100; i++ {
		sl.Put(float64(i), i)
		if sl.Get(float64(i)).Value.(int) != i {
			t.Errorf("SkipList.Get() = %v, want %v", sl.Get(float64(i)).Value.(int), i)
		}
	}

	if sl.Get(1000) != nil {
		t.Errorf("SkipList.Get() = %v, want %v", sl.Get(1000), nil)
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := New()
	for i := 0; i < 100; i++ {
		sl.Put(float64(i), i)
	}
	temp := sl.Delete(50).Value.(int)
	if temp != 50 {
		t.Errorf("SkipList.Delete() = %v, want %v", temp, 50)
	}
	if sl.Get(50) != nil {
		t.Errorf("SkipList.Get() = %v, want %v", sl.Get(50), nil)
	}
}
