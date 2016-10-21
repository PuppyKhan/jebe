// avl_test.go

package avl

import "testing"

func TestHeapSort(t *testing.T) {
	var tree BinaryTree

	tree.Init(10, nil, nil)
	if tree.GetRoot().value != 10 {
		t.Errorf("bad root")
	}
}
