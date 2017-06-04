package m2c

import (
	"math/cmplx"
)

type Matrix struct {
	a, b, c, d complex128
}

var conj = cmplx.Conj

func Add(l Matrix, r Matrix) Matrix {
	return Matrix{l.a + r.a, l.b + r.b, l.c + r.c, l.d + r.d}
}

func Conj(m Matrix) Matrix {
	return Matrix{conj(m.a), conj(m.b), conj(m.c), conj(m.d)}
}

func Eq(l Matrix, r Matrix) bool {
	return (l.a == r.a) && (l.b == r.b) && (l.c == r.c) && (l.d == r.d)
}

func Det(m Matrix) complex128 {
	return m.a*m.d - m.b*m.c
}

func Id() Matrix {
	return Matrix{1, 0, 0, 1}
}

func Inv(m Matrix) Matrix {
	var det = Det(m)

	return Matrix{m.d / det, -m.b / det, -m.c / det, m.a / det}
}

func Mul(l Matrix, r Matrix) Matrix {
	return Matrix{l.a*r.a + l.b*r.c, l.a*r.b + l.b*r.d, l.c*r.a + l.d*r.b, l.c*r.b + l.d*r.d}
}

func Sub(l Matrix, r Matrix) Matrix {
	return Matrix{l.a - r.a, l.b - r.b, l.c - r.c, l.d - r.d}
}
