package main

import "fmt"

// this is comment

func main() {
	var name string = "Sherlock"
	var number int
	const const_number = 3.1416
	fmt.Println("Hello world")
	number = 1 + 1
	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(const_number)
	name = "Sherlock Holmes"
	fmt.Println(name)
	var (
		a = 100
		b = 200
		c = 300
	)
	fmt.Println(`Decalred mulltiple variables in Go fashion but all must be used 
and none should remain unused so adding them`)
	fmt.Println(a + b + c)
	fmt.Println("Multiline string is enclosed between `` ")
	var x = 100
	fmt.Println(x)
}
