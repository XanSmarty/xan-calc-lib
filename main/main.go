package main

import (
	"errors"
	"fmt"
	"io"
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
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
	return &Handler{stdout, calculator}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.stdout, "%d", result)
	if err != nil {
		return fmt.Errorf("%w: %w", errOutputWriter, err)
	}
	return nil
}

var (
	errWrongArgCount   = errors.New("usage: calc <a> <b>")
	errInvalidArgument = errors.New("invalid argument")
	errOutputWriter    = errors.New("output failure")
)
