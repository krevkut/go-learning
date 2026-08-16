// Программист Арсений часто ходит в магазин, но не доверяет подсчётам на кассе.

// Помогите ему и напишите программу, которой подаётся на вход количество продуктов, скидка по карте лояльности (скидка распространяется на все продукты) и цена каждого продукта, а программа выводит итоговую стоимость покупок в чеке с учётом скидок.

// Формат ввода
// Количество N

// Целое число

// N float64 чисел

// Формат вывода
// float64

package main

import (
	"fmt"
	"math"
)

func main() {
	var quantity, percent int
	fmt.Scanln(&quantity)
	fmt.Scanln(&percent)
	var total float64

	for range quantity {
		var price float64
		fmt.Scanln(&price)
		total += price
	}
	res := total * float64(100-percent) / 100
	res = math.Round(res*1000) / 1000
	fmt.Println(res)
}
