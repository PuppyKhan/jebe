// avl.go

package avl

// Item - the type to be sorted
type Item interface{}

// Node of a binary tree
type Node struct {
	value  Item
	left   *Node
	right  *Node
	parent *Node // doubly linked
	height int   // for AVL property
}

// PrioritizeTreeItem - custom comparison for prioritizing tree items
//  basic sort would need "a < b" ("a > b" for hight to low)
type PrioritizeTreeItem func(a, b Item) bool

// EquivalenceTreeItem - custom comparison for equality of tree items, "a == b"
//  needed for search
type EquivalenceTreeItem func(a, b Item) bool

// BinaryTree holds the root of the tree and its comparison function
type BinaryTree struct {
	root   *Node
	lesser PrioritizeTreeItem
	equals EquivalenceTreeItem
}

// MakeNode puts Item into a Node, sets parent, returns pointer
func MakeNode(val Item, p *Node) *Node {
	return &Node{
		value:  val,
		left:   nil,
		right:  nil,
		parent: p,
		height: 0,
	}
}

// Init sets both root node and comparison func
func (t *BinaryTree) Init(root Item, a PrioritizeTreeItem, b EquivalenceTreeItem) {
	// t.root = t.Add(root)
	t.Insert(root)
	if a == nil {
		t.SetLTIntPrioritizeTreeItem()
	} else {
		t.SetPrioritizeTreeItem(a)
	}
	if b == nil {
		t.SetEqIntEquivalenceTreeItem()
	} else {
		t.SetEquivalenceTreeItem(b)
	}
}

// SetPrioritizeTreeItem - "a < b" or whatever comparison is needed
func (t *BinaryTree) SetPrioritizeTreeItem(a PrioritizeTreeItem) {
	t.lesser = a
}

// SetLTIntPrioritizeTreeItem - default "a < b" as ints
func (t *BinaryTree) SetLTIntPrioritizeTreeItem() {
	t.lesser = func(a, b Item) bool {
		return a.(int) < b.(int)
	}
}

// SetEquivalenceTreeItem - "a == b" or whatever comparison is needed
func (t *BinaryTree) SetEquivalenceTreeItem(b EquivalenceTreeItem) {
	t.equals = b
}

// SetEqIntEquivalenceTreeItem - default "a == b" as ints
func (t *BinaryTree) SetEqIntEquivalenceTreeItem() {
	t.equals = func(a, b Item) bool {
		return a.(int) == b.(int)
	}
}

// Insert a new Item to a tree
func (t *BinaryTree) Insert(newValue Item) {
	var y *Node
	x := t.root
	z := MakeNode(newValue, nil)
	for x != nil {
		y = x
		if t.lesser(z.value, x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z // tree was empty
	} else if t.lesser(z.value, y.value) {
		y.left = z
	} else {
		y.right = z
	}

	// update heights for AVL property
	// TODO: entirely contain w fir property in seperate method?
	// y starts as z's parent here from previous step
	for h := z.height; y != nil; y = y.parent {
		h++
		if y.height < h {
			y.height = h
		}
	}

	// now fix AVL property on inserted node & upwards
	RestoreAVLPropertyTree(z)
}

// InsertRecursive a new Item to a tree
//  call RestoreAVLProperty(branchRoot) afterwards
func (t *BinaryTree) InsertRecursive(branchRoot, newValue *Node) *Node {
	if branchRoot == nil {
		return newValue
	} else if t.lesser(newValue.value, branchRoot.value) {
		if branchRoot.left == nil {
			newValue.parent = branchRoot
			branchRoot.left = t.InsertRecursive(branchRoot.left, newValue)
		} else {
			t.InsertRecursive(branchRoot.left, newValue)
		}
		// update heights for AVL property
		if branchRoot.left.height >= branchRoot.height {
			branchRoot.height = branchRoot.left.height + 1
		}
		branchRoot.left = RestoreAVLProperty(branchRoot.left) // correct level?
		return branchRoot.left.parent
	} else {
		if branchRoot.right == nil {
			newValue.parent = branchRoot
			branchRoot.right = t.InsertRecursive(branchRoot.right, newValue)
		} else {
			t.InsertRecursive(branchRoot.right, newValue)
		}
		// update heights for AVL property
		if branchRoot.right.height >= branchRoot.height {
			branchRoot.height = branchRoot.right.height + 1
		}
		branchRoot.right = RestoreAVLProperty(branchRoot.right) // correct level?
		return branchRoot.right.parent
	}
}

// GetRoot helper for treewalk
func (t BinaryTree) GetRoot() *Node {
	return t.root
}

// InOrderTreeWalk does left, current, right
//  closes channel when done
func (t BinaryTree) InOrderTreeWalk(n *Node, c chan Item) {
	x := n
	var last *Node
	for x != nil {
		if x.left != nil && (last == nil || t.lesser(last.value, x.left.value)) {
			x = x.left
		} else {
			if last != x && (last == nil || t.lesser(last.value, x.value)) {
				c <- x.value
				last = x
			}
			if x.right != nil && (last == nil || t.lesser(last.value, x.right.value)) {
				x = x.right
			} else {
				x = x.parent
			}
		}
	}
	close(c)
}

// InOrderTreeWalkRecursive does left, current, right
//  closes channel when done
func (t BinaryTree) InOrderTreeWalkRecursive(n *Node, c chan Item) {
	if n != nil {
		t.InOrderTreeWalkRecursive(n.left, c)
		c <- n.value
		t.InOrderTreeWalkRecursive(n.right, c)
		if t.equals(n.value, t.root.value) {
			close(c)
		}
	}
}

// SearchRecursive to find node with Item in current branch or nil if none
func (t BinaryTree) SearchRecursive(k Item, current *Node) *Node {
	if current == nil || t.equals(current.value, k) {
		return current
	} else if t.lesser(k, current.value) {
		return t.Search(k, current.left)
	} else {
		return t.Search(k, current.right)
	}
}

// Search to find node with Item in current branch or nil if none
func (t BinaryTree) Search(k Item, current *Node) *Node {
	x := current
	if x == nil {
		x = t.root
	}
	kval := k
	for x != nil && !t.equals(x.value, kval) {
		if t.lesser(kval, x.value) {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// GetMinimum finds lowest value
func (t BinaryTree) GetMinimum(current *Node) *Node {
	x := current
	for x != nil && x.left != nil {
		x = x.left
	}
	return x
}

// GetMaximum finds highest value
func (t BinaryTree) GetMaximum(current *Node) *Node {
	x := current
	for x != nil && x.right != nil {
		x = x.right
	}
	return x
}

// GetNext finds successor in order
func (t BinaryTree) GetNext(current *Node) *Node {
	x := current
	if x == nil {
		return nil
	}
	if x.right != nil {
		return t.GetMinimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}
	return y
}

// GetPrevious finds predecessor in order
func (t BinaryTree) GetPrevious(current *Node) *Node {
	x := current
	if x == nil {
		return nil
	}
	if x.left != nil {
		return t.GetMaximum(x.left)
	}
	y := x.parent
	for y != nil && x == y.left {
		x = y
		y = y.parent
	}
	return y
}

// Transplant switches branch u with v
//  Does not update sub branches
func (t *BinaryTree) Transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Delete removes a node and adjusts tree accordingly
func (t *BinaryTree) Delete(z *Node) {
	if z == nil {
		return
	}
	if z.left == nil {
		t.Transplant(z, z.right)
	} else if z.right == nil {
		t.Transplant(z, z.left)
	} else {
		y := t.GetMinimum(z.right)
		if y.parent != z {
			t.Transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.Transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

// GetHeight returns height of node, including -1 for nil nodes
func GetHeight(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

// GetTreeHeight helper return overall tree height
func (t BinaryTree) GetTreeHeight() int {
	// if t.root == nil {
	// 	return -1
	// }
	// return t.root.height
	return GetHeight(t.root)
}

// IsBalanced returns actual height imbalance
//  0 if children of a node are equal,
//  + for left heavy, - for right heavy
//  per AVL property, balanced: height_difference <= +/-1
func IsBalanced(n *Node) int {
	if n == nil {
		return 0
	}
	return GetHeight(n.left) - GetHeight(n.right)

	// // 1 for left heavy, -1 for right heavy
	// bal := GetHeight(n.left) - GetHeight(n.right)
	// if bal > 1 {
	// 	return 1
	// } else if bal < -1 {
	// 	return -1
	// }
	// return 0
}

// FixHeight resets a node's height based on its current children
func FixHeight(n *Node) {
	r := GetHeight(n.right)
	l := GetHeight(n.left)
	if r > l {
		n.height = r + 1
	} else {
		n.height = l + 1
	}
}

// LeftRotate rotates a node with its right child
//  returns new subtree root, n or its replacement
func LeftRotate(n *Node) *Node {
	if n == nil || n.right == nil {
		return n // can't rotate left
	}
	y := n.right
	y.parent = n.parent
	n.parent = y
	n.right = y.left
	y.left = n
	if y.parent != nil { // not tree root
		if y.parent.left == n {
			y.parent.left = y
		} else if y.parent.right == n {
			y.parent.right = y
		}
	}
	FixHeight(n)
	FixHeight(y) // needed?
	return y
}

// RightRotate rotates a node with its left child
//  returns new subtree root, n or its replacement
func RightRotate(n *Node) *Node {
	if n == nil || n.left == nil {
		return n // can't rotate right
	}
	x := n.left
	x.parent = n.parent
	n.parent = x
	n.left = x.right
	x.right = n
	if x.parent != nil { // not tree root
		if x.parent.right == n {
			x.parent.right = x
		} else if x.parent.left == n {
			x.parent.left = x
		}
	}
	FixHeight(n)
	FixHeight(x) // needed?
	return x
}

// RestoreAVLProperty of an inserted node only
//  returns n, or its replacement if rotated
func RestoreAVLProperty(n *Node) *Node {
	b := IsBalanced(n)
	if b < -1 { // too right heavy, fix
		if IsBalanced(n.right) > 0 {
			// right child is left heavy so 2 rotations
			RightRotate(n.right)
		} // else right child is right heavy or balanced, so 1 rotation
		return LeftRotate(n)
	} else if b > 1 { // too left heavy, fix
		if IsBalanced(n.left) < 0 {
			// left child is right heavy so 2 rotations
			LeftRotate(n.left)
		} // else left child is left heavy or balanced, so 1 rotation
		return RightRotate(n)
	} // else AVL balanced (height of children <= +/-1 difference)
	return n
}

// RestoreAVLPropertyTree of an inserted node, traversing upwards
func RestoreAVLPropertyTree(n *Node) {
	for currentNode := n; n != nil; n = n.parent {
		currentNode = RestoreAVLProperty(currentNode)
		// b := IsBalanced(currentNode)
		// if b < -1 { // too right heavy, fix
		// 	if IsBalanced(currentNode.right) > 0 {
		// 		// right child is left heavy so 2 rotations
		// 		RightRotate(currentNode.right)
		// 	} // else right child is right heavy or balanced, so 1 rotation
		// 	currentNode = LeftRotate(currentNode)
		// } else if b > 1 { // too left heavy, fix
		// 	if IsBalanced(currentNode.left) < 0 {
		// 		// left child is right heavy so 2 rotations
		// 		LeftRotate(currentNode.left)
		// 	} // else left child is left heavy or balanced, so 1 rotation
		// 	currentNode = RightRotate(currentNode)
		// } // else AVL balanced (height of children <= +/-1 difference)
	}
}
