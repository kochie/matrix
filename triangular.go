package matrix

import (
	"errors"
)

// IsUpperTriangular will check if the matrix is an Upper Triangular matrix. That is all the elements below the diagonal are zero.
func (m MatrixStruct) IsUpperTriangular() bool {
	for i := 0; i < m.Columns; i++ {
		for j := 0; j < i; j++ {
			if m.Elements[i*m.Columns+j] != 0 {
				return false
			}
		}
	}
	return true
}

// IsLoewrTriangular will check if the matrix is a Lower Triangular matrix. That is all the elements above the diagonal are zero.
func (m MatrixStruct) IsLowerTriangular() bool {
	for i := 0; i < m.Columns; i++ {
		for j := i + 1; j < m.Rows; j++ {
			if m.Elements[i*m.Columns+j] != 0 {
				return false
			}
		}
	}
	return true
}

// TriangleInverse will preform Guass-Jordan elimination on the selected triangular matrix to determine it's inverse.
func (m MatrixStruct) TriangleInverse() (*MatrixStruct, error) {
	if m.Columns != m.Rows {
		return nil, errors.New("Not a square matrix")
	}

	inverse, _ := Eye(m.Columns, m.Rows)

	if m.IsUpperTriangular() {
		for k := m.Rows - 1; k >= 0; k-- {
			t, _ := m.GetValue(k, k)
			inverse.SetValue(k, k, 1/t)
			for i := k - 1; i >= 0; i-- {
				a, _ := m.Minor(i, i, i, k)
				b, _ := inverse.Minor(i, k, k, k)
				c, _ := a.Multiply(b)
				d, _ := m.GetValue(i, i)
				e, _ := c.ScalarMultiply(-1/d).GetValue(0, 0)
				inverse.SetValue(i, k, e)
			}
		}
		return inverse, nil
	}
	if m.IsLowerTriangular() {
		for k := 0; k < m.Rows; k++ {
			t, _ := m.GetValue(k, k)
			inverse.SetValue(k, k, 1/t)
			for i := k + 1; i < m.Rows; i++ {
				a, _ := m.Minor(i, i, k, i-1)
				b, _ := inverse.Minor(k, i-1, k, k)
				c, _ := a.Multiply(b)
				d, _ := m.GetValue(i, i)
				e, _ := c.ScalarMultiply(-1/d).GetValue(0, 0)
				inverse.SetValue(i, k, e)
			}
		}
		return inverse, nil
	}

	return nil, errors.New("Not a triangular matrix")
}
