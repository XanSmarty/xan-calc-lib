package calc

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 1},
		{a: 4, b: 5, want: 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			addition := &Addition{}
			got := addition.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: -1},
		{a: 4, b: 5, want: -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d - %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			subtraction := &Subtraction{}
			got := subtraction.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 0},
		{a: 4, b: 5, want: 20},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d * %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			multiplication := &Multiplication{}
			got := multiplication.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 6, b: 3, want: 2},
		{a: 30, b: 6, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d / %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			division := &Division{}
			got := division.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_ByZeroPanics(t *testing.T) {
	division := &Division{}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("Division.ByZero did not panic")
		}
	}()
	division.Calculate(1, 0)
}
