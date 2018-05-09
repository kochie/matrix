//matrix is a library full of matrix functions.
package matrix

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// Matrix defines a basic dense matrix structure that can store values in a slice.
type Matrix struct {
	rows, columns, capacity int
	elements                []float64
}

func (m Matrix) shortestDimension() int {
	return int(math.Min(float64(m.rows), float64(m.columns)))
}

// NewMatrix will return a Matrix structure containing the elements given in the list. This function will also parse the elements and check for input errors.
func NewMatrix(rows int, columns int, elements []float64) (*Matrix, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}

	capacity := rows * columns

	if len(elements) > capacity {
		return nil, errors.New("More elements than supported in matrix dimensions")
	}
	return &Matrix{
		capacity: capacity,
		rows:     rows,
		columns:  columns,
		elements: elements,
	}, nil
}

// NewEyeMatrix will create an identity matrix with the dimensions given.
func NewEyeMatrix(rows int, columns int) (*Matrix, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}

	elements := make([]float64, rows*columns)

	s := int(math.Min(float64(rows), float64(columns)))

	for i := 0; i < s; i++ {
		elements[i*columns+i] = 1
	}
	return NewMatrix(rows, columns, elements)
}

// NewZerosMatrix will return a new matrix containing zeros in all elements. The size of the matrix will be determined by the dimensions given.
func NewZerosMatrix(rows int, columns int) (*Matrix, error) {
	return NewMatrix(rows, columns, make([]float64, rows*columns))
}

// NewOnesMatrix will return a new matrix containing ones in all elements. The size of the matrix will be determined by the dimensions given.
func NewOnesMatrix(rows int, columns int) (*Matrix, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("Incorrect matrix dimensions")
	}
	elements := make([]float64, rows*columns)
	for i := range elements {
		elements[i] = 1
	}
	return NewMatrix(rows, columns, elements)
}

// Add will return a new matrix that has the sum of the current matrix and the input matrix. Will also check for dimension errors.
func (m Matrix) Add(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.columns != n.columns {
		return nil, errors.New("The dimensions of the matricies must agree!")
	}

	newElements := make([]float64, m.capacity)

	for i := range m.elements {
		newElements[i] = m.elements[i] + n.elements[i]
	}

	return NewMatrix(m.rows, m.columns, newElements)
}

// Subtract will return a new matrix that has the difference of the current matrix and the input matrix. Will also check for dimension errors.
func (m Matrix) Subtract(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.columns != n.columns {
		return nil, errors.New("The dimensions of the matricies must agree!")
	}

	newElements := make([]float64, m.capacity)

	for i := range m.elements {
		newElements[i] = m.elements[i] - n.elements[i]
	}

	return NewMatrix(m.rows, m.columns, newElements)
}

// Transpose will return a new matrix that is the transpose of the current matrix.
func (m Matrix) Transpose() *Matrix {
	newElements := make([]float64, m.capacity)
	for i := 0; i < m.columns; i++ {
		for j := 0; j < m.rows; j++ {
			newElements[i*m.rows+j] = m.elements[j*m.columns+i]
		}
	}

	matrix, _ := NewMatrix(m.columns, m.rows, newElements)
	return matrix
}

// Print will print the matrix to the console.
func (m Matrix) Print() {
	fmt.Println()
	for i := 0; i < m.rows; i++ {
		fmt.Println(m.elements[i*m.columns : (i+1)*m.columns])
	}
	fmt.Println()
}

// Multiply returns a new matrix that is the matrix product of the current and input matrix. The order of the matrix multiplication is right handed, meaning the output is m*n.
func (m Matrix) Multiply(n *Matrix) (*Matrix, error) {
	if m.columns != n.rows {
		return nil, errors.New("matrix dimensions do not agree")
	}

	newElements := make([]float64, m.rows*n.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < n.columns; j++ {
			element := float64(0)
			for k := 0; k < n.rows; k++ {
				element += m.elements[i*m.columns+k] * n.elements[k*n.columns+j]
			}
			newElements[i*n.columns+j] = element
		}
	}

	return NewMatrix(m.rows, n.columns, newElements)
}

// Trace returns the sum of the diagonal values of the matrix.
func (m Matrix) Trace() float64 {
	var sum float64

	for i := 0; i < m.shortestDimension(); i++ {
		sum += m.elements[i*m.columns+i]
	}

	return sum
}

// ScalarMultiply returns a new matrix that is the original matrix multiplied by the input scalar.
func (m Matrix) ScalarMultiply(s float64) *Matrix {
	newElements := make([]float64, len(m.elements))
	for i, elem := range m.elements {
		newElements[i] = elem * s
	}

	matrix, _ := NewMatrix(m.rows, m.columns, newElements)
	return matrix
}

func (m Matrix) norm2() (norm float64) {
	norm = 0
	for i := 0; i < m.columns; i++ {
		sum := float64(0)
		for j := 0; j < m.rows; j++ {
			sum += math.Pow(m.elements[j*m.columns+i], float64(2))
		}
		norm = math.Max(norm, sum)
	}
	return math.Sqrt(norm)
}

// Clone returns a new matrix that is an exact copy of the selected matrix.
func (m Matrix) Clone() *Matrix {
	s := make([]float64, len(m.elements))
	copy(s, m.elements)
	matrix, _ := NewMatrix(m.rows, m.columns, s)
	return matrix
}

// Normal returns the selected matrix normal.
func (m Matrix) Normal(normType string) float64 {
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
func (m Matrix) Minor(row1, row2, column1, column2 int) (*Matrix, error) {
	if row1 > row2 || column1 > column2 {
		return nil, errors.New("Matrix index mismatch")
	}

	if row2 > m.rows-1 || column2 > m.columns-1 {
		return nil, errors.New("Matrix index is mismatched")
	}

	var elem []float64

	for i := column1; i <= column2; i++ {
		for j := row1; j <= row2; j++ {
			elem = append(elem, m.elements[i*m.columns+j])
		}
	}

	return NewMatrix(row2-row1+1, column2-column1+1, elem)
}

// MinorUpdate will update the selected area in a matrix minor with new values and return a new matrix with the replaced numbers.
func (m Matrix) MinorUpdate(y1, y2, x1, x2 int, n *Matrix) (*Matrix, error) {
	if y1 > y2 || x1 > x2 {
		return nil, errors.New("Matrix index mismatch")
	}

	if y2 > m.rows || x2 > m.columns {
		return nil, errors.New("Matrix index is mismatched")
	}

	newMatrix := m.Clone()

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			newMatrix.elements[j*m.columns+i] = n.elements[(j-y1)*n.columns+(i-x1)]
		}
	}

	return newMatrix, nil
}

// GetValue will return the value of an element at a specific index in the matrix.
func (m Matrix) GetValue(y, x int) (float64, error) {
	if y >= m.rows || x >= m.columns {
		return 0, fmt.Errorf("The indexes must be in the bounds of the matrix.\nMatrix is %dx%d, indexes are %d,%d", m.rows, m.columns, y, x)
	}

	return m.elements[y*m.columns+x], nil
}

// SetValue will set the value of an element at a specific index in the matrix.
func (m Matrix) SetValue(y, x int, value float64) error {
	if y >= m.rows || x >= m.columns {
		return fmt.Errorf("The indexes must be in the bounds of the matrix.\nMatrix is %dx%d, indexes are %d,%d", m.rows, m.columns, y, x)
	}

	m.elements[y*m.columns+x] = value
	return nil
}

// IsEqual will determine if two matricies are the same shape and have the same values in the same places.
func (m Matrix) IsEqual(n *Matrix) bool {
	if m.columns != n.columns {
		return false
	}

	if m.rows != n.rows {
		return false
	}

	if len(m.elements) != len(n.elements) {
		return false
	}

	for i := 0; i < len(m.elements); i++ {
		if m.elements[i] != n.elements[i] {
			return false
		}
	}

	return true
}

// QR will return the QR decomposition of the selected matrix.
func (a Matrix) QR() (Q *Matrix, R *Matrix) {
	m := a.rows
	n := a.columns

	Q, _ = NewEyeMatrix(m, m)
	R = a.Clone()

	for k := 0; k < n; k++ {
		x, err := R.Minor(k, m-1, k, k)
		x.Print()
		if err != nil {
			log.Fatal(err)
		}

		y, err := NewZerosMatrix(m-k, 1)
		if err != nil {
			log.Fatal(err)
		}

		xTop, err := x.GetValue(0, 0)

		if err != nil {
			log.Fatal(err)
		}

		if xTop < 0 {
			err = y.SetValue(0, 0, x.Normal("2"))
		} else {
			err = y.SetValue(0, 0, -x.Normal("2"))
		}

		if err != nil {
			log.Fatal(err)
		}

		v, _ := x.Subtract(y)
		fmt.Println("v Normal", v.Normal("2"))
		u := v.ScalarMultiply(1 / v.Normal("2"))

		r1, _ := R.Minor(k, m-1, k, n-1)
		r1.Print()
		uDot, _ := u.Multiply(u.Transpose())
		houseHolder, _ := uDot.ScalarMultiply(2).Multiply(r1)
		temp, _ := r1.Subtract(houseHolder)
		R, _ = R.MinorUpdate(k, m-1, k, n-1, temp)

		q1, _ := Q.Minor(k, m-1, k, n-1)
		q1.Print()
		houseHolder, _ = uDot.ScalarMultiply(2).Multiply(q1)
		temp, _ = q1.Subtract(houseHolder)
		Q, _ = Q.MinorUpdate(k, m-1, k, n-1, temp)
	}

	Q = Q.Transpose()
	return
}
