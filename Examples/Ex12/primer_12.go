package main

import (
	"fmt"
	"math" // Add this line to import the math package
)

func main() {
	var result = math.Abs(-5.67) //абсолютное значение
	fmt.Println(result)

	result = math.Ceil(5.67) // округление вверх до ближайшнго целого
	fmt.Println(result)

	result = math.Floor(5.67) // округление вниз до ближайшнго целого
	fmt.Println(result)

	result = math.Sqrt(16) // квадратный корень
	fmt.Println(result)

	result = math.Pow(2, 3) // возведение в степень
	fmt.Println(result)

	var sinValue = math.Sin(math.Pi / 2) // синус
	fmt.Println(sinValue)

	result = math.Log(10) // логарифм (натуральный)
	fmt.Println(result)

	var maxVal = math.Max(3, 7) // максимальное значение
	var minVal = math.Min(3, 7) // минимальное значение
	fmt.Println(maxVal, minVal)

	result = math.Mod(10, 3) // остаток от деления
	fmt.Println(result)

	result = math.Round(5.67) // округление до ближайшего целого
	fmt.Println(result)

	var posInf = math.Inf(1)  // Возвращает положительную бесконечность
	var negInf = math.Inf(-1) // Возвращает отрицательную бесконечность
	fmt.Println(posInf, negInf)

	var nan = math.NaN() // Возвращает "Not a Number" (NaN)
	fmt.Println(nan)

	result = math.Exp(2) // Возвращает экспоненту в степени
	fmt.Println(result)

	result = math.Exp2(3) // Возвращает экспоненту в степени 2
	fmt.Println(result)

	result = math.Expm1(1) // Возвращает экспоненту в степени -1
	fmt.Println(result)

	result = math.Log10(100) // Возвращает десятичный логарифм числа
	fmt.Println(result)

	result = math.Log2(8) // Возвращает двоичный логарифм числа
	fmt.Println(result)

	result = math.Log1p(1) // Возвращает натуральный логарифм числа
	fmt.Println(result)

	var isNegative = math.Signbit(-5) // Возвращает true, если число отрицательное
	fmt.Println(isNegative)
}
