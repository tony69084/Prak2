package main

import "fmt"

func main() {

	var num1, num2, num3 float64 // Числа, которые вводит пользователь

	fmt.Println("Проверить, что есть совпадающие числа.")

	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)

	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)

	if num1 == num2 || num2 == num3 || num3 == num1 {
		fmt.Println("Два числа или более совпадают.")
	} else {
		fmt.Println("Все числа разные")
	}

}
