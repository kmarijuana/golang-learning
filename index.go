package main

import "fmt"

func main() {
	fmt.Println("HELLO")
	foo()

	fmt.Println("EAMPLE")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm foo")
}
func bar() {
	fmt.Println("barsssdasdsdsadsaa")
}
