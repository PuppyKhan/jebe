// avl_test.go

package avl

import (
	"strings"
	"testing"
)

func TestBinarySearchTreeWalk(t *testing.T) {
	tests := []struct {
		givenArray        []Item
		givenPriorityFunc func(a, b Item) bool
		givenEqualityFunc func(a, b Item) bool
		givenDeletables   []Item

		wantArray      []Item
		wantFinalArray []Item
	}{
		{
			[]Item{},
			nil, // default int sort
			nil,
			[]Item{},

			[]Item{},
			[]Item{},
		},
		{
			[]Item{5, 2, 9, 7, 1, 3, 4, 24, 14, 34, -1, 12, 18, 10, 16},
			nil, // default int sort
			nil,
			[]Item{},

			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
		},
		{
			[]Item{5, 2, 9, 7, 1, 3, 4, 24, 14, 34, -1, 12, 18, 10, 16},
			nil, // default int sort
			nil,
			[]Item{24, -1, 5},

			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
			[]Item{1, 2, 3, 4, 7, 9, 10, 12, 14, 16, 18, 34},
		},
		{
			[]Item{"kenny", "kyle", "eric", "chef", "stan", "timmy"},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) < 0)
			},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) == 0)
			},
			[]Item{"timmy"},

			[]Item{"chef", "eric", "kenny", "kyle", "stan", "timmy"},
			[]Item{"chef", "eric", "kenny", "kyle", "stan"},
		},
	}

	for i, test := range tests {
		var tree BinaryTree

		for x, n := range test.givenArray {
			if x == 0 {
				tree.Init(n, test.givenPriorityFunc, test.givenEqualityFunc)
				if !tree.equals(tree.GetRoot().value, n) {
					t.Errorf("%d: bad root", i)
				}
			} else {
				tree.Insert(n)
			}
		}

		ch1 := make(chan Item, 1)
		go tree.InOrderTreeWalk(tree.GetRoot(), ch1)

		x := 0
		for y := range ch1 {
			if !tree.equals(y, test.wantArray[x]) {
				t.Errorf("%d: Invalid order, found: %d, expected: %d", i, y, test.wantArray[x])
				return
			}
			x++
		}

		for _, n := range test.givenDeletables {
			x := tree.Search(n, nil)
			tree.Delete(x)
		}

		ch2 := make(chan Item, 1)
		go tree.InOrderTreeWalk(tree.GetRoot(), ch2)

		x = 0
		for y := range ch2 {
			if !tree.equals(y, test.wantFinalArray[x]) {
				t.Errorf("%d: Invalid order, found: %d, expected: %d", i, y, test.wantFinalArray[x])
				return
			}
			x++
		}
	}
}

func TestBinarySearchTreeWalkRecursives(t *testing.T) {
	tests := []struct {
		givenArray        []Item
		givenPriorityFunc func(a, b Item) bool
		givenEqualityFunc func(a, b Item) bool
		givenDeletables   []Item

		wantArray      []Item
		wantFinalArray []Item
	}{
		{
			[]Item{},
			nil, // default int sort
			nil,
			[]Item{},

			[]Item{},
			[]Item{},
		},
		{
			[]Item{5, 2, 9, 7, 1, 3, 4, 24, 14, 34, -1, 12, 18, 10, 16},
			nil, // default int sort
			nil,
			[]Item{},

			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
		},
		{
			[]Item{5, 2, 9, 7, 1, 3, 4, 24, 14, 34, -1, 12, 18, 10, 16},
			nil, // default int sort
			nil,
			[]Item{24, -1, 5},

			[]Item{-1, 1, 2, 3, 4, 5, 7, 9, 10, 12, 14, 16, 18, 24, 34},
			[]Item{1, 2, 3, 4, 7, 9, 10, 12, 14, 16, 18, 34},
		},
		{
			[]Item{"kenny", "kyle", "eric", "chef", "stan", "timmy"},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) < 0)
			},
			func(a, b Item) bool {
				return (strings.Compare(a.(string), b.(string)) == 0)
			},
			[]Item{"timmy"},

			[]Item{"chef", "eric", "kenny", "kyle", "stan", "timmy"},
			[]Item{"chef", "eric", "kenny", "kyle", "stan"},
		},
	}

	for i, test := range tests {
		var tree BinaryTree

		for x, n := range test.givenArray {
			if x == 0 {
				tree.Init(n, test.givenPriorityFunc, test.givenEqualityFunc)
				if !tree.equals(tree.GetRoot().value, n) {
					t.Errorf("%d: bad root", i)
				}
			} else {
				tree.InsertRecursive(tree.GetRoot(), MakeNode(n, nil))
			}
		}

		// InOrderTreeWalkRecursive() only works on non nil tree
		if tree.GetRoot() == nil {
			if len(test.wantArray) > 0 || len(test.wantFinalArray) > 0 {
				t.Errorf("%d: Empty result when non empty wanted", i)
			}
			continue
		}

		ch1 := make(chan Item, 1)
		go tree.InOrderTreeWalkRecursive(tree.GetRoot(), ch1)

		x := 0
		for y := range ch1 {
			if !tree.equals(y, test.wantArray[x]) {
				t.Errorf("%d: Invalid order, found: %d, expected: %d", i, y, test.wantArray[x])
				return
			}
			x++
		}

		for _, n := range test.givenDeletables {
			x := tree.SearchRecursive(n, tree.GetRoot())
			tree.Delete(x)
		}

		ch2 := make(chan Item, 1)
		go tree.InOrderTreeWalkRecursive(tree.GetRoot(), ch2)

		x = 0
		for y := range ch2 {
			if !tree.equals(y, test.wantFinalArray[x]) {
				t.Errorf("%d: Invalid order, found: %d, expected: %d", i, y, test.wantFinalArray[x])
				return
			}
			x++
		}
	}
}
