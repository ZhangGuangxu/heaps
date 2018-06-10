package heaps

import (
	"errors"
	"fmt"
)

// ErrEmptyHeap tells us the heap is empty.
var ErrEmptyHeap = errors.New("empty heap")

const (
	invalidTail = -1
)

// IntHeap is a struct represents binary heap contains only int-items.
type IntHeap struct {
	compare func(a, b int) bool // index to a is smaller than index to b
	way     int
	data    []int
	tail    int // the index of the last item in data
}

// NewMaxBinaryIntHeap returns an instance of a max binary IntHeap.
func NewMaxBinaryIntHeap() *IntHeap {
	return NewIntHeapWithSize(true, 2, 1)
}

// NewMinBinaryIntHeap returns an instance of a min binary IntHeap.
func NewMinBinaryIntHeap() *IntHeap {
	return NewIntHeapWithSize(false, 2, 1)
}

// NewIntHeapWithSize returns an instance of a max or min n-way IntHeap with init-size.
func NewIntHeapWithSize(max bool, nWay int, s int) *IntHeap {
	var f func(a, b int) bool
	if max {
		f = func(a, b int) bool {
			return a < b
		}
	} else {
		f = func(a, b int) bool {
			return a > b
		}
	}
	return &IntHeap{
		compare: f,
		way:     nWay,
		data:    make([]int, s),
		tail:    invalidTail,
	}
}

// IsEmpty returns true when heap is empty.
func (h *IntHeap) IsEmpty() bool {
	return h.tail == invalidTail
}

// Insert inserts an item into heap.
func (h *IntHeap) Insert(x int) {
	if h.tail+1 >= len(h.data) {
		h.makeSpace()
	}

	h.tail++
	h.data[h.tail] = x
	h.siftUp()
}

func (h *IntHeap) makeSpace() {
	d := make([]int, len(h.data)*2+1)
	copy(d, h.data)
	h.data = d
}

// Pop removes the root node from the heap and returns that node.
// Pop returns error when heap is empty. So you had better make sure
// the heap is not empty before you invode Pop on it.
func (h *IntHeap) Pop() (int, error) {
	if h.IsEmpty() {
		return 0, ErrEmptyHeap
	}

	v := h.data[0]
	h.data[0] = h.data[h.tail]
	h.tail--
	h.siftDown()
	return v, nil
}

func (h *IntHeap) siftUp() {
	idx := h.tail
	parentIdx := 0

	for {
		if idx == 0 {
			return
		}

		r := idx % h.way
		if r == 0 {
			parentIdx = (idx - h.way) / h.way
		} else {
			parentIdx = (idx - r) / h.way
		}
		if h.compare(h.data[parentIdx], h.data[idx]) {
			h.data[parentIdx], h.data[idx] = h.data[idx], h.data[parentIdx]
			idx = parentIdx
		} else {
			return
		}
	}
}

func (h *IntHeap) siftDown() {
	if h.tail <= 0 {
		return
	}

	idx := 0
	swap := false

	for {
		idx, swap = h.compareWithChildren(idx)
		if !swap {
			return
		}
	}
}

func (h *IntHeap) compareWithChildren(idx int) (newIdx int, swap bool) {
	newIdx = idx
	max := h.data[idx]

	for a := 1; a <= h.way; a++ {
		i := idx*h.way + a
		if i > h.tail {
			break
		}
		if h.compare(max, h.data[i]) {
			newIdx = i
			max = h.data[i]
		}
	}

	swap = newIdx != idx
	if swap {
		h.data[idx], h.data[newIdx] = h.data[newIdx], h.data[idx]
	}
	return
}

func (h *IntHeap) show() {
	for i := 0; i <= h.tail; i++ {
		fmt.Printf("%d ", h.data[i])
	}
	fmt.Println()
}
