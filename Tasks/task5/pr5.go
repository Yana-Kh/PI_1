/* Задача: По данному целому числу, найдите его квадрат*/
package main

import (
  "fmt"
)

func main(){
  
   var a, result int
   fmt.Print("Введите число: ")
   fmt.Scan(&a) // считаем переменную 'a' с консоли
   result = a * a
   fmt.Println("Возведение в квадрат", result)
  }
