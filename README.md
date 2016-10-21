# Jebe

Implementing a couple of algorithms in Go for fun

Each acts upon `type Item interface{}` for value and needs custom defined comparison functions so they work for any data, not just int.

## Packages

### Heap

HeapSort: Sort underlying array in place using HeapSort algorithm

Can be used as a max, min or custom priority heap by setting the comparison with a custom PrioritizeHeapItem(), with Push/Pop aliases
- Push could dynamically grow heap, thus needing a copy operation

Satifies sort.Interface

```go
import "github.com/PuppyKhan/jebe/heap"
```

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

### BST

Binary Search Tree, order can be custom defined

```go
import "github.com/PuppyKhan/jebe/bst"
```

Follows pseudocode from "Introduction to Algorithms" by Cormen, Leiserson, Rivest, Stein

### AVL

AVL Sort using balanced Binary Search Tree
(unfinished)

```go
import "github.com/PuppyKhan/jebe/avl"
```

## Jebe meaning

Jebe is the name of one of Chinggis Khaan's greatest warriors, whose name means "weapon" - though probably something more specific like a particular type of arrowhead.

This offers a wonderful tagline:

_These are a good collection of arrows to have in your quiver._

## Author

Luigi Kapaj <puppy at viahistoria.com>
