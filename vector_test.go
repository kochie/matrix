package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector(t *testing.T) {
	assert := assert.New(t)

	v1, err := Vector(1, 0, 0)

	assert.Nil(err)
	assert.Equal(v1.Rows, 3)
	assert.Equal(v1.Columns, 1)
	assert.Equal(v1.Capacity, 3)
	assert.Equal(v1.Elements, []float64{1, 0, 0})

	v2, err := Vector(3, 4, 5, 6)
	v2 = v2.Transpose()

	assert.Nil(err)
	assert.Equal(v2.Rows, 1)
	assert.Equal(v2.Columns, 4)
	assert.Equal(v2.Capacity, 4)
	assert.Equal(v2.Elements, []float64{3, 4, 5, 6})
}

func TestIsVector(t *testing.T) {
	assert := assert.New(t)

	v1, err := Vector(1, 0, 0)
	assert.Nil(err)
	assert.True(v1.IsVector())

	v2, err := Matrix(1, 1, []float64{1})
	assert.Nil(err)
	assert.True(v2.IsVector())

	v3, err := Matrix(2, 2, []float64{1, 2, 3, 4})
	assert.Nil(err)
	assert.False(v3.IsVector())

	v4, err := Matrix(3, 1, []float64{1, 2, 3})
	assert.Nil(err)
	assert.True(v4.Transpose().IsVector())
}

func TestUnit(t *testing.T) {
	assert := assert.New(t)

	v3, err := Vector(2, 2, 0)
	assert.Nil(err)
	v3_u, err := v3.Unit()
	v4, err := Vector(0.7071067811865475, 0.7071067811865475, 0)
	assert.Nil(err)
	assert.Equal(v3_u.Rows, v4.Rows)
	assert.Equal(v3_u.Columns, v4.Columns)
	assert.Equal(v3_u.Capacity, v4.Capacity)
	assert.Equal(v3_u.Elements, v4.Elements)
}
