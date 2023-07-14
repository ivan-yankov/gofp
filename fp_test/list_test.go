package fp_test

import (
	"fmt"
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestListZip(t *testing.T) {
	exp := fp.ListOf(
		fp.PairOf("zero", 0),
		fp.PairOf("one", 1),
		fp.PairOf("two", 2),
	)

	assert.True(t, fp.ListZip(fp.ListOf[int](), fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.ListZip(fp.ListOf(1, 2), fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.ListZip(fp.ListOf[int](), fp.ListOf("1", "2")).IsEmpty())

	assert.True(t, fp.ListZip(fp.ListOf("zero", "one", "two"), fp.ListOf(0, 1, 2)).Equals(exp))
	assert.True(t, fp.ListZip(fp.ListOf("zero", "one", "two"), fp.ListOf(0, 1, 2, 3, 4)).Equals(exp))
	assert.True(t, fp.ListZip(fp.ListOf("zero", "one", "two", "next", "one more"), fp.ListOf(0, 1, 2)).Equals(exp))
}

func TestListFoldLeft(t *testing.T) {
	f := func(i int, acc string) string {
		return acc + fmt.Sprint(i)
	}
	assert.Equal(t, "", fp.ListFoldLeft(fp.ListOf[int](), f, ""))
	assert.Equal(t, "12345", fp.ListFoldLeft(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestListFoldRightLeft(t *testing.T) {
	f := func(i int, acc string) string {
		return acc + fmt.Sprint(i)
	}
	assert.Equal(t, "", fp.ListFoldRight(fp.ListOf[int](), f, ""))
	assert.Equal(t, "54321", fp.ListFoldRight(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestListFoldCount(t *testing.T) {
	f := func(i int, x int, acc string) string {
		return acc + "[" + fmt.Sprint(i) + "," + fmt.Sprint(x) + "]"
	}
	assert.Equal(t, "", fp.ListFoldCount(fp.ListOf[int](), f, ""))
	assert.Equal(t, "[0,1][1,2][2,3][3,4][4,5]", fp.ListFoldCount(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestListStartsWith(t *testing.T) {
	assert.False(t, fp.ListStartsWith(fp.ListOf[int](), fp.ListOf[int]()))
	assert.False(t, fp.ListStartsWith(fp.ListOf[int](), fp.ListOf(1)))
	assert.False(t, fp.ListStartsWith(fp.ListOf(1), fp.ListOf[int]()))
	assert.True(t, fp.ListStartsWith(fp.ListOf(1), fp.ListOf(1)))
	assert.True(t, fp.ListStartsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListStartsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3)))
	assert.False(t, fp.ListStartsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5, 6)))
	assert.False(t, fp.ListStartsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)))
}

func TestListEndsWith(t *testing.T) {
	assert.False(t, fp.ListEndsWith(fp.ListOf[int](), fp.ListOf[int]()))
	assert.False(t, fp.ListEndsWith(fp.ListOf[int](), fp.ListOf(1)))
	assert.False(t, fp.ListEndsWith(fp.ListOf(1), fp.ListOf[int]()))
	assert.True(t, fp.ListEndsWith(fp.ListOf(1), fp.ListOf(1)))
	assert.True(t, fp.ListEndsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3)))
	assert.False(t, fp.ListEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5, 6)))
	assert.False(t, fp.ListEndsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(3, 4, 5)))
	assert.False(t, fp.ListEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3, 4, 5, 6)))
	assert.False(t, fp.ListEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(6, 1, 2, 3, 4, 5)))
}

func TestListContainsSlice(t *testing.T) {
	assert.True(t, fp.ListContainsSlice(fp.ListOf[int](), fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.ListContainsSlice(fp.ListOf(1), fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.ListContainsSlice(fp.ListOf[int](), fp.ListOf(1)).NonDefined())
	assert.True(t, fp.ListContainsSlice(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)).NonDefined())

	assert.Equal(t, 0, fp.ListContainsSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3, 4, 5)).GetOrElse(-1))
	assert.Equal(t, 3, fp.ListContainsSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5)).GetOrElse(-1))
	assert.Equal(t, 2, fp.ListContainsSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(3, 4, 5)).GetOrElse(-1))
}
