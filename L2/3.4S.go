// Напишите программу, которая будет выполнять поиск товара по его названию. Пользователь должен ввести название товара или его часть, после чего программа должна вывести в консоль цену данного товара в формате:

// Название товара: цена товара

// Доступные товары и их цены:
// "Клавиатура JZ9": 19200
// "Наушники N45": 9600
// "Смартфон S10": 55000
// Если пользователь введет название товара, которого нет в списке, программа должна вывести сообщение:

// Товар "название_введенного_товара" не найден.

// Поиск должен быть нечувствителен к регистру. Например, если пользователь введет "наУШНИКИ n45", программа должна вернуть цену 9600.

// Примеры:
// Введите название товара: наУШНИКИ n45
// Наушники N45: 9600

// Введите название товара: сМАРТ
// Смартфон S10: 55000

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	keyboard := "Клавиатура JZ9"
	headphone := "Наушники N45"
	phone := "Смартфон S10"
	pricekeyboard := 19200
	priceheadphone := 9600
	pricephone := 55000

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите название товара: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка ввода: %s", err)
	}
	input := scanner.Text()

	switch {
	case strings.Contains(strings.ToLower(keyboard), strings.ToLower(input)):
		fmt.Printf("%s: %d", keyboard, pricekeyboard)
	case strings.Contains(strings.ToLower(headphone), strings.ToLower(input)):
		fmt.Printf("%s: %d", headphone, priceheadphone)
	case strings.Contains(strings.ToLower(phone), strings.ToLower(input)):
		fmt.Printf("%s: %d", phone, pricephone)
	default:
		fmt.Printf("Товар %s не найден", input)
	}
}
