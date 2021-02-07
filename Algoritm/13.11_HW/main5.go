package main

import (
	"errors"
	"fmt"
	//"log"
)

/*
## A Tree Node
Based on the above definition of a binary tree, a tree node consists of
* a value,
* a left subtree, and
* a right subtree.
By the way, this is a *recursive* data structure: Each subtree of a node is also a node containing subtrees.
In this minimal setup, the tree contains simple string data.
*/

// `Node` contains the search value, some data, a left child node, and a right child node.
type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func main() {
	n := &Node{value: 8}

	n.Insert(1)
	n.Insert(2)
	n.Insert(3)
	n.Insert(4)
	n.Insert(5)
	n.Insert(6)
	n.Insert(7)

	n.PrintInorder()
	fmt.Println("")
	//PrintInorder prints the elements in order
	(func(n *Node) PrintInorder)()
	{

		if n == nil {

			return
		}

		n.left.PrintInorder()
		fmt.Print(n.value)
		n.right.PrintInorder()
	}
}

/* ## Node Operations
### Insert
To insert a value into a sorted tree, we need to find the correct place to add a new node with this value. Here is how:
1. Start at the root node.
2. Compare the new value with the current node's value.
	* If it is the same, stop. We have that value already.
	* If it is smaller, repeat 2. with the left child node. If there is no left child node, add a new one with the new value. Stop.
	* IF it is greater, repeat 2. with the right child node, or create a new one if none exists. Stop.
Sounds quite easy, doesn't it? Just keep in mind we do not take care of keeping the tree balanced. Doing so adds a bit of complexity but for now we don't care about this.
HYPE[Insert](TreeInsert.html)
The Insert method we define here works *recursively*. That is, it calls itself but with one of the child nodes as the new receiver. If you are unfamiliar with recursion, see the little example [here](https://en.wikipedia.org/wiki/Recursion#In_computer_science) or have a look at [this factorial function](https://play.golang.org/p/feMIAgYWg3).
*/

// `Insert` inserts new data into the tree, at the position determined by the search value.
// Return values:
//
// * `true` if the data was successfully inserted,
// * `false` if the data value already exists in the tree.
func (n *Node) Insert(value, data string) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	// If the data is already in the tree, return.
	case value == n.Value:
		return nil
	// If the data value is less than the current node's value, and if the left child node is `nil`, insert a new left child node. Else call `Insert` on the left subtree.
	case value < n.Valuealue:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	// If the data value is greater than the current node's value, do the same but for the right subtree.
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}
