# mymatrix


![GitHub top language](https://img.shields.io/github/languages/top/Zachary2237/mymatrix)&nbsp;
![go_version](https://img.shields.io/badge/go%20version-1.18-green)

`mymatrix` is a very small matrix calculation library.

It implements some `row vector` and `matrix` construction and operation methods.

# Quick Start

Fisrt, create a demo and import the mymatrix:

```shell
$ go mod init demo

$ go get github.com/Zachary2237/mymatrix
```

# Usage

Here is a simple application example.

```go
package main

import (
	"example/matrix"
	"fmt"
)

func init() {}

func main() {
	m1 := matrix.Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := matrix.Matrix([][]float64{[]float64{10, 20, 30}, []float64{40, 50, 60}})
	m3, _ := m1.Add(m2)
	fmt.Println(m3)
}

```

# API

The `row vector` operation implements the methods in the following interface.

```go
type IRowVec interface {
	Add(rv2 RowVector) (rv3 RowVector, err error)   //Vector addition
	Minus(rv2 RowVector) (rv3 RowVector, err error) //Vector subtraction
	Dot(rv2 RowVector) (dot float64, err error)     //vector dot product
	Cross(rv2 RowVector) (rv3 RowVector, err error) //vector fork multiplication(2-D or 3-D matrix only)
	Length() (l float64)                            //Vector modulus length
	Transpose() (cv ColumnVector)                   //Vector transpose
}
```

The `matrix` operations implement the methods in the following interface.

```go
type IMat interface {
	Add(m2 Matrix) (m3 Matrix, err error)    //Matrix Addition
	Minus(m2 Matrix) (m3 Matrix, err error)  //Matrix addition and subtraction
	MatMul(m2 Matrix) (m3 Matrix, err error) //Matrix multiplication
	Transpose() (m2 Matrix)                  //Matrix transpose
}
```
