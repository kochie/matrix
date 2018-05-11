package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUpperTriangular(t *testing.T) {
	assert := assert.New(t)

	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 0, 6, 7, 8, 0, 0, 11, 12, 0, 0, 0, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	assert.True(a.IsUpperTriangular())

	b, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 0, 5, 11, 12, 0, 0, 0, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	assert.False(b.IsUpperTriangular())
}

func BenchmarkIsUpperTriangular(b *testing.B) {
	a, _ := Matrix(4, 4, []float64{1, 2, 3, 4, 0, 6, 7, 8, 0, 0, 11, 12, 0, 0, 0, 16})
	for n := 0; n < b.N; n++ {
		_ = a.IsUpperTriangular()
	}
}

func TestIsLowerTriangular(t *testing.T) {
	assert := assert.New(t)

	a, err := Matrix(4, 4, []float64{1, 0, 0, 0, 5, 6, 0, 0, 9, 10, 11, 0, 13, 14, 15, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	assert.True(a.IsLowerTriangular())

	b, err := Matrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 0, 5, 11, 12, 0, 0, 0, 16})
	if err != nil {
		t.Errorf("Caught Error %s", err.Error())
	}

	assert.False(b.IsLowerTriangular())
}

func BenchmarkIsLowerTriangular(b *testing.B) {
	a, _ := Matrix(4, 4, []float64{1, 0, 0, 0, 5, 6, 0, 0, 9, 10, 11, 0, 13, 14, 15, 16})
	for n := 0; n < b.N; n++ {
		_ = a.IsLowerTriangular()
	}
}

func TestTriangleInverse(t *testing.T) {
	assert := assert.New(t)

	a, err := Matrix(4, 4, []float64{1, 2, 3, 4, 0, 5, 6, 7, 0, 0, 8, 9, 0, 0, 0, 10})
	assert.Nil(err)

	b, err := a.TriangleInverse()
	assert.Nil(err)

	c, err := Matrix(4, 4, []float64{1, -0.4, -0.07499999999999996, -0.05249999999999999, 0, 0.2, -0.15000000000000002, -0.0050000000000000044, 0, 0, 0.125, -0.1125, 0, 0, 0, 0.1})
	assert.Nil(err)

	assert.Equal(c.Rows, b.Rows)
	assert.Equal(c.Columns, b.Columns)
	assert.Equal(c.Elements, b.Elements)

	d, err := Matrix(4, 4, []float64{1, 0, 0, 0, 2, 3, 0, 0, 4, 5, 6, 0, 7, 8, 9, 10})
	assert.Nil(err)

	e, err := d.TriangleInverse()
	assert.Nil(err)

	f, err := Matrix(4, 4, []float64{1, 0, 0, 0, -0.6666666666666666, 0.3333333333333333, 0, 0, -0.11111111111111116, -0.27777777777777773, 0.16666666666666666, 0, -0.06666666666666665, -0.016666666666666698, -0.15000000000000002, 0.1})
	assert.Nil(err)

	assert.Equal(e.Rows, f.Rows)
	assert.Equal(e.Columns, f.Columns)
	assert.Equal(e.Elements, f.Elements)
}

func Benchmark_TriangleInverse_Upper(b *testing.B) {
	a, _ := Matrix(4, 4, []float64{1, 2, 3, 4, 0, 5, 6, 7, 0, 0, 8, 9, 0, 0, 0, 10})
	for n := 0; n < b.N; n++ {
		_, _ = a.TriangleInverse()
	}
}

func Benchmark_TriangleInverse_Lower(b *testing.B) {
	a, _ := Matrix(4, 4, []float64{1, 0, 0, 0, 2, 3, 0, 0, 4, 5, 6, 0, 7, 8, 9, 10})
	for n := 0; n < b.N; n++ {
		_, _ = a.TriangleInverse()
	}
}
