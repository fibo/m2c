// Package m2c implements complex matrices 2x2.
package m2c

import (
	"fmt"
	"math/cmplx"
)

var (
	conj = cmplx.Conj
)

// A matrix with two rows, two columns and complex numbers as values.
type Matrix struct {
	A, B, C, D complex128
}

// CannotInvertMatrixError complains when the argument of the inversion operator is a matrix with determinant equal to zero.
type CannotInvertMatrixError struct {
	matrix Matrix
}

func (e *CannotInvertMatrixError) Error() string {
	return fmt.Sprintf("Cannot invert a matrix with determinant=0\n%v", e.matrix)
}

// Add two matrices.
func Add(l Matrix, r Matrix) Matrix {
	return Matrix{l.A + r.A, l.B + r.B, l.C + r.C, l.D + r.D}
}

// Conj returns the conjugated matrix.
func Conj(m Matrix) Matrix {
	return Matrix{conj(m.A), conj(m.B), conj(m.C), conj(m.D)}
}

// Eq checks if two matrices are equal.
func Eq(l Matrix, r Matrix) bool {
	return (l.A == r.A) && (l.B == r.B) && (l.C == r.C) && (l.D == r.D)
}

// Det computes matrix determinant.
func Det(m Matrix) complex128 {
	return m.A*m.D - m.B*m.C
}

// Inv inverts given matrix respect to multiplication.
// If it has determinant equal to zero, an error will be returned
// as second argument.
func Inv(m Matrix) (Matrix, error) {
	var det = Det(m)

	if 0 == det {
		return Matrix{}, &CannotInvertMatrixError{m}
	}

	return Matrix{m.D / det, -m.B / det, -m.C / det, m.A / det}, nil
}

// Mul multiplies two matrices. This operator is not commutative.
func Mul(l Matrix, r Matrix) Matrix {
	return Matrix{l.A*r.A + l.B*r.C, l.A*r.B + l.B*r.D, l.C*r.A + l.D*r.B, l.C*r.B + l.D*r.D}
}

// Neg computes matrix inverse respect to addition.
func Neg(m Matrix) Matrix {
	return Matrix{-m.A, -m.B, -m.C, -m.D}
}

// Scalar multiplies matrix by a complex number.
func Scalar(m Matrix, c complex128) Matrix {
	return Matrix{m.A * c, m.B * c, m.C * c, m.D * c}
}

// Sub subtracts two matrices.
func Sub(l Matrix, r Matrix) Matrix {
	return Matrix{l.A - r.A, l.B - r.B, l.C - r.C, l.D - r.D}
}

// T returns the transposed matrix.
func T(m Matrix) Matrix {
	return Matrix{m.A, m.C, m.B, m.D}
}
