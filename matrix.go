// Complex matrices 2x2.
package m2c

import (
	"fmt"
	"math/cmplx"
)

// A matrix with two rows and two columns with Complex values.
type Matrix struct {
	a, b, c, d complex128
}

func (m *Matrix) String() string {
	return fmt.Sprintf("⌈%v\t%v⌉\n⌊%v\t%v⌋", m.a, m.b, m.c, m.d)
}

// It is not possible to invert a matrix if determinant is zero.
type CannotInvertMatrixError struct {
	matrix Matrix
}

func (e *CannotInvertMatrixError) Error() string {
	return fmt.Sprintf("Cannot invert a matrix with determinant=0\n%v", e.matrix)
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
func ID() Matrix {
	return Matrix{1, 0, 0, 1}
}

// Invert given matrix. If it has determinant equal to zero, an error
// will be returned.
func Inv(m Matrix) (Matrix, error) {
	var det = Det(m)

	if 0 == det {
		return Matrix{}, &CannotInvertMatrixError{m}
	}

	return Matrix{m.d / det, -m.b / det, -m.c / det, m.a / det}, nil
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
