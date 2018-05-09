package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector(t *testing.T) {
	assert := assert.New(t)

	v1, err := Vector(1, 0, 0)

	assert.Nil(err)
	assert.Equal(v1.rows, 3)
	assert.Equal(v1.columns, 1)
	assert.Equal(v1.capacity, 3)
	assert.Equal(v1.elements, []float64{1, 0, 0})

	v2, err := Vector(3, 4, 5, 6)
	v2 = v2.Transpose()

	assert.Nil(err)
	assert.Equal(v2.rows, 1)
	assert.Equal(v2.columns, 4)
	assert.Equal(v2.capacity, 4)
	assert.Equal(v2.elements, []float64{3, 4, 5, 6})
}

func TestIsVector(t *testing.T) {
	assert := assert.New(t)

	v1, err := Vector(1, 0, 0)
	assert.Nil(err)
	assert.True(v1.IsVector())

	v2, err := NewMatrix(1, 1, []float64{1})
	assert.Nil(err)
	assert.True(v2.IsVector())

	v3, err := NewMatrix(2, 2, []float64{1, 2, 3, 4})
	assert.Nil(err)
	assert.False(v3.IsVector())

	v4, err := NewMatrix(3, 1, []float64{1, 2, 3})
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
	assert.Equal(v3_u.rows, v4.rows)
	assert.Equal(v3_u.columns, v4.columns)
	assert.Equal(v3_u.capacity, v4.capacity)
	assert.Equal(v3_u.elements, v4.elements)
}
