// Exercise: Fibonacci closure
// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	elemPre := 0
	elemPost := 1
	sum := 0
	return func() int {
		sum, elemPre, elemPost = elemPre, elemPost, elemPre+elemPost
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
