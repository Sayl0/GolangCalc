package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[rune]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,
}
var intToRoman = map[int]string{
	100: "C", 90: "XC", 50: "L", 40: "XL",
	10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
}

func romanToArabic(roman string) int {
	total := 0
	prevuValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		char := rune(roman[i])
		value := romanToInt[char]

		if value < prevuValue {
			total -= value
		} else {
			total += value
		}
		prevuValue = value
	}
	return total
}

func arabicToRoman(num int) string {
	roman := ""

	for num > 0 {
		arabKey := 0
		for value := range intToRoman {
			if value <= num && value > arabKey { // XVI
				arabKey = value
			}
		}
		roman += intToRoman[arabKey]
		num -= arabKey
	}
	return roman
}

func isRoman(str string) bool {
	for i := 0; i < len(str); i++ {
		char := rune(str[i])
		_, exists := romanToInt[char]

		if !exists {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n	Калькулятор принимает на вход два арабских или римских числа от 1 до 10 включительно.\n	Введите уравнение через пробел: ")

	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSpace(input)
	parts := strings.Fields(input) // ["1", "+", "2"]

	if len(parts) != 3 {
		if len(parts) > 3 {
			panic("	Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		} else {
			panic("	Выдача паники, так как строка не является математической операцией.")
		}
	}

	isRoman1 := isRoman(parts[0])
	isRoman2 := isRoman(parts[2])
	if isRoman1 != isRoman2 {
		panic("	Выдача паники, так как используются одновременно разные системы счисления.")
	}

	var num1, num2 int
	if isRoman1 {
		num1 = romanToArabic(strings.ToUpper(parts[0]))
		num2 = romanToArabic(strings.ToUpper(parts[2]))
	} else {
		num1, _ = strconv.Atoi(parts[0])
		num2, _ = strconv.Atoi(parts[2])
	}

	operator := parts[1]
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("	Деление на ноль!")
		}
		result = num1 / num2
	default:
		panic("	Разрешено использовать только +, -, * или /")
	}

	if isRoman1 {
		if result == 0 {
			panic("	В римской системе счисления нет нуля!")
		} else if result < 0 {
			panic("	Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else {
			fmt.Println("	Ответ: ", arabicToRoman(result))
		}
	} else {
		fmt.Println("	Ответ: ", result)
	}
}
