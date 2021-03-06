# Jebe

Implementing a couple of algorithms in Go for fun

Each acts upon `type Item interface{}` for value and needs custom defined comparison functions so they work for any data, not just int.

## Packages

### Heap

HeapSort: Sort underlying array in place using HeapSort algorithm

Can be used as a max, min or custom priority heap by setting the comparison with a custom PrioritizeHeapItem(), with Push/Pop aliases
- Push() could dynamically grow heap, thus needing a copy operation so original may not be sorted

Satisfies sort.Interface

```go
import "github.com/PuppyKhan/jebe/heap"
```

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

### BST

Binary Search Tree, order can be custom defined.

Both recursive and nonrecursive methods given as an excercise, use the nonrecursive in practice. Use:
- `Insert()` instead of `InsertRecursive()`
- `InOrderTreeWalk()` instead of `InOrderTreeWalkRecursive()`
- `Search()` instead of `SearchRecursive()`

```go
import "github.com/PuppyKhan/jebe/bst"
```

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

### AVL

AVL Sort using balanced Binary Search Tree

Supports sorting, traversal, and priority queue functionality, on min and max

```go
import "github.com/PuppyKhan/jebe/avl"
```

Follows:
- https://www.youtube.com/watch?v=FNeL18KsWPc
- http://www.geeksforgeeks.org/avl-tree-set-2-deletion/
- https://courses.cs.washington.edu/courses/cse332/10sp/lectures/lecture8.pdf

## Jebe meaning

Jebe is the name of one of Chinggis Khaan's greatest warriors, whose name means "weapon" - though probably something more specific like a particular type of arrowhead.

This offers a wonderful tagline:

_These are a good collection of arrows to have in your quiver._

## Author

Luigi Kapaj <puppy at viahistoria.com>
