package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrix(t *testing.T) {
	a, err := Matrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	if a.Rows != 2 {
		t.Errorf("Row incorrect, should be %d, received %d", 2, a.Rows)
	}
	if a.Columns != 3 {
		t.Errorf("Column incorrect, should be %d, received %d", 3, a.Columns)
	}
	if a.Capacity != 6 {
		t.Errorf("Incorrect capacity, should be %d, received %d", 6, a.Capacity)
	}

	a, err = Matrix(0, 4, []float64{1, 2, 3, 4})
	if err == nil {
		t.Errorf("Did not error on incorrect matrix values")
	}

	a, err = Matrix(4, -1, []float64{1, 2, 3, 4})
	if err == nil {
		t.Errorf("Did not error on incorrect matrix values")
	}

	a, err = Matrix(4, 1, []float64{1, 2, 3, 4, 5})
	if err == nil {
		t.Errorf("Did not error on incorrect matrix values")
	}
}

func TestEye(t *testing.T) {
	a, err := Eye(3, 3)
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	for i := 0; i < a.Columns; i++ {
		for j := 0; j < a.Rows; j++ {
			if i == j && a.Elements[i*a.Columns+j] != 1 {
				a.Print()
				t.Errorf("Identity matrix incorrect in the diagonal!")
			} else if i != j && a.Elements[i*a.Columns+j] != 0 {
				a.Print()
				t.Errorf("Identity matrix incorrect in the extra space!")
			}
		}
	}
}

func TestGetValue(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	a, err := Matrix(4, 4, list)

	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a11, err := a.GetValue(i, j)
			if err != nil {
				t.Errorf("Caught Error %s", err.Error())
			}
			if a11 != list[i*4+j] {
				t.Errorf("Wrong Number %f, should be %f at position %d, %d", a11, list[i*4+j], i, j)
			}
		}
	}

	_, err = a.GetValue(5, 5)
	if err == nil {
		t.Errorf("Did not throw error for out of bounds matrix grab.")
	}
}

func TestMinor(t *testing.T) {
	assert := assert.New(t)

	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	b, err := a.Minor(1, 3, 1, 3)
	assert.Equal(b.Rows, 3)
	assert.Equal(b.Columns, 3)
	assert.Equal(b.Capacity, 9)
	assert.Equal(b.Elements, []float64{6, 7, 8, 10, 11, 12, 14, 15, 16})

	c, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	d, err := c.Minor(0, 3, 0, 0)
	assert.Equal(d.Rows, 4)
	assert.Equal(d.Columns, 1)
	assert.Equal(d.Capacity, 4)
	assert.Equal(d.Elements, []float64{1, 5, 9, 13})

}

func TestTranspose(t *testing.T) {
	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	b, err := Matrix(4, 4, []float64{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	if !a.Transpose().IsEqual(b) {
		t.Errorf("Transpose is not equal to the transpose!")
	}

	c, err := Matrix(4, 1, []float64{1, 2, 3, 4})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	d, err := Matrix(1, 4, []float64{1, 2, 3, 4})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	if !c.Transpose().IsEqual(d) {
		t.Errorf("Transpose is not equal to the transpose!")
	}

	e, err := Matrix(4, 2, []float64{1, 5, 2, 6, 3, 7, 4, 8})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	f, err := Matrix(2, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	if !e.Transpose().IsEqual(f) {
		t.Errorf("Transpose is not equal to the transpose!")
	}
}

func TestMultiply(t *testing.T) {
	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	b, err := a.Multiply(a)
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	c, err := Matrix(4, 4, []float64{90, 100, 110, 120, 202, 228, 254, 280, 314, 356, 398, 440, 426, 484, 542, 600})
	if !b.IsEqual(c) {
		t.Errorf("Incorrect Matrix Multiplication")
	}

	d, err := Matrix(8, 1, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	e, err := Matrix(1, 8, []float64{1, 2, 3, 4, 5, 6, 7, 8})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	f, err := d.Multiply(e)
	g, err := Matrix(8, 8, []float64{1, 2, 3, 4, 5, 6, 7, 8, 2, 4, 6, 8, 10, 12, 14, 16, 3, 6, 9, 12, 15, 18, 21, 24, 4, 8, 12, 16, 20, 24, 28, 32, 5, 10, 15, 20, 25, 30, 35, 40, 6, 12, 18, 24, 30, 36, 42, 48, 7, 14, 21, 28, 35, 42, 49, 56, 8, 16, 24, 32, 40, 48, 56, 64})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	if !g.IsEqual(f) {
		g.Print()
		f.Print()
		t.Errorf("Incorrect Matrix")
	}
}

func TestQR(t *testing.T) {
	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}
	Q, R := a.QR()

	A, _ := Q.Multiply(R)
	A.Print()
}

func BenchmarkQR(b *testing.B) {
	a, _ := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	for n := 0; n < b.N; n++ {
		_, _ = a.QR()
	}
}
