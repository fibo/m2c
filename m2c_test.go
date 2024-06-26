package m2c

import (
	"testing"
)

var m1 = Matrix{0, 1, 1i, 0}
var m2 = Matrix{1 + 1i, 1 - 0.5i, 2 + 2i, 2 - 1i}
var one = Matrix{1, 0, 0, 1} // identity matrix
var minusOne = Matrix{-1, 0, 0, -1}
var j = Matrix{0, 1, -1, 0} // symplectic matrix
var zero = Matrix{0, 0, 0, 0}

func TestAdd(t *testing.T) {
	cases := []struct {
		a Matrix
		b Matrix
		c Matrix
	}{
		{one, zero, one},
	}

	for _, matrix := range cases {
		got := Add(matrix.a, matrix.b)

		if got != matrix.c {
			t.Errorf("Add(%v, %v) == %v, want %v", matrix.a, matrix.b, got, matrix.c)
		}
	}
}

func TestConj(t *testing.T) {
	cases := []struct {
		in   Matrix
		conj Matrix
	}{
		{one, one},
		{m1, Matrix{0, 1, -1i, 0}},
		{m2, Matrix{1 - 1i, 1 + 0.5i, 2 - 2i, 2 + 1i}},
	}

	for _, matrix := range cases {
		got := Conj(matrix.in)

		if got != matrix.conj {
			t.Errorf("Conj(%v) == %v, want %v", matrix.in, got, matrix.conj)
		}
	}
}

func TestEq(t *testing.T) {
	if !Eq(m1, m1) {
		t.Errorf("Matrix %v is not equal to itself", m1)
	}
}

func TestDet(t *testing.T) {
	cases := []struct {
		in  Matrix
		det complex128
	}{
		{one, 1},
		{m1, -1i},
		{m2, 0},
		{Mul(m1, m1), -1},
	}

	for _, matrix := range cases {
		got := Det(matrix.in)

		if got != matrix.det {
			t.Errorf("Det(%v) == %v, want %v", matrix.in, got, matrix.det)
		}
	}
}

func TestInv(t *testing.T) {
	cases := []struct {
		a   Matrix
		inv Matrix
	}{
		{one, one},
	}

	for _, matrix := range cases {
		got, err := Inv(matrix.a)

		if err != nil {
			t.Error(err)
		}

		if got != matrix.inv {
			t.Errorf("Inv(%v) == %v, want %v", matrix.a, got, matrix.inv)
		}
	}
}

func TestMul(t *testing.T) {
	m1Inv, _ := Inv(m1)

	cases := []struct {
		a Matrix
		b Matrix
		c Matrix
	}{
		{one, one, one},
		{m1, one, m1},
		{m2, one, m2},
		{m1, m2, Matrix{2 + 2i, 2 - 1i, -1 + 1i, 0.5 + 1i}},
		{m1, m1Inv, one},
	}

	for _, matrix := range cases {
		got := Mul(matrix.a, matrix.b)

		if got != matrix.c {
			t.Errorf("Mul(%v, %v) == %v, want %v", matrix.a, matrix.b, got, matrix.c)
		}
	}
}

func TestNeg(t *testing.T) {
	cases := []struct {
		in  Matrix
		neg Matrix
	}{
		{one, minusOne},
		{zero, zero},
	}

	for _, matrix := range cases {
		got := Neg(matrix.in)

		if got != matrix.neg {
			t.Errorf("Neg(%v) == %v, want %v", matrix.in, got, matrix.neg)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		a Matrix
		b Matrix
		c Matrix
	}{
		{one, one, zero},
	}

	for _, matrix := range cases {
		got := Sub(matrix.a, matrix.b)

		if got != matrix.c {
			t.Errorf("Sub(%v, %v) == %v, want %v", matrix.a, matrix.b, got, matrix.c)
		}
	}
}

func TestT(t *testing.T) {
	cases := []struct {
		a Matrix
		b Matrix
	}{
		{one, one},
		{zero, zero},
		{j, Matrix{0, -1, 1, 0}},
	}

	for _, matrix := range cases {
		got := T(matrix.a)

		if got != matrix.b {
			t.Errorf("T(%v) == %v, want %v", matrix.a, got, matrix.b)
		}
	}
}
