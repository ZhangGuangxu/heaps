package heaps

import (
	"testing"
)

func TestMaxIntBinaryHeap(t *testing.T) {
	a := NewMinBinaryIntHeap()
	if !a.IsEmpty() {
		t.Error("a is not empty")
	}

	b := NewIntHeap(false, 2, 110)
	if !b.IsEmpty() {
		t.Error("b is not empty")
	}
	b.Insert(10)
	b.Insert(9)
	b.Pop()

	h := NewMaxBinaryIntHeap()
	h.Insert(1)
	if h.IsEmpty() {
		t.Error("h should not be empty after Insert")
	}
	h.Insert(2)
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	if h.tail != 4 {
		t.Errorf("h.tail got %d, want %d", h.tail, 4)
	}
	h.Insert(5)
	h.show()

	v, err := h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 5 {
		t.Errorf("h.Pop got %d, want %d", v, 5)
	}

	v, err = h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 5 {
		t.Errorf("h.Pop got %d, want %d", v, 5)
	}

	v, err = h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 4 {
		t.Errorf("h.Pop got %d, want %d", v, 4)
	}

	v, err = h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 3 {
		t.Errorf("h.Pop got %d, want %d", v, 3)
	}

	v, err = h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 2 {
		t.Errorf("h.Pop got %d, want %d", v, 2)
	}
	h.show()

	v, err = h.Pop()
	if err != nil {
		t.Error("h.Pop got error")
	}
	if v != 1 {
		t.Errorf("h.Pop got %d, want %d", v, 1)
	}
	h.show()

	_, err = h.Pop()
	if err != ErrEmptyHeap {
		t.Errorf("h.Pop should return %v", ErrEmptyHeap)
	}
	h.show()
}

func TestMin4WayIntHeap(t *testing.T) {
	h := NewMinIntHeapWithWay(false, 4)

	for i, j := 0, 100; i < 21; {
		h.Insert(j)
		i++
		j--
	}
	h.show()

	v, err := h.Pop()
	if v != 80 {
		t.Errorf("h.Pop() got %d, want %d", v, 80)
	}
	if err != nil {
		t.Errorf("h.Pop() got %v, want %v", err, nil)
	}
	h.show()

	v, err = h.Pop()
	if v != 81 {
		t.Errorf("h.Pop() got %d, want %d", v, 80)
	}
	if err != nil {
		t.Errorf("h.Pop() got %v, want %v", err, nil)
	}
	h.show()

	for i := 0; i < 18; i++ {
		h.Pop()
	}
	h.show()
	v, err = h.Pop()
	if v != 100 {
		t.Errorf("h.Pop() got %d, want %d", v, 100)
	}
	v, err = h.Pop()
	if err != ErrEmptyHeap {
		t.Errorf("h.Pop() got %v, want %v", v, ErrEmptyHeap)
	}
}
