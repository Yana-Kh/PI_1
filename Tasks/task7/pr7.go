/* Задача: Дано неотрицательное целое число. Найдите число 
десятков (то есть вторую цифру справа). На вход дается 
натуральное число N, не превосходящее 10000. Выведите одно 
целое число - число десятков.*/
package main
import "fmt"
func main(){

   var a, result int
   fmt.Print("Введите число: ")
   fmt.Scan(&a) // считаем переменную 'a' с консоли
   result = (a % 100 - a % 10) / 10
   fmt.Println("Количество дсятков: ", result)
  }