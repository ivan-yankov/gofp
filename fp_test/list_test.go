package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestListTabulate(t *testing.T) {
	// ListOf uses ListTabulate
	// add expected elements in a different way to avoid testing ListTabulate against itself
	exp := fp.ListOf[int]().Add(2).Add(4).Add(6).Add(8).Add(10).Reverse()
	assert.True(
		t,
		fp.ListTabulate(5, func(i int) int { return (i + 1) * 2 }).
			Equals(exp),
	)
}

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

func TestListFill(t *testing.T) {
	assert.True(t, fp.ListFill(5, 1).Equals(fp.ListOf(1, 1, 1, 1, 1)))
}

func TestListRange(t *testing.T) {
	assert.True(t, fp.ListRange(1, 5).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListRange(6, 3).Equals(fp.ListOf(6, 7, 8)))
}

func TestListAdd(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Add(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf[int](1).Add(2).Equals(fp.ListOf(2, 1)))
	assert.True(t, fp.ListOf[int](1, 2).Add(3).Equals(fp.ListOf(3, 1, 2)))
	assert.True(t, fp.ListOf[int](1, 2).Add(3).Add(4).Equals(fp.ListOf(4, 3, 1, 2)))
}

func TestListIsEmpty(t *testing.T) {
	assert.True(t, fp.ListOf[int]().IsEmpty())
	assert.False(t, fp.ListOf[int](1, 2, 3).IsEmpty())
}

func TestListNonEmpty(t *testing.T) {
	assert.False(t, fp.ListOf[int]().NonEmpty())
	assert.True(t, fp.ListOf[int](1, 2, 3).NonEmpty())
}

func TestListHeadOption(t *testing.T) {
	assert.True(t, fp.ListOf[int]().HeadOption().NonDefined())

	assert.True(t, fp.ListOf[int](1).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf[int](1).HeadOption().GetOrElse(0))

	assert.True(t, fp.ListOf[int](1, 2, 3).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf[int](1, 2, 3).HeadOption().GetOrElse(0))
}

func TestListLastOption(t *testing.T) {
	assert.True(t, fp.ListOf[int]().LastOption().NonDefined())

	assert.True(t, fp.ListOf[int](1).LastOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf[int](1).LastOption().GetOrElse(0))

	assert.True(t, fp.ListOf[int](1, 2, 3).LastOption().IsDefined())
	assert.Equal(t, 3, fp.ListOf[int](1, 2, 3).LastOption().GetOrElse(0))
}

func TestListTail(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Tail().IsEmpty())

	assert.True(t, fp.ListOf[int](1).Tail().IsEmpty())

	assert.True(t, fp.ListOf[int](1, 2, 3).Tail().NonEmpty())
	assert.Equal(t, 2, fp.ListOf[int](1, 2, 3).Tail().HeadOption().GetOrElse(0))
}

func TestListEquals(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf[int](1).Equals(fp.ListOf[int](1)))
	assert.True(t, fp.ListOf[int](1, 2, 3).Equals(fp.ListOf[int](1, 2, 3)))

	assert.False(t, fp.ListOf[int]().Equals(fp.ListOf[int](1)))
	assert.False(t, fp.ListOf[int](1, 2).Equals(fp.ListOf[int](2, 1)))
	assert.False(t, fp.ListOf[int](1, 2).Equals(fp.ListOf[int](1, 3)))
	assert.False(t, fp.ListOf[int](1, 2).Equals(fp.ListOf[int](1, 2, 3)))
}

func TestListReverse(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Reverse().Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf[int](1).Reverse().Equals(fp.ListOf[int](1)))
	assert.True(t, fp.ListOf[int](1, 2, 3).Reverse().Equals(fp.ListOf[int](3, 2, 1)))
	assert.True(t, fp.ListRange(1, 5).Reverse().Reverse().Equals(fp.ListRange(1, 5)))
}

func TestListAppend(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Append(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf[int](1).Append(2).Equals(fp.ListOf(1, 2)))
	assert.True(t, fp.ListOf[int](1, 2).Append(3).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf[int](1, 2).Append(3).Append(4).Equals(fp.ListOf(1, 2, 3, 4)))
}

func TestListConcat(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Concat(fp.ListOf[int]()).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf[int]().Concat(fp.ListOf[int](1)).Equals(fp.ListOf[int](1)))
	assert.True(t, fp.ListOf[int](1).Concat(fp.ListOf[int](2, 3)).Equals(fp.ListOf[int](1, 2, 3)))
	assert.True(t, fp.ListOf[int](1, 2, 3).Concat(fp.ListOf[int](4, 5)).Equals(fp.ListOf[int](1, 2, 3, 4, 5)))
}
