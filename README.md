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
import "m2c"
```

Get the identity matrix

```go
var id = m2c.I()
fmt.Printf("%v", id)
// {(1+0i) (0+0i) (0+0i) (1+0i)}
```

Multiply two matrices.

```go
var a = m2c.Matrix{1, 0, 0, 1}
var b = m2c.Matrix{1, 1+i, 0, 1-i}

var c = m2c.Mul(a, b)
```

Invert a matrix.

```go
var a = m2c.Matrix{2i, 0, 0, 1}

var invA, err = m2c.Inv(a)

if err != nil {
	log.Fatal(err)
}
```

## License

[MIT](https://fibo.github.io/mit-license/)

