package m2c

import (
	"fmt"
	"math/cmplx"
)

// Matrix with two rows, two columns and Complex numbers as values.
type Matrix struct {
	A, B, C, D complex128
}

// Matrix contructor.
func NewMatrix(a, b, c, d complex128) Matrix {
	return Matrix{A: a, B: b, C: c, D: d}
}

// CannotInvertMatrixError complains cause it is not possible to invert a matrix if determinant is zero.
type CannotInvertMatrixError struct {
	matrix Matrix
}

func (e *CannotInvertMatrixError) Error() string {
	return fmt.Sprintf("Cannot invert a matrix with determinant=0\n%v", e.matrix)
}

var conj = cmplx.Conj

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

// I returns the identity matrix.
func I() Matrix {
	return Matrix{1, 0, 0, 1}
}

// J returns the symplectic matrix.
func J() Matrix {
	return Matrix{0, 1, -1, 0}
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

// Zero returns the matrix with all zeros.
func Zero() Matrix {
	return Matrix{0, 0, 0, 0}
}
