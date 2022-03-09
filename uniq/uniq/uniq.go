package uniq

import (
	"strconv"
	"strings"
)

// Параметры uniq.
type Options struct {
	UseCount       bool // Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.
	OutputRepeated bool // Вывести только те строки, которые повторились во входных данных.
	OutputUnique   bool // Вывести только те строки, которые не повторились во входных данных.
	NumFields      int  // Не учитывать первые numFields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.
	NumChars       int  // Не учитывать первые numChars символов в строке. При использовании вместе с параметром.
	IgnoreCase     bool // Не учитывать регистр букв.
}

func isEqual(line1 string, line2 string, options Options) (bool, error) {
	if options.IgnoreCase {
		line1, line2 = strings.ToLower(line1), strings.ToLower(line2)
	}
	if options.NumChars > 0 {
		var numChars1, numChars2 int
		if options.NumChars > len(line1) {
			numChars1 = len(line1)
		} else {
			numChars1 = options.NumChars
		}
		if options.NumChars > len(line2) {
			numChars2 = len(line2)
		} else {
			numChars2 = options.NumChars
		}
		line1 = line1[numChars1:]
		line2 = line2[numChars2:]
	}
	return line1 == line2, nil
}

func filterLines(lines []string, options Options) ([]string, error) {
	var linesFiltered []string
	if len(lines) < 1 {
		return lines, nil
	}
	curLine := lines[0]
	linesFiltered = append(linesFiltered, curLine)
	for _, line := range lines {
		equal, err := isEqual(curLine, line, options)
		if err != nil {
			return linesFiltered, err
		}
		if !equal {
			curLine = line
			linesFiltered = append(linesFiltered, line)
		}
	}
	return linesFiltered, nil
}

func countLine(lines []string, targetLine string, options Options) (int, error) {
	counter := 0
	for _, line := range lines {
		equal, err := isEqual(line, targetLine, options)
		if err != nil {
			return 0, err
		}
		if equal {
			counter++
		}
	}
	return counter, nil
}

func countLines(lines []string, options Options) ([]string, error) {
	var linesCounted []string
	for _, line := range lines {
		lineNum, err := countLine(lines, line, options)
		if err != nil {
			return linesCounted, err
		}
		linesCounted = append(linesCounted, strconv.Itoa(lineNum)+" "+line)
	}
	return linesCounted, nil
}

func filterRepeated(lines []string, options Options) ([]string, error) {
	var linesRepeated []string
	for _, line := range lines {
		lineNum, err := countLine(lines, line, options)
		if err != nil {
			return linesRepeated, nil
		}
		if lineNum > 1 {
			linesRepeated = append(linesRepeated, line)
		}
	}
	return linesRepeated, nil
}

func filterUnique(lines []string, options Options) ([]string, error) {
	var linesUnique []string
	for _, line := range lines {
		lineNum, err := countLine(lines, line, options)
		if err != nil {
			return linesUnique, err
		}
		if lineNum == 1 {
			linesUnique = append(linesUnique, line)
		}
	}
	return linesUnique, nil
}

func cutFields(lines []string, numFields int) (res []string, cutLines []string) {
	res = lines
	counter := 0
	for i, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			cutLines = lines[:i]
			res = lines[i:]
			counter++
		}
		if counter == numFields {
			break
		}
	}
	return
}

func GetDefaultOptions() Options {
	options := Options{}
	options.UseCount = false
	options.OutputRepeated = false
	options.OutputUnique = false
	options.NumFields = 0
	options.NumChars = 0
	options.IgnoreCase = false
	return options
}

func Uniq(lines []string, options Options) (res []string, err error) {
	var cutLines []string
	if options.NumFields > 0 {
		lines, cutLines = cutFields(lines, options.NumFields)
	}
	if options.UseCount {
		res, err = countLines(lines, options)
	} else if options.OutputRepeated {
		res, err = filterRepeated(lines, options)
	} else if options.OutputUnique {
		res, err = filterUnique(lines, options)
	} else {
		res, err = filterLines(lines, options)
	}
	res = append(cutLines, res...)
	return
}
