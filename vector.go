package matrix

import "errors"

// Vector will return a column vector consisting of the inputted values. To get a row vector, immediately follow with the Transpose function.
func Vector(nums ...float64) (*MatrixStruct, error) {
	return Matrix(len(nums), 1, nums)
}

// IsVector does a simple check to determine if a given matrix is actually a vector.
func (m MatrixStruct) IsVector() bool {
	if m.Rows == 1 || m.Columns == 1 {
		return true
	}
	return false
}

// Unit return a unit vector if the given matrix is actually a vector.
func (m MatrixStruct) Unit() (*MatrixStruct, error) {
	if m.IsVector() {
		return m.ScalarMultiply(1 / m.Normal("2")), nil
	}
	return nil, errors.New("Matrix is not a Vector")
}
