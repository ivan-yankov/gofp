package fp_test

import (
	"fmt"
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestSeqZip(t *testing.T) {
	exp := fp.ListOf(
		fp.PairOf("zero", 0),
		fp.PairOf("one", 1),
		fp.PairOf("two", 2),
	)

	assert.True(t, fp.SeqZip(fp.ListOf[int](), fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.SeqZip(fp.ListOf(1, 2), fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.SeqZip(fp.ListOf[int](), fp.ListOf("1", "2")).IsEmpty())

	assert.True(t, fp.SeqZip(fp.ListOf("zero", "one", "two"), fp.ListOf(0, 1, 2)).Equals(exp))
	assert.True(t, fp.SeqZip(fp.ListOf("zero", "one", "two"), fp.ListOf(0, 1, 2, 3, 4)).Equals(exp))
	assert.True(t, fp.SeqZip(fp.ListOf("zero", "one", "two", "next", "one more"), fp.ListOf(0, 1, 2)).Equals(exp))
}

func TestSeqZipWithIndex(t *testing.T) {
	exp := fp.ListOf(
		fp.PairOf("zero", 0),
		fp.PairOf("one", 1),
		fp.PairOf("two", 2),
	)

	assert.True(t, fp.SeqZipWithIndex(fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.SeqZipWithIndex(fp.ListOf("zero", "one", "two")).Equals(exp))
}

func TestSeqFoldLeft(t *testing.T) {
	f := func(i int, acc string) string {
		return acc + fmt.Sprint(i)
	}
	assert.Equal(t, "", fp.SeqFoldLeft(fp.ListOf[int](), f, ""))
	assert.Equal(t, "12345", fp.SeqFoldLeft(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestSeqFoldRightLeft(t *testing.T) {
	f := func(i int, acc string) string {
		return acc + fmt.Sprint(i)
	}
	assert.Equal(t, "", fp.SeqFoldRight(fp.ListOf[int](), f, ""))
	assert.Equal(t, "54321", fp.SeqFoldRight(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestSeqFoldCount(t *testing.T) {
	f := func(i int, x int, acc string) string {
		return acc + "[" + fmt.Sprint(i) + "," + fmt.Sprint(x) + "]"
	}
	assert.Equal(t, "", fp.SeqFoldCount(fp.ListOf[int](), f, ""))
	assert.Equal(t, "[0,1][1,2][2,3][3,4][4,5]", fp.SeqFoldCount(fp.ListOf(1, 2, 3, 4, 5), f, ""))
}

func TestSeqMap(t *testing.T) {
	f := func(x int) string {
		return "r" + fmt.Sprint(x)
	}

	assert.True(t, fp.SeqMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.SeqMap(fp.ListOf(1), f).Equals(fp.ListOf("r1")))
	assert.True(t, fp.SeqMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("r1", "r2", "r3")))
}

func TestSeqReverseMap(t *testing.T) {
	f := func(x int) string {
		return "r" + fmt.Sprint(x)
	}

	assert.True(t, fp.SeqReverseMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.SeqReverseMap(fp.ListOf(1), f).Equals(fp.ListOf("r1")))
	assert.True(t, fp.SeqReverseMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("r3", "r2", "r1")))
}

func TestSeqFlatMap(t *testing.T) {
	f := func(x int) fp.Seq[string] {
		return fp.ListOf("result", "r"+fmt.Sprint(x))
	}

	assert.True(t, fp.SeqFlatMap(fp.ListOf[int](), f).IsEmpty())
	assert.True(t, fp.SeqFlatMap(fp.ListOf(1), f).Equals(fp.ListOf("result", "r1")))
	assert.True(t, fp.SeqFlatMap(fp.ListOf(1, 2, 3), f).Equals(fp.ListOf("result", "r1", "result", "r2", "result", "r3")))
}

func TestSeqSliding(t *testing.T) {
	assert.True(t, fp.SeqSliding(fp.ListOf[int](), 0, 0).IsEmpty())
	assert.True(t, fp.SeqSliding(fp.ListOf[int](), 0, 1).IsEmpty())
	assert.True(t, fp.SeqSliding(fp.ListOf[int](), 1, 0).IsEmpty())
	assert.True(t, fp.SeqSliding(fp.ListOf(1), -1, 1).IsEmpty())
	assert.True(t, fp.SeqSliding(fp.ListOf(1), 1, -1).IsEmpty())

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3), 1, 1).Equals(fp.ListOf(fp.ListOf(1), fp.ListOf(2), fp.ListOf(3))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 2, 1).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(2, 3), fp.ListOf(3, 4), fp.ListOf(4, 5), fp.ListOf(5, 6))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6, 7), 2, 1).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(2, 3), fp.ListOf(3, 4), fp.ListOf(4, 5), fp.ListOf(5, 6), fp.ListOf(6, 7))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5), 2, 2).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(3, 4), fp.ListOf(5))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 2, 2).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(3, 4), fp.ListOf(5, 6))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6, 7), 2, 2).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(3, 4), fp.ListOf(5, 6), fp.ListOf(7))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 2, 3).Equals(fp.ListOf(fp.ListOf(1, 2), fp.ListOf(4, 5))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 3, 1).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(2, 3, 4), fp.ListOf(3, 4, 5), fp.ListOf(4, 5, 6))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6, 7), 3, 1).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(2, 3, 4), fp.ListOf(3, 4, 5), fp.ListOf(4, 5, 6), fp.ListOf(5, 6, 7))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 3, 2).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(3, 4, 5), fp.ListOf(5, 6))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6, 7), 3, 2).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(3, 4, 5), fp.ListOf(5, 6, 7))))

	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6), 3, 3).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(4, 5, 6))))
	assert.True(t, fp.SeqSliding(fp.ListOf(1, 2, 3, 4, 5, 6, 7), 3, 3).Equals(fp.ListOf(fp.ListOf(1, 2, 3), fp.ListOf(4, 5, 6), fp.ListOf(7))))
}

func TestSeqStartsWith(t *testing.T) {
	assert.False(t, fp.SeqStartsWith(fp.ListOf[int](), fp.ListOf[int]()))
	assert.False(t, fp.SeqStartsWith(fp.ListOf[int](), fp.ListOf(1)))
	assert.False(t, fp.SeqStartsWith(fp.ListOf(1), fp.ListOf[int]()))
	assert.True(t, fp.SeqStartsWith(fp.ListOf(1), fp.ListOf(1)))
	assert.True(t, fp.SeqStartsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3)))
	assert.True(t, fp.SeqStartsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3)))
	assert.False(t, fp.SeqStartsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5, 6)))
	assert.False(t, fp.SeqStartsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)))
}

func TestSeqEndsWith(t *testing.T) {
	assert.False(t, fp.SeqEndsWith(fp.ListOf[int](), fp.ListOf[int]()))
	assert.False(t, fp.SeqEndsWith(fp.ListOf[int](), fp.ListOf(1)))
	assert.False(t, fp.SeqEndsWith(fp.ListOf(1), fp.ListOf[int]()))
	assert.True(t, fp.SeqEndsWith(fp.ListOf(1), fp.ListOf(1)))
	assert.True(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3)))
	assert.False(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5, 6)))
	assert.False(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(3, 4, 5)))
	assert.False(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3, 4, 5, 6)))
	assert.False(t, fp.SeqEndsWith(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(6, 1, 2, 3, 4, 5)))
}

func TestListFindSlice(t *testing.T) {
	assert.True(t, fp.SeqFindSlice(fp.ListOf[int](), fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.SeqFindSlice(fp.ListOf(1), fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.SeqFindSlice(fp.ListOf[int](), fp.ListOf(1)).NonDefined())
	assert.True(t, fp.SeqFindSlice(fp.ListOf(1, 2, 3), fp.ListOf(1, 2, 3, 4, 5)).NonDefined())

	assert.Equal(t, 0, fp.SeqFindSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(1, 2, 3, 4, 5)).GetOrElse(-1))
	assert.Equal(t, 3, fp.SeqFindSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(4, 5)).GetOrElse(-1))
	assert.Equal(t, 2, fp.SeqFindSlice(fp.ListOf(1, 2, 3, 4, 5), fp.ListOf(3, 4, 5)).GetOrElse(-1))
}
