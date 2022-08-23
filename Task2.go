package main

import "fmt"

func main() {

	var num1, num2, num3 float64 // Числа, которые вводит пользователь

	fmt.Println("Проверить, что одно из чисел — положительное.")

	fmt.Println("Введите первое число:")
	fmt.Scan(&num1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&num2)

	fmt.Println("Введите третье число:")
	fmt.Scan(&num3)

	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Одно или более из введеных вами чисел положительное")
	} else {
		fmt.Println("Ни одно из введеных вами чисел не является положительным")
	}

}
