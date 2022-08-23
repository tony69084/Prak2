package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	/* Число, которое программа угадывает и индикатор числа
	(больше, меньше или равно указанному программой) */
	var num, check int

	fmt.Println("Игра «Угадай число».")

	fmt.Println("Загадайте число от 1 до 10, когда загадете введите 1.")

	// Генератор случайного числа от 1 до 10
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(9) + 1

	// Первая попытка отгадать число
	fmt.Println("Вы загадали число", num)
	fmt.Println("Если да, введите 1, если оно больше введите 2, если меньше введите 3")
	fmt.Scan(&check)

	if check == 1 {
		fmt.Println("Ура, число угадано.")
		os.Exit(0)
	} else if check == 2 {
		if num+3 <= 10 {
			num += 3
			fmt.Println("Тогда может быть это число", num)
		} else if num+2 <= 10 {
			num += 2
			fmt.Println("Тогда может быть это число", num)
		} else {
			num++
			fmt.Println("Тогда может быть это число", num)
		}
	} else if check == 3 {
		if num-3 >= 1 {
			num -= 3
			fmt.Println("Тогда может быть это число", num)
		} else if num-2 >= 1 {
			num -= 2
			fmt.Println("Тогда может быть это число", num)
		} else {
			num--
			fmt.Println("Тогда может быть это число", num)
		}
	}
	// Записываем предыдущее значение индикатора в новую переменную
	checkpr1 := check

	// Вторая попытка отгадать число
	fmt.Println("Если да, введите 1, если оно больше введите 2, если меньше введите 3")
	fmt.Scan(&check)

	if check == 1 {
		fmt.Println("Ура, число угадано.")
		os.Exit(0)
	} else if checkpr1 == check {
		if check == 2 {
			if num+3 <= 10 {
				num += 3
				fmt.Println("Тогда может быть это число", num)
			} else if num+2 <= 10 {
				num += 2
				fmt.Println("Тогда может быть это число", num)
			} else {
				num++
				fmt.Println("Тогда может быть это число", num)
			}
		} else if check == 3 {
			if num-3 >= 1 {
				num -= 3
				fmt.Println("Тогда может быть это число", num)
			} else if num-2 >= 1 {
				num -= 2
				fmt.Println("Тогда может быть это число", num)
			} else {
				num--
				fmt.Println("Тогда может быть это число", num)
			}
		}
	} else if checkpr1 > check {
		if num+2 <= 10 {
			num += 2
			fmt.Println("Тогда может быть это число", num)
		} else {
			num++
			fmt.Println("Тогда может быть это число", num)
		}
	} else if checkpr1 < check {
		if num-2 >= 1 {
			num -= 2
			fmt.Println("Тогда может быть это число", num)
		} else {
			num--
			fmt.Println("Тогда может быть это число", num)
		}
	}

	// Третья попытка отгадать число
	checkpr2 := check
	fmt.Println("Если да, введите 1, если оно больше введите 2, если меньше введите 3")
	fmt.Scan(&check)

	if check == 1 {
		fmt.Println("Ура, число угадано.")
		os.Exit(0)
	} else if checkpr2 == check && checkpr1 == check {
		if check == 2 {
			if num+3 <= 10 {
				num += 3
				fmt.Println("Тогда может быть это число", num)
			} else if num+2 <= 10 {
				num += 2
				fmt.Println("Тогда может быть это число", num)
			} else {
				num++
				fmt.Println("Тогда может быть это число", num)
			}
		} else if check == 3 {
			if num-3 >= 1 {
				num -= 3
				fmt.Println("Тогда может быть это число", num)
			} else if num-2 >= 1 {
				num -= 2
				fmt.Println("Тогда может быть это число", num)
			} else {
				num--
				fmt.Println("Тогда может быть это число", num)
			}
		}
	} else if checkpr1 == checkpr2 && checkpr2 > check {
		{
			if num+2 <= 10 {
				num += 2
				fmt.Println("Тогда может быть это число", num)
			} else {
				num++
				fmt.Println("Тогда может быть это число", num)
			}
		}
	} else if checkpr1 == checkpr2 && checkpr2 < check {
		if num-2 >= 1 {
			num -= 2
			fmt.Println("Тогда может быть это число", num)
		} else {
			num--
			fmt.Println("Тогда может быть это число", num)
		}
	} else if checkpr1 == check && checkpr2 < check {
		num--
	} else if checkpr1 == check && checkpr2 > check {
		num++
	}

	// Четвертая попытка отгадать число
	fmt.Println("Если да, введите 1, если оно больше введите 2, если меньше введите 3")
	fmt.Scan(&check)

	if check == 1 {
		fmt.Println("Ура, число угадано.")
		os.Exit(0)
	} else {
		fmt.Println("Жаль, мне не удалось отгадать ваше число.")
	}
}
