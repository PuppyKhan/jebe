// heap.go

package heap

import "errors"

// Item - the type to be sorted
type Item interface{}

// PrioritizeHeapItem - custom comparison for prioritizing heap items
//  basic max heap would need "a > b" ("a < b" for min heap)
//  maybe create a default? https://newfivefour.com/golang-interface-type-assertions-switch.html
type PrioritizeHeapItem func(a, b Item) bool

// BinaryHeap - the basic heap with a slice and a comparison method
type BinaryHeap struct {
	array       []Item
	greater     PrioritizeHeapItem
	logicalSize uint
}

// Parent node number of node i
func Parent(i uint) (uint, error) {
	if i == 0 {
		return 0, errors.New("no parent")
	}
	return (i - 1) / 2, nil // integer division truncates, thus floor result
}

// Left child node number of node i
func Left(i uint) uint {
	return (2 * i) + 1
}

// Right child node number of node i
func Right(i uint) uint {
	return (2 * i) + 2
}

// ArraySize - real size of slice as uint
func (h BinaryHeap) ArraySize() uint {
	return uint(len(h.array))
}

// Size as uint (logical size for sorting)
func (h BinaryHeap) Size() uint {
	return h.logicalSize
}

// Value - the item at location i
func (h BinaryHeap) Value(i uint) Item {
	return h.array[i]
}

// SwapHeapItem - reverse two items in heap
func (h *BinaryHeap) SwapHeapItem(a, b uint) {
	tmp := h.array[a]
	h.array[a] = h.array[b]
	h.array[b] = tmp
}

// SetPrioritizeHeapItem - "a > b" or whatever comparison is needed
func (h *BinaryHeap) SetPrioritizeHeapItem(a PrioritizeHeapItem) {
	h.greater = a
}

// SetGTIntPrioritizeHeapItem - default "a > b" as ints
func (h *BinaryHeap) SetGTIntPrioritizeHeapItem() {
	h.greater = func(a, b Item) bool {
		return a.(int) > b.(int)
	}
}

// SetArray copies slice header into BinaryHeap
//  note this means the array items themselves are used in place
func (h *BinaryHeap) SetArray(array []Item) {
	h.array = array
	h.logicalSize = h.ArraySize()
}

// MaxHeapify - fix a branch of the heap
//  NOTE: uses 0 based array indexing
func (h *BinaryHeap) MaxHeapify(i uint) {
	l := Left(i)
	r := Right(i)
	largest := i
	if l < h.Size() && h.greater(h.array[l], h.array[i]) {
		largest = l
	}
	if r < h.Size() && h.greater(h.array[r], h.array[largest]) {
		largest = r
	}
	if largest != i {
		h.SwapHeapItem(i, largest)
		h.MaxHeapify(largest) // TODO: loop instead of recursion
	}
}

// BuildMaxHeap - convert unsorted array into Max Heap
func (h *BinaryHeap) BuildMaxHeap() {
	length := h.Size()
	for i := (length / 2); i > 0; i-- {
		// i is 1 based index here to avoid wraparound of uint to uint_max
		h.MaxHeapify(i - 1) // so adjust to 0 based when used
	}
}

// Sort - sort array
//  initializes heap, sorts, returns sorted slice
//  if gt is nil, default casts HeapItems as type int
func (h *BinaryHeap) Sort(array []Item, gt PrioritizeHeapItem) []Item {
	if array == nil {
		return nil
	}
	h.SetArray(array)
	if gt == nil {
		h.SetGTIntPrioritizeHeapItem()
	} else {
		h.greater = gt
	}
	h.BuildMaxHeap()
	for i := h.Size(); i > 1; i-- {
		// i is 1 based index here to avoid wraparound of uint to uint_max
		h.SwapHeapItem(0, i-1) // so adjust to 0 based when used
		h.logicalSize--
		h.MaxHeapify(0)
	}
	return h.array
}

// Using heap as a priority queue

// Insert - add new item to heap, grow if necessary
func (h *BinaryHeap) Insert(key Item) {
	if h.Size() < h.ArraySize() {
		h.array[h.logicalSize] = nil
		h.logicalSize++
	} else {
		h.SetArray(append(h.array, nil))
	}
	h.ReplaceItem(h.logicalSize-1, key)
}

// Push - alias for Insert()
func (h *BinaryHeap) Push(key Item) {
	h.Insert(key)
}

// Maximum - returns prioritzed Item without removing it from heap
func (h BinaryHeap) Maximum() Item {
	return h.array[0]
}

// Peek - alias for Maximum()
func (h BinaryHeap) Peek() Item {
	return h.Maximum()
}

// ExtractMax - returns prioritzed Item and removes it from heap
func (h *BinaryHeap) ExtractMax() Item {
	if h.logicalSize < 1 {
		return nil
	}
	h.logicalSize--
	h.SwapHeapItem(0, h.logicalSize)
	h.MaxHeapify(0)
	return h.array[h.logicalSize]
}

// Pop - alias for ExtractMax()
func (h *BinaryHeap) Pop() Item {
	return h.ExtractMax()
}

// ReplaceItem - (instead of IncreaseKey)
func (h *BinaryHeap) ReplaceItem(i uint, key Item) {
	if i >= h.Size() { // ArraySize() instead?
		return
	}
	h.array[i] = key
	for n, err := i, error(nil); err == nil && n >= 0; n, err = Parent(n) {
		h.MaxHeapify(n)
	}
}

// For satisfying sort.Interface

// Len - alias for Size() except as signed int
func (h BinaryHeap) Len() int {
	return int(h.Size())
}

// Swap - alias for SwapHeapItem() except as signed ints
func (h *BinaryHeap) Swap(a, b int) {
	h.SwapHeapItem(uint(a), uint(b))
}

// Less - actually returns <= (aka: not >) so may be less optimal for such use
func (h *BinaryHeap) Less(a, b int) bool {
	return !h.greater(h.array[a], h.array[b])
}
