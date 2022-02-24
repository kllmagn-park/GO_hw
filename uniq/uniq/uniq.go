package uniq

import (
	"strconv"
	"strings"
)

func isEqual(line1 string, line2 string, ignoreCase bool) (bool, error) {
	if ignoreCase {
		line1, line2 = strings.ToLower(line1), strings.ToLower(line2)
	}
	return line1 == line2, nil
}

func filterLines(lines []string, ignoreCase bool) ([]string, error) {
	var linesFiltered []string
	if len(lines) < 1 {
		return lines, nil
	}
	curLine := lines[0]
	linesFiltered = append(linesFiltered, curLine)
	for i := 1; i < len(lines); i++ {
		equal, err := isEqual(curLine, lines[i], ignoreCase)
		if err != nil {
			return linesFiltered, err
		}
		if !equal {
			curLine = lines[i]
			linesFiltered = append(linesFiltered, lines[i])
		}
	}
	return linesFiltered, nil
}

func countLine(lines []string, targetLine string, ignoreCase bool) (int, error) {
	counter := 0
	for _, line := range lines {
		equal, err := isEqual(line, targetLine, ignoreCase)
		if err != nil {
			return 0, err
		}
		if equal {
			counter++
		}
	}
	return counter, nil
}

func countLines(lines []string, ignoreCase bool) ([]string, error) {
	var linesCounted []string
	for i := 0; i < len(lines); i++ {
		lineNum, err := countLine(lines, lines[i], ignoreCase)
		if err != nil {
			return linesCounted, err
		}
		linesCounted = append(linesCounted, strconv.Itoa(lineNum)+" "+lines[i])
	}
	return linesCounted, nil
}

func filterRepeated(lines []string, ignoreCase bool) ([]string, error) {
	var linesRepeated []string
	for i := 0; i < len(lines); i++ {
		lineNum, err := countLine(lines, lines[i], ignoreCase)
		if err != nil {
			return linesRepeated, nil
		}
		if lineNum > 1 {
			linesRepeated = append(linesRepeated, lines[i])
		}
	}
	return linesRepeated, nil
}

func filterUnique(lines []string, ignoreCase bool) ([]string, error) {
	var linesUnique []string
	for i := 0; i < len(lines); i++ {
		lineNum, err := countLine(lines, lines[i], ignoreCase)
		if err != nil {
			return linesUnique, err
		}
		if lineNum == 1 {
			linesUnique = append(linesUnique, lines[i])
		}
	}
	return linesUnique, nil
}

// Параметры uniq.
type Options struct {
	UseCount       bool // Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.
	OutputRepeated bool // Вывести только те строки, которые повторились во входных данных.
	OutputUnique   bool // Вывести только те строки, которые не повторились во входных данных.
	NumFields      int  // Не учитывать первые numFields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.
	NumChars       int  // Не учитывать первые numChars символов в строке. При использовании вместе с параметром.
	IgnoreCase     bool // Не учитывать регистр букв.
}

func Uniq(lines []string, options Options) (res []string, err error) {
	if options.UseCount {
		res, err = countLines(lines, options.IgnoreCase)
	} else if options.OutputRepeated {
		res, err = filterRepeated(lines, options.IgnoreCase)
	} else if options.OutputUnique {
		res, err = filterUnique(lines, options.IgnoreCase)
	} else {
		res, err = filterLines(lines, options.IgnoreCase)
	}
	return
}
