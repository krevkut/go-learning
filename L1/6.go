// Программист Арсений любит равные числа, а неравные не любит.
// Напишите программу, которая сравнивает три числа и выводит ‘Максимальное равенство‘, если все числа равны или ‘Не равны‘ в противном случае.

package main

import "fmt"

func main() {
	var chisl1, chisl2, chisl3 float64
	fmt.Scan(&chisl1, &chisl2, &chisl3)

	if chisl1 == chisl2 && chisl2 == chisl3 {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}

}
