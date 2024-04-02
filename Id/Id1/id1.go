// Вариант 28 (8). Объем и площадь цилиндрической банки:
// Задайте переменные для радиуса и высоты цилиндрической банки.
// Рассчитайте и выведите объем и полную поверхность цилиндра

package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	var height float64
	fmt.Print("Enter first numb:")
	fmt.Scan(&radius)
	fmt.Print("Enter second numb:")
	fmt.Scan(&height)

	fmt.Println("Volume: ", math.Pi*math.Pow(radius, 2)*height)
	fmt.Println("Surface area: ", 2*math.Pi*math.Pow(radius, 2)+2*math.Pi*radius*height)

}
