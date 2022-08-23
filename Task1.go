package main

import "fmt"

func main() {

	var x, y float64 // Числа, которые вводит пользователь

	fmt.Println("Определение координатной плоскости точки.")

	fmt.Println("Введите x:") // Ввод х
	fmt.Scan(&x)

	fmt.Println("Введите y:") // Ввод у
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка принадлежит к первой четверти.")
	}

	if x < 0 && y > 0 {
		fmt.Println("Точка принадлежит ко второй четверти.")
	}

	if x < 0 && y < 0 {
		fmt.Println("Точка принадлежит к третьей четверти.")
	}

	if x > 0 && y < 0 {
		fmt.Println("Точка принадлежит к четвертой четверти.")
	}
}
