package main

import "fmt"

func main() {
	var i = 10
	fmt.Println("Loops in Go lang")
	fmt.Println("Demonstrating if condition")
	if i > 10 {
		fmt.Println("i is greater than 10 ")
	} else if i < 10 {
		fmt.Println(" i is greater than 1")
	} else {
		fmt.Println("i = 10 ")
	}
	fmt.Println("Demonstrating for loop")
	for j := 0; j < i; j++ {
		fmt.Println(j)
	} // go has no while so use for as while

	fmt.Println("Demonstrating switch case")
	switch i {
	case 1:
		println(" i = 1 ")
	case 2:
		println(" i = 2 ")
	case 10:
		fmt.Println("i = 10 ")
	}
} // main method ends here
