package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целое число: ")
	input, err := reader.ReadString('\n') // Читаем до перевода строки

	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)          // Убираем лишние пробелы и \n
	fmt.Printf("Вы ввели число: %q\n", input) // %q для вывода в кавычках (если строка)
}
