/*
	Задача: Напишите программу, которая для заданных значений a и b вычисляет площадь поверхности и объем тела, образованного вращением эллипса, заданного уравнением:

x^2/a^2 +y^2/b^2 = 1
вокруг оси Ox.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Print("Enter a numb: ")
	fmt.Scan(&a)
	fmt.Print("Enter b numb: ")
	fmt.Scan(&b)
	// Шаг для интегрирования
	step := 0.0001

	// Инициализация переменных для хранения площади поверхности и объема
	surfaceArea := 0.0
	volume := 0.0

	// Интегрирование для вычисления площади поверхности и объема
	for x := -a; x <= a; x += step {
		y := b * math.Sqrt(1-math.Pow(x/a, 2)) // Функция описывающая эллипс

		surfaceArea += 2 * math.Pi * y * step
		volume += math.Pi * math.Pow(y, 2) * step
	}

	fmt.Println("Surface area:", surfaceArea)
	fmt.Println("Volume:", volume)
}
