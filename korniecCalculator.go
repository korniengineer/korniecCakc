package main // Инициализация пакета main

import ( // Импорт необходимых библиотек для проекта
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]int{ // Инициализация римских чисел для
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var convIntToRoman = [14]int{ // Числа для конвертации
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}
var a, b *int                          // Инициализация указателей
var operators = map[string]func() int{ // Инициализация операторов через функции
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const ( // Константы, содержащие в себе ошибки
	LOW = "Вывод ошибки, так как строка " +
		"не является математической операцией."
	HIGH = "Вывод ошибки, так как формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Вывод ошибки, так как используются " +
		"одновременно разные системы счисления."
	DIV = "Вывод ошибки, так как в римской системе " +
		"нет отрицательных чисел."
	ZERO  = "Вывод ошибки, так как в римской системе нет числа 0."
	RANGE = "Калькулятор умеет работать только с арабскими целыми " +
		"числами или римскими цифрами от 1 до 10 включительно"
)

func base(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(HIGH)
	case len(operator) < 1:
		panic(LOW)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(SCALE)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(RANGE)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(RANGE)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(ZERO)
	} else if romanResult < 0 {
		panic(DIV)
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
func main() {
	fmt.Println("Welcome to kata-calculator")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(s)))
	}
}
