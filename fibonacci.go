package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1
	return func() int {
		num := prev
		prev, curr = curr, prev+curr
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f())
	}
}
