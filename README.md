Implementing a couple of algorithms in Go for fun

Each acts upon "type Item interface{}" for value and needs custom defined comparison functions so they work for any data, not just int.

Heap
====
HeapSort: Sort underlying array in place using HeapSort algorithm

Can be used as a max, min or custom priority heap by setting the comparison with a custom PrioritizeHeapItem(), with Push/Pop aliases
- Push could dynamically grow heap, thus needing a copy operation

Satifies sort.Interface

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

BST
===
Binary Search Tree, order can be custom defined

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

AVL
===
AVL Sort using balanced Binary Search Tree
(unfinished)
