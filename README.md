# m2c

> go for 2x2 complex matrices

[![GoReference](https://pkg.go.dev/badge/github.com/fibo/m2c)](https://pkg.go.dev/github.com/fibo/m2c)

## Description

Complex matrices 2x2 (a.k.a. **m2c**) are a wonderful land to explore.
They are matrices of the form

```
     | a  b |
     | c  d |
```

Where a, b, c, d are complex numbers. I found Golang as the perfect
language to implement them since it has `complex128` data type, yeah!

## Examples

First of all import `m2c`.

```go
import "github.com/fibo/m2c"
```

Create an identity matrix

```go
var id = m2c.Matrix{1, 0, 0, 1}
fmt.Println(id) // {(1+0i) (0+0i) (0+0i) (1+0i)}
```

Multiply two matrices.

```go
var a = m2c.Matrix{1, 0, 0, 1}
var b = m2c.Matrix{1, 1 + 1i, 0, 1 - 1i}

var c = m2c.Mul(a, b)
fmt.Println(c) // {(1+0i) (1+1i) (1+1i) (1-1i)}
```

Invert a matrix.

```go
var d = m2c.Matrix{2i, 0, 0, 1}

var invD, err = m2c.Inv(d)

if err != nil {
	log.Fatal(err)
}

fmt.Println(invD) // {(0-0.5i) (-0+0i) (-0+0i) (1+0i)}
```

## License

[MIT](https://fibo.github.io/mit-license/)

