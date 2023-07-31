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

func TestListMap(t *testing.T) {
	f := func(x int) string {
		return "r" + fmt.Sprint(x)
	}

	assert.True(t, fp.ListMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.ListMap(fp.ListOf(1), f).Equals(fp.ListOf("r1")))
	assert.True(t, fp.ListMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("r1", "r2", "r3")))
}

func TestListReverseMap(t *testing.T) {
	f := func(x int) string {
		return "r" + fmt.Sprint(x)
	}

	assert.True(t, fp.ListReverseMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.ListReverseMap(fp.ListOf(1), f).Equals(fp.ListOf("r1")))
	assert.True(t, fp.ListReverseMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("r3", "r2", "r1")))
}

func TestListFlatMap(t *testing.T) {
	f := func(x int) fp.Seq[string] {
		return fp.ListOf("result", "r"+fmt.Sprint(x))
	}

	assert.True(t, fp.ListFlatMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.ListFlatMap(fp.ListOf(1), f).Equals(fp.ListOf("result", "r1")))
	assert.True(t, fp.ListFlatMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("result", "r1", "result", "r2", "result", "r3")))
}
