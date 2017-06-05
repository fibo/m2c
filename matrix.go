// Complex matrices 2x2.
package m2c

import (
	"errors"
	"math/cmplx"
)

type Matrix struct {
	a, b, c, d complex128
}

var conj = cmplx.Conj

// Add two matrices.
func Add(l Matrix, r Matrix) Matrix {
	return Matrix{l.a + r.a, l.b + r.b, l.c + r.c, l.d + r.d}
}

// Return the conjugated matrix.
func Conj(m Matrix) Matrix {
	return Matrix{conj(m.a), conj(m.b), conj(m.c), conj(m.d)}
}

// Check it tow matrices are equal.
func Eq(l Matrix, r Matrix) bool {
	return (l.a == r.a) && (l.b == r.b) && (l.c == r.c) && (l.d == r.d)
}

// Compute matrix determinant.
func Det(m Matrix) complex128 {
	return m.a*m.d - m.b*m.c
}

// Returns the identity matrix.
func Id() Matrix {
	return Matrix{1, 0, 0, 1}
}

// Invert given matrix. If it has determinant equal to zero, the function
// will panic.
func Inv(m Matrix) Matrix {
	var det = Det(m)

	if 0 == det {
		panic(errors.New("Cannot invert a matrix with determinant=0"))
	}

	return Matrix{m.d / det, -m.b / det, -m.c / det, m.a / det}
}

// Multiply two matrices. This operator is not commutative.
func Mul(l Matrix, r Matrix) Matrix {
	return Matrix{l.a*r.a + l.b*r.c, l.a*r.b + l.b*r.d, l.c*r.a + l.d*r.b, l.c*r.b + l.d*r.d}
}

// Multiply matrix by a complex.
func Scalar(m Matrix, c complex128) Matrix {
	return Matrix{m.a * c, m.b * c, m.c * c, m.d * c}
}

// Subtract two matrices.
func Sub(l Matrix, r Matrix) Matrix {
	return Matrix{l.a - r.a, l.b - r.b, l.c - r.c, l.d - r.d}
}

// Return the transposed matrix.
func Tr(m Matrix) Matrix {
	return Matrix{m.a, m.c, m.b, m.d}
}
