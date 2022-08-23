package main

import (
	"fmt"
)

func main() {

	// Стоимость товара и номиналы трёх монет
	var productPrice, denom1, denom2, denom3 int

	fmt.Println("Сумма без сдачи.")

	fmt.Println("Введите стоимость товара:")
	fmt.Scan(&productPrice)

	fmt.Println("Введите номинал первой монеты:")
	fmt.Scan(&denom1)

	fmt.Println("Введите номинал второй монеты:")
	fmt.Scan(&denom2)

	fmt.Println("Введите номинал третьей монеты:")
	fmt.Scan(&denom3)

	if productPrice == denom1 || productPrice == denom2 || productPrice == denom3 ||
		productPrice == (denom1+denom2) || productPrice == (denom1+denom3) ||
		productPrice == (denom3+denom2) || productPrice == (denom1+denom2+denom3) {
		fmt.Println("Вы сможете оплатить товар без сдачи.")
	} else {
		fmt.Println("Вы не сможете оплатить товар без сдачи.")
	}

}
