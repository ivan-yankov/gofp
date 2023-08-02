package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestArrayOfGoSlice_Empty(t *testing.T) {
	a := fp.ArrayOfGoSlice[int]([]int{})
	assert.True(t, a.HeadOption().NonDefined())
	assert.True(t, a.Tail().IsEmpty())
	assert.True(t, a.IsEmpty())
}

func TestArrayOfGoSlice_One(t *testing.T) {
	a := fp.ArrayOfGoSlice([]int{1})
	assert.Equal(t, 1, a.HeadOption().GetOrElse(0))
	assert.True(t, a.Tail().IsEmpty())
	assert.False(t, a.IsEmpty())
}

func TestArrayOfGoSlice_NonEmpty(t *testing.T) {
	a := fp.ArrayOfGoSlice([]int{1, 2, 3})
	assert.Equal(t, 1, a.HeadOption().GetOrElse(0))
	assert.Equal(t, 2, a.Tail().HeadOption().GetOrElse(0))
	assert.Equal(t, 3, a.Tail().Tail().HeadOption().GetOrElse(0))
	assert.True(t, a.Tail().Tail().Tail().IsEmpty())
	assert.False(t, a.IsEmpty())
}

func TestArrayOf_Empty(t *testing.T) {
	a := fp.ArrayOf[int]()
	assert.True(t, a.HeadOption().NonDefined())
	assert.True(t, a.Tail().IsEmpty())
	assert.True(t, a.IsEmpty())
}

func TestArrayOf_One(t *testing.T) {
	a := fp.ArrayOf(1)
	assert.Equal(t, 1, a.HeadOption().GetOrElse(0))
	assert.True(t, a.Tail().IsEmpty())
	assert.False(t, a.IsEmpty())
}

func TestArrayOf_NonEmpty(t *testing.T) {
	a := fp.ArrayOf(1, 2, 3)
	assert.Equal(t, 1, a.HeadOption().GetOrElse(0))
	assert.Equal(t, 2, a.Tail().HeadOption().GetOrElse(0))
	assert.Equal(t, 3, a.Tail().Tail().HeadOption().GetOrElse(0))
	assert.True(t, a.Tail().Tail().Tail().IsEmpty())
	assert.False(t, a.IsEmpty())
}

func TestArrayRangeStep(t *testing.T) {
	assert.True(t, fp.ArrayRangeStep(0, -2, 1).IsEmpty())
	assert.True(t, fp.ArrayRangeStep(0, 2, -1).IsEmpty())
	assert.True(t, fp.ArrayRangeStep(0, 0, 1).IsEmpty())

	assert.True(t, fp.ArrayRangeStep(1, 1, 1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayRangeStep(0, 5, 1).Equals(fp.ArrayOf(0, 1, 2, 3, 4)))
	assert.True(t, fp.ArrayRangeStep(1, 5, 2).Equals(fp.ArrayOf(1, 3, 5, 7, 9)))
	assert.True(t, fp.ArrayRangeStep(5, 3, 10).Equals(fp.ArrayOf(5, 15, 25)))
	assert.True(t, fp.ArrayRangeStep(-3, 4, 5).Equals(fp.ArrayOf(-3, 2, 7, 12)))
}

func TestArrayRange(t *testing.T) {
	assert.True(t, fp.ArrayRange(0, -2).IsEmpty())
	assert.True(t, fp.ArrayRange(0, 0).IsEmpty())

	assert.True(t, fp.ArrayRange(1, 1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayRange(0, 5).Equals(fp.ArrayOf(0, 1, 2, 3, 4)))
	assert.True(t, fp.ArrayRange(1, 5).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayRange(6, 3).Equals(fp.ArrayOf(6, 7, 8)))
}

func TestArrayTabulate(t *testing.T) {
	assert.True(
		t,
		fp.ArrayTabulate(5, func(i int) int { return (i + 1) * 2 }).
			Equals(fp.ArrayRangeStep(2, 5, 2)),
	)
}

func TestArrayFill(t *testing.T) {
	assert.True(t, fp.ArrayFill(5, 1).Equals(fp.ArrayOf(1, 1, 1, 1, 1)))
}
