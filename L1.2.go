// Программист Арсений проверил работы своих учеников в кружке по информатике и выставил всем баллы. Осталось только сформировать финальную оценку.

// Напишите программу, которой подаётся на вход количество учеников и балл каждого ученика. На основании этих данных выведите оценку:

// 5, если балл от 90 до 100
// 4, если балл от 75 до 89
// 3, если балл от 50 до 74
// 2, если балл от 0 до 49
// Неверный балл в иных случаях
// Формат ввода
// Целое число N количество учеников

// N float64 чисел

// Формат вывода
// Целое число или строка

package main

import (
	"fmt"
)

func main() {
	var quantity int
	fmt.Scanln(&quantity)

	for range quantity {
		var quality float64
		fmt.Scanln(&quality)
		if (0 <= quality) && (quality <= 49) {
			fmt.Println(2)
		} else if (50 <= quality) && (quality <= 74) {
			fmt.Println(3)
		} else if (75 <= quality) && (quality <= 89) {
			fmt.Println(4)
		} else if (90 <= quality) && (quality <= 100) {
			fmt.Println(5)
		} else {
			fmt.Println("Неверный балл")
		}
	}
}
