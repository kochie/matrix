//matrix is a library full of matrix functions.
package matrix

import (
	"errors"
	"fmt"
	"math"
)

// MatrixStruct defines a basic dense matrix structure that can store values in a slice.
type MatrixStruct struct {
	Rows, Columns, Capacity int
	Elements                []float64
}

func (m MatrixStruct) shortestDimension() int {
	return int(math.Min(float64(m.Rows), float64(m.Columns)))
}

// Matrix will return a MatrixStruct structure containing the Elements given in the list. This function will also parse the Elements and check for input errors.
func Matrix(rows int, columns int, elements []float64) (*MatrixStruct, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}

	capacity := rows * columns

	if len(elements) > capacity {
		return nil, errors.New("More Elements than supported in matrix dimensions")
	}
	return &MatrixStruct{
		Capacity: capacity,
		Rows:     rows,
		Columns:  columns,
		Elements: elements,
	}, nil
}

// Eye will create an identity matrix with the dimensions given.
func Eye(rows, columns int) (*MatrixStruct, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}

	elements := make([]float64, rows*columns)

	s := int(math.Min(float64(rows), float64(columns)))

	for i := 0; i < s; i++ {
		elements[i*columns+i] = 1
	}
	return Matrix(rows, columns, elements)
}

// Zeros will return a new matrix containing zeros in all Elements. The size of the matrix will be determined by the dimensions given.
func Zeros(rows int, columns int) (*MatrixStruct, error) {
	return Matrix(rows, columns, make([]float64, rows*columns))
}

// Ones will return a new matrix containing ones in all Elements. The size of the matrix will be determined by the dimensions given.
func Ones(rows int, columns int) (*MatrixStruct, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}
	elements := make([]float64, rows*columns)
	for i := range elements {
		elements[i] = 1
	}
	return Matrix(rows, columns, elements)
}

// Add will return a new matrix that has the sum of the current matrix and the input matrix. Will also check for dimension errors.
func (m MatrixStruct) Add(n *MatrixStruct) (*MatrixStruct, error) {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		return nil, errors.New("The dimensions of the matricies must agree!")
	}

	newElements := make([]float64, m.Capacity)

	for i := range m.Elements {
		newElements[i] = m.Elements[i] + n.Elements[i]
	}

	return Matrix(m.Rows, m.Columns, newElements)
}

// Subtract will return a new matrix that has the difference of the current matrix and the input matrix. Will also check for dimension errors.
func (m MatrixStruct) Subtract(n *MatrixStruct) (*MatrixStruct, error) {
	if m.Rows != n.Rows || m.Columns != n.Columns {
		return nil, errors.New("The dimensions of the matricies must agree!")
	}

	newElements := make([]float64, m.Capacity)

	for i := range m.Elements {
		newElements[i] = m.Elements[i] - n.Elements[i]
	}

	return Matrix(m.Rows, m.Columns, newElements)
}

// Transpose will return a new matrix that is the transpose of the current matrix.
func (m MatrixStruct) Transpose() *MatrixStruct {
	newElements := make([]float64, m.Capacity)
	for i := 0; i < m.Columns; i++ {
		for j := 0; j < m.Rows; j++ {
			newElements[i*m.Rows+j] = m.Elements[j*m.Columns+i]
		}
	}

	matrix, _ := Matrix(m.Columns, m.Rows, newElements)
	return matrix
}

// Print will print the matrix to the console.
func (m MatrixStruct) Print() {
	fmt.Println()
	for i := 0; i < m.Rows; i++ {
		fmt.Println(m.Elements[i*m.Columns : (i+1)*m.Columns])
	}
	fmt.Println()
}

// Multiply returns a new matrix that is the matrix product of the current and input matrix. The order of the matrix multiplication is right handed, meaning the output is m*n.
func (m MatrixStruct) Multiply(n *MatrixStruct) (*MatrixStruct, error) {
	if m.Columns != n.Rows {
		return nil, errors.New("matrix dimensions do not agree")
	}

	newElements := make([]float64, m.Rows*n.Columns)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < n.Columns; j++ {
			element := float64(0)
			for k := 0; k < n.Rows; k++ {
				element += m.Elements[i*m.Columns+k] * n.Elements[k*n.Columns+j]
			}
			newElements[i*n.Columns+j] = element
		}
	}

	return Matrix(m.Rows, n.Columns, newElements)
}

// Trace returns the sum of the diagonal values of the matrix.
func (m MatrixStruct) Trace() float64 {
	var sum float64

	for i := 0; i < m.shortestDimension(); i++ {
		sum += m.Elements[i*m.Columns+i]
	}

	return sum
}

// ScalarMultiply returns a new matrix that is the original matrix multiplied by the input scalar.
func (m MatrixStruct) ScalarMultiply(s float64) *MatrixStruct {
	newElements := make([]float64, len(m.Elements))
	for i, elem := range m.Elements {
		newElements[i] = elem * s
	}

	matrix, _ := Matrix(m.Rows, m.Columns, newElements)
	return matrix
}

func (m MatrixStruct) norm2() (norm float64) {
	norm = 0
	for i := 0; i < m.Columns; i++ {
		sum := float64(0)
		for j := 0; j < m.Rows; j++ {
			sum += math.Pow(m.Elements[j*m.Columns+i], float64(2))
		}
		norm = math.Max(norm, sum)
	}
	return math.Sqrt(norm)
}

// Clone returns a new matrix that is an exact copy of the selected matrix.
func (m MatrixStruct) Clone() *MatrixStruct {
	s := make([]float64, len(m.Elements))
	copy(s, m.Elements)
	matrix, _ := Matrix(m.Rows, m.Columns, s)
	return matrix
}

// Normal returns the selected matrix normal.
func (m MatrixStruct) Normal(normType string) float64 {
	switch normType {
	case ("2"):
		{
			return m.norm2()
		}
	default:
		{
			return m.norm2()
		}
	}
}

// Minor will return a new matrix whos values are the minor of the selected matrix at the given coordinates.
func (m MatrixStruct) Minor(column1, column2, row1, row2 int) (*MatrixStruct, error) {
	if row1 > row2 || column1 > column2 {
		return nil, errors.New("Matrix index mismatch")
	}

	if row2 > m.Rows-1 || column2 > m.Columns-1 {
		return nil, errors.New("Matrix index is mismatched")
	}

	var elem []float64

	for i := column1; i <= column2; i++ {
		for j := row1; j <= row2; j++ {
			elem = append(elem, m.Elements[i*m.Columns+j])
		}
	}

	return Matrix(column2-column1+1, row2-row1+1, elem)
}

// MinorUpdate will update the selected area in a matrix minor with new values and return a new matrix with the replaced numbers.
func (m MatrixStruct) MinorUpdate(y1, y2, x1, x2 int, n *MatrixStruct) (*MatrixStruct, error) {
	if y1 > y2 || x1 > x2 {
		return nil, errors.New("Matrix index mismatch")
	}

	if y2 > m.Rows || x2 > m.Columns {
		return nil, errors.New("Matrix index is mismatched")
	}

	newMatrix := m.Clone()

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			newMatrix.Elements[j*m.Columns+i] = n.Elements[(j-y1)*n.Columns+(i-x1)]
		}
	}

	return newMatrix, nil
}

// GetValue will return the value of an element at a specific index in the matrix.
func (m MatrixStruct) GetValue(y, x int) (float64, error) {
	if y >= m.Rows || x >= m.Columns {
		return 0, fmt.Errorf("The indexes must be in the bounds of the matrix.\nMatrix is %dx%d, indexes are %d,%d", m.Rows, m.Columns, y, x)
	}

	return m.Elements[y*m.Columns+x], nil
}

// SetValue will set the value of an element at a specific index in the matrix.
func (m MatrixStruct) SetValue(y, x int, value float64) error {
	if y >= m.Rows || x >= m.Columns {
		return fmt.Errorf("The indexes must be in the bounds of the matrix.\nMatrix is %dx%d, indexes are %d,%d", m.Rows, m.Columns, y, x)
	}

	m.Elements[y*m.Columns+x] = value
	return nil
}

// IsEqual will determine if two matricies are the same shape and have the same values in the same places.
func (m MatrixStruct) IsEqual(n *MatrixStruct) bool {
	if m.Columns != n.Columns {
		return false
	}

	if m.Rows != n.Rows {
		return false
	}

	if len(m.Elements) != len(n.Elements) {
		return false
	}

	for i := 0; i < len(m.Elements); i++ {
		if m.Elements[i] != n.Elements[i] {
			return false
		}
	}

	return true
}

// QR will return the QR decomposition of the selected matrix using the householder decomposition.
func (m MatrixStruct) QR() (Q *MatrixStruct, R *MatrixStruct) {
	M := m.Rows
	N := m.Columns

	Q, _ = Eye(M, M)
	R = m.Clone()

	for k := 0; k < N; k++ {
		x, _ := R.Minor(k, M-1, k, k)
		y, _ := Zeros(M-k, 1)

		xTop, _ := x.GetValue(0, 0)

		if xTop < 0 {
			y.SetValue(0, 0, x.Normal("2"))
		} else {
			y.SetValue(0, 0, -x.Normal("2"))
		}

		v, _ := x.Subtract(y)
		u := v.ScalarMultiply(1 / v.Normal("2"))

		r1, _ := R.Minor(k, M-1, k, N-1)
		uDot, _ := u.Multiply(u.Transpose())
		houseHolder, _ := uDot.ScalarMultiply(2).Multiply(r1)
		temp, _ := r1.Subtract(houseHolder)
		R, _ = R.MinorUpdate(k, M-1, k, N-1, temp)

		q1, _ := Q.Minor(k, M-1, 0, M-1)
		houseHolder, _ = uDot.ScalarMultiply(2).Multiply(q1)
		temp, _ = q1.Subtract(houseHolder)
		Q, _ = Q.MinorUpdate(k, M-1, 0, M-1, temp)
	}

	Q = Q.Transpose()
	return
}

// IsSquare will return true if the matrix is square.
func (m MatrixStruct) IsSquare() bool {
	if m.Rows == m.Columns {
		return true
	}
	return false
}

// Prune will zero out small elements in a matrix to allow faster computations. Tolerance in set at 1e-10.
func (m MatrixStruct) Prune() (*MatrixStruct, error) {
	elements := make([]float64, m.Capacity)
	copy(elements, m.Elements)
	tol := 1e-10

	for i := 0; i < len(elements); i++ {
		if elements[i] < tol && elements[i] > -tol {
			elements[i] = 0
		}
	}

	return Matrix(m.Rows, m.Columns, elements)
}

// Inverse will compute the inverse matrix of a given matrix. FYI at the moment there is no check to see if the matrix is invertable.
func (m MatrixStruct) Inverse() (*MatrixStruct, error) {
	if !m.IsSquare() {
		return nil, errors.New("Not a square matrix")
	}

	Q, R := m.QR()
	Q_t := Q.Transpose()
	R, _ = R.Prune()
	R_inv, _ := R.TriangleInverse()
	inv, _ := R_inv.Multiply(Q_t)
	return inv, nil
}
