package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestListOf_Empty(t *testing.T) {
	lst := fp.ListOf[int]()
	assert.True(t, lst.HeadOption().NonDefined())
	assert.True(t, lst.Tail().IsEmpty())
	assert.True(t, lst.IsEmpty())
}

func TestListOf_One(t *testing.T) {
	lst := fp.ListOf(1)
	assert.Equal(t, 1, lst.HeadOption().GetOrElse(0))
	assert.True(t, lst.Tail().IsEmpty())
	assert.False(t, lst.IsEmpty())
}

func TestListOf_NonEmpty(t *testing.T) {
	lst := fp.ListOf(1, 2, 3)
	assert.Equal(t, 1, lst.HeadOption().GetOrElse(0))
	assert.Equal(t, 2, lst.Tail().HeadOption().GetOrElse(0))
	assert.Equal(t, 3, lst.Tail().Tail().HeadOption().GetOrElse(0))
	assert.True(t, lst.Tail().Tail().Tail().IsEmpty())
	assert.False(t, lst.IsEmpty())
}

func TestListRangeStep(t *testing.T) {
	assert.True(t, fp.ListRangeStep(0, -2, 1).IsEmpty())
	assert.True(t, fp.ListRangeStep(0, 2, -1).IsEmpty())
	assert.True(t, fp.ListRangeStep(0, 0, 1).IsEmpty())

	assert.True(t, fp.ListRangeStep(1, 1, 1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListRangeStep(0, 5, 1).Equals(fp.ListOf(0, 1, 2, 3, 4)))
	assert.True(t, fp.ListRangeStep(1, 5, 2).Equals(fp.ListOf(1, 3, 5, 7, 9)))
	assert.True(t, fp.ListRangeStep(5, 3, 10).Equals(fp.ListOf(5, 15, 25)))
	assert.True(t, fp.ListRangeStep(-3, 4, 5).Equals(fp.ListOf(-3, 2, 7, 12)))
}

func TestListRange(t *testing.T) {
	assert.True(t, fp.ListRange(0, -2).IsEmpty())
	assert.True(t, fp.ListRange(0, 0).IsEmpty())

	assert.True(t, fp.ListRange(1, 1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListRange(0, 5).Equals(fp.ListOf(0, 1, 2, 3, 4)))
	assert.True(t, fp.ListRange(1, 5).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListRange(6, 3).Equals(fp.ListOf(6, 7, 8)))
}

func TestListTabulate(t *testing.T) {
	assert.True(
		t,
		fp.ListTabulate(5, func(i int) int { return (i + 1) * 2 }).
			Equals(fp.ListRangeStep(2, 5, 2)),
	)
}

func TestListFill(t *testing.T) {
	assert.True(t, fp.ListFill(5, 1).Equals(fp.ListOf(1, 1, 1, 1, 1)))
}
