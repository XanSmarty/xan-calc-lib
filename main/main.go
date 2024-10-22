package main

import (
	"fmt"
	"os"
	"strconv"

	calc "github.com/XanSmarty/xan-calc-lib"
)

type Calculator interface {
	Calculate(a, b int) int
}

func main() {
	args := os.Args[1:]
	a, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	var calculator calc.Addition
	result := calculator.Calculate(a, b)
	fmt.Println(result)
}
