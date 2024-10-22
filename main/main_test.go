package main

import (
	"bytes"
	"errors"
	"testing"

	calc "github.com/XanSmarty/xan-calc-lib"
)

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	if !errors.Is(err, errWrongArgCount) {
		t.Errorf("want: %v, got %v", errWrongArgCount, err)
	}
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "42"})
	if !errors.Is(err, errInvalidArgument) {
		t.Errorf("want: %v, got %v", errInvalidArgument, err)
	}
}

func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"42", "INVALID"})
	if !errors.Is(err, errInvalidArgument) {
		t.Errorf("want: %v, got %v", errInvalidArgument, err)
	}
}

func TestHandler_ResultWrittenToOutput(t *testing.T) {
	buf := &bytes.Buffer{}

	handler := NewHandler(buf, &calc.Addition{})
	err := handler.Handle([]string{"2", "3"})
	if err != nil {
		t.Errorf("want: %v, got %v", nil, err)
	}
	if buf.String() != "5" {
		t.Errorf("want: %v, got %v", "5", buf.String())
	}
}

func TestHandler_NoCalculatorProvided(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{})
}
