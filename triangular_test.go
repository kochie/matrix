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
	a.Print()
	b.Print()
}
