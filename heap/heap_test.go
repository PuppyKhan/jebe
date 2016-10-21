// heap_test.go

package heap

import (
	"strings"
	"testing"
)

func TestHeapSort(t *testing.T) {
	// HeapSort
	var heap BinaryHeap

	tests := []struct {
		givenArray        []Item
		givenPriorityFunc func(a, b Item) bool

		wantArray []Item
	}{
		{
			[]Item{"danny", "13", "55", "bob", "6", "17999", "99", "22"},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) > 0)
			},

			[]Item{"13", "17999", "22", "55", "6", "99", "bob", "danny"},
		},
		{
			[]Item{"danny", "13", "55", "bob", "6", "17999", "99", "22"},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) < 0) // reverse sort
			},

			[]Item{"danny", "bob", "99", "6", "55", "22", "17999", "13"},
		},
		{
			[]Item{23, 13, 55, 10, 6, 99, 22},
			nil, // default int sort

			[]Item{6, 10, 13, 22, 23, 55, 99},
		},
		{
			[]Item{},
			nil, // default int sort

			[]Item{},
		},
		{
			[]Item{23},
			nil, // default int sort

			[]Item{23},
		},
		{
			[]Item{23, 24},
			nil, // default int sort

			[]Item{23, 24},
		},
		{
			[]Item{24, 23},
			nil, // default int sort

			[]Item{23, 24},
		},
	}

	for i, test := range tests {
		sortedArray := heap.Sort(test.givenArray, test.givenPriorityFunc)

		for x := 0; x < len(sortedArray); x++ {
			if sortedArray[x] != test.wantArray[x] {
				t.Errorf("%d: Invalid order, found: %d, expected: %d", i, sortedArray[x], test.wantArray[x])
				return
			}

		}
	}
}
