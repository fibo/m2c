package m2c

import (
	"math/cmplx"
)

type Matrix struct {
	a complex128
	b complex128
	c complex128
	d complex128
}

func Conj(in Matrix) Matrix {
	out := Matrix{}

	out.a = in.a
	out.b = in.b
	out.c = in.c
	out.d = in.d

	return out
}
