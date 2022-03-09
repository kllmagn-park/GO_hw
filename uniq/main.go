package main

import (
	"uniq"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func readStdin() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}
	return output, nil
}

func writeStdout(text string) {
	fmt.Print(text)
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return output, err
	}

	return output, nil
}

func writeFile(filename string, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		return err
	}

	return nil
}

func main() {

	useCount := flag.Bool("c", false, "Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	outputRepeated := flag.Bool("d", false, "Вывести только те строки, которые повторились во входных данных.")
	outputUnique := flag.Bool("u", false, "Вывести только те строки, которые не повторились во входных данных.")
	numFields := flag.Int("f", 0, "Не учитывать первые numFields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.")
	numChars := flag.Int("s", 0, "Не учитывать первые numChars символов в строке.")
	ignoreCase := flag.Bool("i", false, "Не учитывать регистр букв.")
	showHelp := flag.Bool("h", false, "Показать это сообщение.")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	var lines []string
	var err error
	if flag.NArg() == 0 {
		lines, err = readStdin()
		if (err != nil) {
			log.Fatal(err)
			os.Exit(1)
		}
	} else if flag.NArg() <= 2 {
		inputFile := flag.Arg(0)
		lines, err = readFile(inputFile)
		if (err != nil) {
			log.Fatal(err)
			os.Exit(1)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	options := uniq.Options{*useCount, *outputRepeated, *outputUnique, *numFields, *numChars, *ignoreCase}

	lines, err = uniq.Uniq(lines, options)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	text := strings.Join(lines, "\n")
	if flag.NArg() <= 1 {
		writeStdout(text)
	} else {
		writeFile(flag.Arg(1), text)
	}
}
