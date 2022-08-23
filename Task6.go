package main

import (
	"fmt"
)

func main() {

	// Четырёхзначный номер билета
	var ticketNum int

	fmt.Println("Счастливый билет.")

	fmt.Println("Введите номер билета:")
	fmt.Scan(&ticketNum)

	// Разбиваем номер билета на символы
	// Первый символ
	firstDigit := ticketNum % 10
	// Второй символ
	secondDigit := (ticketNum / 10) % 10
	// Третий символ
	thirdDigit := (ticketNum / 100) % 10
	// Второй символ
	fourthDigit := (ticketNum / 1000) % 10

	if firstDigit == fourthDigit && secondDigit == thirdDigit {
		fmt.Println("Зеркальный билет.")
	} else if firstDigit+secondDigit == thirdDigit+fourthDigit {
		fmt.Println("Счастливый билет")
	} else {
		fmt.Println("Обычный билет")
	}

}
