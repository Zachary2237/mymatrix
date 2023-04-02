package matrix

//对向量和矩阵的测试
import (
	"testing"
)

func TestRowVector_GetShape(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3, 4.4})
	t.Log(rv1.GetShape())
}

func TestRowVector_Mul(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3, 4.4})
	rv1.Mul(2.0)
	t.Log("rv1 = ", rv1)
}

func TestNewRowVector(t *testing.T) {
	t.Log(NewRowVector(5))
}

func TestRowVector_Add(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Add(rv2)
	t.Log("rv3 = ", rv3)
}

func TestRowVector_Minus(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Minus(rv2)
	t.Log("rv3 = ", rv3)
}

func TestRowVector_Cross(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 4})
	rv2 := RowVector([]float64{14, 17, 16})
	rv3, _ := rv1.Cross(rv2)
	t.Log("rv3 = ", rv3)
}

func TestRowVector_Length(t *testing.T) {
	rv := RowVector([]float64{3, 4})
	l := rv.Length()
	t.Log("length = ", l)
}

func TestRowVector_Transpose(t *testing.T) {
	rv := RowVector([]float64{14, 17, 16})
	cv := rv.Transpose()
	t.Log("ColomnVector = ", cv)
}

func TestMatrix_GetShape(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("shape of m = ", m.GetShape())
}

func TestMatrix_Mul(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("m = ", m)
	m.Mul(2)
	t.Log("m = ", m)
}

func TestNewMatrix(t *testing.T) {
	m, _ := NewMatrix(3, 4)
	t.Log("m = ", m)
}

func TestMatrix_Add(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{14, 12, 10}, []float64{16, 17, 19}})
	m3, _ := m1.Add(m2)
	t.Log("m3 = ", m3)
}

func TestMatrix_Minus(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{14, 12, 10}, []float64{16, 17, 19}})
	m3, _ := m1.Minus(m2)
	t.Log("m3 = ", m3)
}

func TestMatrix_MatMul(t *testing.T) {
	m1 := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	m2 := Matrix([][]float64{[]float64{11, 12}, []float64{13, 14}, []float64{15, 16}})
	m3, _ := m1.MatMul(m2)
	t.Log("m3 = ", m3)
}

func TestMatrix_Transpose(t *testing.T) {
	m := Matrix([][]float64{[]float64{1, 2, 3}, []float64{4, 5, 6}})
	t.Log("m = ", m)
	mt := m.Transpose()
	t.Log("mt = ", mt)
}

func BenchmarkRowVector_GetShape(b *testing.B) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3, 4.4})
	for i := 0; i < b.N; i++ {
		rv1.GetShape()
	}
}
