package main

import (
	"fmt"
	"math"
)

func main() {

	// Коэффициенты квадратного уравнения.
	var a, b, c float64

	fmt.Println("Решение квадратного уравнения.")

	fmt.Println("Введите коэффициент a:")
	fmt.Scan(&a)

	fmt.Println("Введите коэффициент b:")
	fmt.Scan(&b)

	fmt.Println("Введите коэффициент c:")
	fmt.Scan(&c)

	// Формула дискриминанта квадратного уравнения
	D := math.Pow(b, 2) - 4*a*c

	if D < 0 {
		fmt.Println("Корней нет.")
	} else if D == 0 {
		x := (-b) / (2 * a)
		fmt.Println("Уравнение имеет единственный корень, равный", x)
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Уравнение имеет два корня, которые равны", x1, "и", x2, "соответственно")
	}
}
