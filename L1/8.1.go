// Программист Арсений мечтает создать машину времени. Конечно, работа не простая, но кто-то же должен это сделать. Арсений решил начать с малого и хочет написать код, который будет выводить информацию, на сколько лет в будущее он переместился.
// Напишите программу, которая считывает с консоли время будущего и настоящего и возвращает указание, сколько лет «назад» был нынешний год. Строка должна выглядеть так: X year ago

// Формат ввода
// Время в формате yyyy-MM-dd

// Время в формате yyyy-MM-dd

package main

import (
	"fmt"
	"time"
)

func main() {
	var datef, datet string
	fmt.Scanln(&datef)
	fmt.Scanln(&datet)
	daform := "2006-01-02"

	formatteddatef, _ := time.Parse(daform, datef)
	a := formatteddatef.Year()

	formatteddatet, _ := time.Parse(daform, datet)
	b := formatteddatet.Year()
	fmt.Println(a-b, "year ago")
}
