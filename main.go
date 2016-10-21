// main.go

package main

import (
	"fmt"
	"strings"

	"github.com/PuppyKhan/jebe/bst"
	"github.com/PuppyKhan/jebe/heap"
)

func main() {
	// Some example code

	// HeapSort
	fmt.Println("HeapSort")
	var myHeap heap.BinaryHeap

	// myArray := []heap.Item{23, 13, 55, 10, 6, 99, 22}
	myArray := []heap.Item{"danny", "13", "55", "bob", "6", "17", "99", "22"}

	greater := func(a, b heap.Item) bool {
		return (strings.Compare(a.(string), b.(string)) > 0) // reverse sort
	}

	fmt.Println("Unsorted")
	for i := 0; i < len(myArray); i++ {
		fmt.Println("item", i, ": ", myArray[i])
	}

	// sortedArray := myHeap.Sort(myArray, nil)
	sortedArray := myHeap.Sort(myArray, greater)

	fmt.Println("Sorted")
	for i := 0; i < len(sortedArray); i++ {
		fmt.Println("item", i, ": ", sortedArray[i])
	}

	// Priority Queue
	fmt.Println("Priority Queue")
	var secHeap heap.BinaryHeap

	secArray := make([]heap.Item, 0, 10)

	secHeap.SetArray(secArray)

	secHeap.SetGTIntPrioritizeHeapItem()

	fmt.Println("Insert:", 10)
	secHeap.Insert(10)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("Insert:", 5)
	secHeap.Insert(5)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("Insert:", 15)
	secHeap.Insert(15)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("Insert:", 11)
	secHeap.Insert(11)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("Insert:", 2)
	secHeap.Insert(2)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("Insert:", 3)
	secHeap.Insert(3)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("ExtractMax:", secHeap.ExtractMax())
	fmt.Println("ExtractMax:", secHeap.ExtractMax())

	fmt.Println("Insert:", 23)
	secHeap.Insert(23)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("ExtractMax:", secHeap.ExtractMax())
	fmt.Println("ExtractMax:", secHeap.ExtractMax())

	fmt.Println("Insert:", 1)
	secHeap.Insert(1)
	fmt.Println("Peek Maximum:", secHeap.Maximum())

	fmt.Println("ExtractMax:", secHeap.ExtractMax())

	// BinaryTree
	var tree bst.BinaryTree

	fmt.Println("Insert:", 5)

	tree.Init(5, nil, nil) // default Item is type int
	ch1 := make(chan bst.Item, 1)
	go tree.InOrderTreeWalkRecursive(tree.GetRoot(), ch1)
	x := <-ch1
	fmt.Println(x)

	nodes := []int{2, 9, 7, 1, 3, 4, 24, 14, 34, -1, 12, 18, 10, 16}

	for sz, i := range nodes {
		fmt.Println("Insert:", i)

		// tree.InsertRecursive(tree.GetRoot(), bst.MakeNode(i, nil))
		tree.Insert(i)

		// bst.InOrderTreeWalkRecursive(tree.GetRoot())
		// tree.InOrderTreeWalk(tree.GetRoot())

		ch2 := make(chan bst.Item, sz+1)
		go tree.InOrderTreeWalk(tree.GetRoot(), ch2)
		// go tree.InOrderTreeWalkRecursive(tree.GetRoot(), ch2)
		for y := range ch2 {
			fmt.Println(y)
		}
	}

	fmt.Println("Delete:", 2)
	tree.Delete(tree.Search(2, tree.GetRoot()))
	ch3 := make(chan bst.Item)
	go tree.InOrderTreeWalk(tree.GetRoot(), ch3)
	for z := range ch3 {
		fmt.Println(z)
	}

}
