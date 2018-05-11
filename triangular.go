package matrix

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
func (m MatrixStruct) TriangleInverse(*MatrixStruct, error) {
	if m.IsUpperTriangular() {

	}
}
