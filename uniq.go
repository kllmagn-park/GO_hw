package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsEqual(line1 string, line2 string, ignoreCase bool) bool {
	if ignoreCase {
		line1, line2 = strings.ToLower(line1), strings.ToLower(line2)
	}
	return line1 == line2
}

func readStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func writeStdout(text string) {
	fmt.Print(text)
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func writeFile(filename string, text string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		log.Fatal(err)
	}
}

func filterLines(lines []string, ignoreCase bool) []string {
	var linesFiltered []string
	if len(lines) < 1 {
		return lines
	}
	curLine := lines[0]
	linesFiltered = append(linesFiltered, curLine)
	for i := 1; i < len(lines); i++ {
		if !IsEqual(curLine, lines[i], ignoreCase) {
			curLine = lines[i]
			linesFiltered = append(linesFiltered, lines[i])
		}
	}
	return linesFiltered
}

func countLine(lines []string, targetLine string, ignoreCase bool) int {
	counter := 0
	for _, line := range lines {
		if IsEqual(line, targetLine, ignoreCase) {
			counter++
		}
	}
	return counter
}

func countLines(lines []string, ignoreCase bool) []string {
	var linesCounted []string
	for i := 0; i < len(lines); i++ {
		lineNum := countLine(lines, lines[i], ignoreCase)
		linesCounted = append(linesCounted, strconv.Itoa(lineNum)+" "+lines[i])
	}
	return linesCounted
}

func filterRepeated(lines []string, ignoreCase bool) []string {
	var linesRepeated []string
	for i := 0; i < len(lines); i++ {
		lineNum := countLine(lines, lines[i], ignoreCase)
		if lineNum > 1 {
			linesRepeated = append(linesRepeated, lines[i])
		}
	}
	return linesRepeated
}

func filterUnique(lines []string, ignoreCase bool) []string {
	var linesUnique []string
	for i := 0; i < len(lines); i++ {
		lineNum := countLine(lines, lines[i], ignoreCase)
		if lineNum == 1 {
			linesUnique = append(linesUnique, lines[i])
		}
	}
	return linesUnique
}

func main() {

	useCount := flag.Bool("c", false, "Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	outputRepeated := flag.Bool("d", false, "Вывести только те строки, которые повторились во входных данных.")
	outputUnique := flag.Bool("u", false, "Вывести только те строки, которые не повторились во входных данных.")
	//numFields := flag.Int("f", 0, "Не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.")
	//numChars := flag.Int("s", 0, "Не учитывать первые num_chars символов в строке. При использовании вместе с параметром.")
	ignoreCase := flag.Bool("i", false, "Не учитывать регистр букв.")

	flag.Parse()

	var lines []string
	if flag.NArg() == 0 {
		lines = readStdin()
	} else if flag.NArg() <= 2 {
		inputFile := flag.Arg(0)
		lines = readFile(inputFile)
	} else {
		log.Fatal("Слишком много аргументов.")
		return
	}

	// обработка строк
	if *useCount {
		lines = countLines(lines, *ignoreCase)
	} else if *outputRepeated {
		lines = filterRepeated(lines, *ignoreCase)
	} else if *outputUnique {
		lines = filterUnique(lines, *ignoreCase)
	} else {
		lines = filterLines(lines, *ignoreCase)
	}
	// окончание обработки

	text := strings.Join(lines, "\n")
	if flag.NArg() <= 1 {
		writeStdout(text)
	} else {
		writeFile(flag.Arg(1), text)
	}
}
