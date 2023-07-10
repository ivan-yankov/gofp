package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestListOf_Empty(t *testing.T) {
	lst := fp.ListOf[int]()
	assert.True(t, lst.HeadOption().NonDefined())
}

func TestListOf_One(t *testing.T) {
	lst := fp.ListOf(1)
	assert.Equal(t, 1, lst.HeadOption().GetOrElse(0))
}

func TestListOf_NonEmpty(t *testing.T) {
	lst := fp.ListOf(1, 2, 3)
	assert.Equal(t, 1, lst.HeadOption().GetOrElse(0))
	assert.Equal(t, 2, lst.Tail().HeadOption().GetOrElse(0))
	assert.Equal(t, 3, lst.Tail().Tail().HeadOption().GetOrElse(0))
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

func TestListAdd_Empty(t *testing.T) {
	lst := fp.ListOf[int]().Add(1)
	assert.True(t, lst.NonEmpty())
	assert.True(t, lst.Tail().IsEmpty())
	assert.Equal(t, 1, lst.HeadOption().GetOrElse(0))
}

func TestListAdd_NonEmpty2(t *testing.T) {
	lst := fp.ListOf[int](1).Add(2)
	assert.True(t, lst.NonEmpty())
	assert.True(t, lst.Tail().NonEmpty())
	assert.Equal(t, 2, lst.HeadOption().GetOrElse(0))
	assert.Equal(t, 1, lst.Tail().HeadOption().GetOrElse(0))
}

func TestListAdd_NonEmpty3(t *testing.T) {
	lst := fp.ListOf[int](1, 2).Add(3)
	assert.True(t, lst.NonEmpty())
	assert.True(t, lst.Tail().NonEmpty())
	assert.Equal(t, 3, lst.HeadOption().GetOrElse(0))
	assert.Equal(t, 1, lst.Tail().HeadOption().GetOrElse(0))
	assert.Equal(t, 2, lst.Tail().Tail().HeadOption().GetOrElse(0))
}
