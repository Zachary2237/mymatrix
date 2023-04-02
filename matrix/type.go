package matrix

//定义向量和矩阵数据的类型

type RowVector []float64
type ColumnVector [][1]float64
type Matrix [][]float64

type IVecMat interface {
	GetShape() (s [2]int) //获取维度
	Mul(c float64)        //标量乘法
}

//行向量的运算方法
type IRowVec interface {
	Add(rv2 RowVector) (rv3 RowVector, err error)   //向量加法
	Minus(rv2 RowVector) (rv3 RowVector, err error) //向量减法
	Dot(rv2 RowVector) (dot float64, err error)     //向量点乘
	Cross(rv2 RowVector) (rv3 RowVector, err error) //向量叉乘
	Length() (l float64)                            //向量长度
	Transpose() (cv ColumnVector)                   //转置
}

type IMat interface {
	Add(m2 Matrix) (m3 Matrix, err error)
	Minus(m2 Matrix) (m3 Matrix, err error)
	MatMul(m2 Matrix) (m3 Matrix, err error)
	Transpose() (m2 Matrix)
}
