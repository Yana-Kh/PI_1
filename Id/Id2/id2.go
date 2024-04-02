// Вариант 28 (4). Даны два числа. Найти их сумму, разность, произведение,
// а также частное от деления первого числа на второе

package main

import (
	"fmt"
	"os"
)

func main() {
	var first float32
	var second float32
	fmt.Print("Enter first numb:")
	fmt.Fscan(os.Stdin, &first)
	fmt.Print("Enter second numb:")
	fmt.Fscan(os.Stdin, &second)
	fmt.Println("Sum:", first+second)
	fmt.Println("Subt:", first-second)
	fmt.Println("Mult:", first*second)
	fmt.Println("Div:", first/second)

}
