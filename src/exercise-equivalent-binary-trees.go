package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t != nil {
		walkRecursive(t.Left, ch)
		ch <- t.Value
		walkRecursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var br bool
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i == <-ch2 {
			br = true
		} else {
			br = false
			break
		}
	}
	return br
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("1 and 1 same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 2 same: ", Same(tree.New(1), tree.New(2)))
}
