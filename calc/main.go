package main

import (
	calc "calc"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Использование: %s выражение\n", os.Args[0])
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
	var exp calc.Expression
	input := flag.Arg(0)
	ans, err := exp.Calc(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(ans)
}
