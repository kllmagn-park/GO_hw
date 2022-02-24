package exeval

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // просто добавляем новое значение в конец стека
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // получить индекс самого верхнего элемента
		element := (*s)[index] // получить соответствующий индексу элемент
		*s = (*s)[:index]      // удалить элемент из стека путем обрезки слайса
		return element, true
	}
}

func (s *Stack) Contains(val string) bool {
	for _, v := range *s {
		if v == val {
			return true
		}
	}
	return false
}

var operPriorities = map[string]int{
	"*/": 2,
	"+-": 1,
}

// Произвести форматирование выражения для дальнейшего парсинга.
func format(formula string) string {
	formula = strings.ReplaceAll(formula, " ", "")
	return formula
}

// Получить числовой приоритет соответствующего оператора.
func getPriority(operator string) int {
	for k, v := range operPriorities {
		if strings.Contains(k, operator) {
			return v
		}
	}
	return -1
}

// Произвести преобразование выражение из инфиксной в постфиксную запись.
func parse(formula string) string {
	form := []rune(formula)
	var stack Stack
	var parsed string = ""
	for i := 0; i < len(form); i++ {
		ch := form[i]
		chs := string(ch)
		priority := getPriority(chs)
		if unicode.IsNumber(ch) {
			// любое однозначное число
			parsed += chs
		} else if priority != -1 {
			// любой оператор
			val, exists := stack.Pop()
			stack.Push(val)
			if !exists || getPriority(val) < priority || stack.Contains("(") {
				stack.Push(chs)
			} else {
				for true {
					val, exists := stack.Pop()
					if !exists {
						break
					}
					if val == "(" || val == ")" || getPriority(val) < priority {
						stack.Push(val)
						break
					}
					parsed += val
				}
				stack.Push(chs)
			}
		} else if chs == "(" {
			// левая скобка
			stack.Push(chs)
		} else if chs == ")" {
			// правая скобка
			for true {
				val, exists := stack.Pop()
				if !exists {
					break
				}
				if val == "(" {
					break
				}
				parsed += val
			}
		} else {
			// остальные символы
			panic("Неизвестный символ: " + chs)
		}
	}
	for i := 0; i < len(stack); i++ {
		parsed += string(stack[len(stack)-(i+1)])
	}
	return parsed
}

// Вычислить заданную постфиксную запись выражения.
func eval(formula string) float64 {
	form := []rune(formula)
	var stack []string
	for i := 0; i < len(form); i++ {
		ch := form[i]
		chs := string(ch)
		if getPriority(chs) != -1 {
			val1, _ := strconv.ParseFloat(stack[len(stack)-1], 64)
			val2, _ := strconv.ParseFloat(stack[len(stack)-2], 64)
			stack = stack[:len(stack)-2]
			var v3 float64
			switch ch {
			case '+':
				v3 = val1 + val2
			case '-':
				v3 = val2 - val1
			case '*':
				v3 = val1 * val2
			case '/':
				v3 = val2 / val1
			default:
				panic("Неизвестный оператор.")
			}
			val3 := strconv.FormatFloat(v3, 'f', -1, 64)
			stack = append(stack, val3)
		} else if unicode.IsNumber(ch) {
			stack = append(stack, chs)
		}
	}
	res, _ := strconv.ParseFloat(stack[0], 64)
	return res
}

// Вычислить заданное выражение в строковом формате.
func calc(formula string) float64 {
	formula = format(formula)
	parsed := parse(formula)
	fmt.Println(parsed)
	return eval(parsed)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Print("Недостаточно аргументов.")
	} else if flag.NArg() > 1 {
		fmt.Print("Слишком много аргументов.")
	}
	input := flag.Arg(0)
	ans := calc(input)
	fmt.Println(ans)
}
