package main

import "fmt"

func main() {
	fmt.Print("hello, world")
	fmt.Print("hello, world")

	fmt.Println("\n")

	fmt.Println("hello, world \n")
	fmt.Print("hello, world")

	fmt.Println("\n")

	fmt.Print("Ivan", 27)
	fmt.Println("Ivan", 27)
	fmt.Print(33, 27)

	fmt.Println("\n")

	name := "Ivan"
	age := 27
	fmt.Println("My name is", name, "and I am", age, "years old.")

}
