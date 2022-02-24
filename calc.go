package calc

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"log"
	"os"
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

func (s *Stack) Clear() {
	*s = nil
}

var operPriorities = map[string]int{
	"*/": 2,
	"+-": 1,
}

type Expression struct { stack Stack }

// Произвести форматирование выражения для дальнейшего парсинга.
func (e* Expression) format(formula string) (string, error) {
	formula = strings.ReplaceAll(formula, " ", "")
	return formula, nil
}

// Получить числовой приоритет соответствующего оператора.
func (e* Expression) getPriority(operator string) int {
	for k, v := range operPriorities {
		if strings.Contains(k, operator) {
			return v
		}
	}
	return -1
}

// Произвести преобразование выражение из инфиксной в постфиксную запись.
func (e* Expression) parse(formula string) (string, error) {
	form := []rune(formula)
	var parsed string = ""
	for i := 0; i < len(form); i++ {
		ch := form[i]
		chs := string(ch)
		priority := e.getPriority(chs)
		if unicode.IsNumber(ch) {
			// любое однозначное число
			parsed += chs
		} else if priority != -1 {
			// любой оператор
			val, exists := e.stack.Pop()
			if exists {
				e.stack.Push(val)
			}
			if !exists || e.getPriority(val) < priority || e.stack.Contains("(") {
				e.stack.Push(chs)
			} else {
				for true {
					val, exists := e.stack.Pop()
					if !exists {
						break
					}
					if val == "(" || val == ")" || e.getPriority(val) < priority {
						e.stack.Push(val)
						break
					}
					parsed += val
				}
				e.stack.Push(chs)
			}
		} else if chs == "(" {
			// левая скобка
			e.stack.Push(chs)
		} else if chs == ")" {
			// правая скобка
			for true {
				val, exists := e.stack.Pop()
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
			return "", fmt.Errorf("Неизвестный символ: " + chs)
		}
	}
	exists := true
	for exists {
		val, exists := e.stack.Pop()
		if !exists {
			break
		}
		parsed += val
	}
	return parsed, nil
}

// Вычислить заданную постфиксную запись выражения.
func (e *Expression) eval(formula string) (float64, error) {
	defer func() {
		e.stack.Clear()
	}()
	form := []rune(formula)
	for i := 0; i < len(form); i++ {
		ch := form[i]
		chs := string(ch)
		if e.getPriority(chs) != -1 {
			val1Str, exists1 := e.stack.Pop()
			val2Str, exists2 := e.stack.Pop()
			if (!exists1 || !exists2) {
				return 0, fmt.Errorf("Неверный формат постфиксной записи.")
			}
			val1, err := strconv.ParseFloat(val1Str, 64)
			if (err != nil) {
				return 0, err
			}
			val2, err := strconv.ParseFloat(val2Str, 64)
			if (err != nil) {
				return 0, err
			}
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
				return 0, fmt.Errorf("Неизвестный оператор: "+chs)
			}
			val3 := strconv.FormatFloat(v3, 'f', -1, 64)
			e.stack.Push(val3)
		} else if unicode.IsNumber(ch) {
			e.stack.Push(chs)
		} else {
			return 0, fmt.Errorf("Неизвестный символ постфиксной записи: " + chs)
		}
	}
	res, _ := strconv.ParseFloat(e.stack[0], 64)
	return res, nil
}

// Вычислить заданное выражение в строковом формате.
func (e *Expression) Calc(formula string) (float64, error) {
	formula, err := e.format(formula)
	if (err != nil) {
		return 0, err
	}
	parsed, err := e.parse(formula)
	if (err != nil) {
		return 0, err
	}
	return e.eval(parsed)
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s expression\n", os.Args[0])
     	flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
    	os.Exit(1)
	} else if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}
	var exp Expression
	input := flag.Arg(0)
	ans, err := exp.Calc(input)
	if (err != nil) {
		log.Fatal(err)
		return
	}
	fmt.Println(ans)
}
