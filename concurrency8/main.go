// Exercise: Equivalent Binary Trees
// 1. Implement the Walk function.
//
// 2. Test the Walk function.
//
// The function tree.New(k) constructs a randomly-structured (but always sorted)
// binary tree holding the values k, 2k, 3k, ..., 10k.
// Create a new channel ch and kick off the walker:
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel.
// It should be the numbers 1, 2, 3, ..., 10.
//
// 3. Implement the Same function using Walk to determine
// whether t1 and t2 store the same values.
//
// 4. Test the Same function.
// Same(tree.New(1), tree.New(1)) should return true,
// and Same(tree.New(1), tree.New(2)) should return false.
package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t.Left != nil {
			walker(t.Left)
		}
		ch <- t.Value
		if t.Right != nil {
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

// // Same determines whether the trees
// // t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// var num int
	var t1Values []int
	var t2Values []int

	for {

		n1, ok1 := <-ch1
		n2, ok2 := <-ch2

		t1Values = append(t1Values, n1)
		t2Values = append(t2Values, n2)

		if !ok1 && !ok2 {
			break
		}
	}

	sort.Ints(t1Values)
	sort.Ints(t2Values)

	return Equal(t1Values, t2Values)
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	isSame := Same(tree.New(10), tree.New(10))
	fmt.Println(isSame)
}
