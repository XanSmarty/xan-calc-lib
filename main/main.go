package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	calc "github.com/XanSmarty/xan-calc-lib"
)

type Calculator interface {
	Calculate(a, b int) int
}

func main() {

	handler := NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	stdout     *os.File
	calculator *calc.Addition
}

func NewHandler(stdout *os.File, calculator *calc.Addition) *Handler {
	return &Handler{stdout, calculator}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: calc <a> <b>")
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.stdout, "%d\n", result)
	return err
}
