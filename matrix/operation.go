package matrix

import (
	"errors"
	"fmt"
	"math"
)

//定义对向量和矩阵的操作

// GetShape 获取行向量维度,行向量的底层为[]float64
func (rv1 RowVector) GetShape() (s [2]int) {
	s[0] = 1
	s[1] = len(rv1)
	return
}

// Mul 计算行向量与标量相乘
func (rv1 *RowVector) Mul(c float64) {
	l := rv1.GetShape()[1]
	rvTemp := make([]float64, l, l)
	copy(rvTemp, *rv1)
	for i := 0; i < l; i++ {
		rvTemp[i] *= c
	}
	copy(*rv1, rvTemp)

	//为什么不用这种方式呢?
	//for i := 0; i < l; i++ {
	//	(*rv1)[i] *= c
	//}
}

// NewRowVector 根据指定长度l初始化向量,若l小于0,返回err
func NewRowVector(l int) (rv1 RowVector, err error) {
	if l <= 0 {
		err = errors.New("the dimension of row vectors must > 0")
		return
	}

	rv1 = make([]float64, l, l)
	return
}

// Add 两个行向量的加法
func (rv1 RowVector) Add(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}

	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, _ := range rv1 {
		rv3[i] = rv1[i] + rv2[i]
	}

	return
}

// Minus 两个行向量的减法
func (rv1 RowVector) Minus(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}

	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, _ := range rv1 {
		rv3[i] = rv1[i] - rv2[i]
	}
	return
}

// Dot 两个行向量的点乘运算
func (rv1 RowVector) Dot(rv2 RowVector) (dot float64, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}

	for i, _ := range rv1 {
		dot += rv1[i] * rv2[i]
	}
	return
}

// Cross 二维和三维向量的叉乘
func (rv1 RowVector) Cross(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}

	dim := rv1.GetShape()[1]
	if dim != 2 && dim != 3 {
		err = errors.New("we can only calculate the 2 or 3 dimensions row vector")
		return
	}

	rv3, _ = NewRowVector(3)
	switch dim {
	case 2:
		rv3[0] = 0
		rv3[1] = 0
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	case 3:
		rv3[0] = rv1[1]*rv2[2] - rv1[2]*rv2[1]
		rv3[1] = rv1[2]*rv2[0] - rv1[0]*rv2[2]
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	}
	return
}

// Length 计算行向量的模长
func (rv1 RowVector) Length() (l float64) {
	for _, v := range rv1 {
		l += v * v
	}
	l = math.Sqrt(l)
	return
}

// Transpose 行向量的转置,输出为列向量
func (rv1 RowVector) Transpose() (cv ColumnVector) {
	//Go语言会对ColumnVector和[][1]float64两个类型进行隐式转换
	cv = make([][1]float64, rv1.GetShape()[1], rv1.GetShape()[1])
	for i, v := range rv1 {
		cv[i][0] = v
	}
	return
}

// GetShape 获取矩阵维度
func (m1 Matrix) GetShape() (s [2]int) {
	s[0] = len(m1)
	s[1] = len(m1[0])
	return
}

// Mul 矩阵与标量c的乘法运算
func (m1 *Matrix) Mul(c float64) {
	row, col := m1.GetShape()[0], m1.GetShape()[1]
	mTemp := make([][]float64, row, row)
	copy(mTemp, *m1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mTemp[i][j] *= c
		}
	}

	copy(*m1, mTemp)
}

// NewMatrix 矩阵的初始化方法,r为行数,c为列数
func NewMatrix(r, c int) (mat Matrix, err error) {
	if r <= 0 || c <= 0 {
		err = errors.New("rows and columns of the matrix must > 0")
		return
	}

	mat = make([][]float64, r, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c, c)
	}
	return
}

// Add 两个矩阵的加法
func (m1 Matrix) Add(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}

	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, err = NewMatrix(r, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return
}

// Minus 两个矩阵的减法
func (m1 Matrix) Minus(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}

	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, err = NewMatrix(r, c)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return
}

// MatMul 矩阵乘法
func (m1 Matrix) MatMul(m2 Matrix) (m3 Matrix, err error) {
	r1, c1 := m1.GetShape()[0], m1.GetShape()[1]
	r2, c2 := m2.GetShape()[0], m2.GetShape()[1]

	if c1 != r2 {
		err = errors.New("the shape of the two matrix are not match")
		return
	}

	m3, _ = NewMatrix(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			v1 := RowVector(m1[i])
			v2 := make([]float64, r2, r2)
			for k := 0; k < r2; k++ {
				v2[k] = m2[k][j]
			}

			m3[i][j], _ = v1.Dot(v2)
		}
	}
	return
}

// Transpose 矩阵的转置运算
func (m1 Matrix) Transpose() (m2 Matrix) {
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m2, _ = NewMatrix(c, r)
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			m2[i][j] = m1[j][i]
		}
	}

	return
}
