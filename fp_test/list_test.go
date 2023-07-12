package fp_test

import (
	"fmt"
	"strings"
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

func TestListAdd(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Add(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1).Add(2).Equals(fp.ListOf(2, 1)))
	assert.True(t, fp.ListOf(1, 2).Add(3).Equals(fp.ListOf(3, 1, 2)))
	assert.True(t, fp.ListOf(1, 2).Add(3).Add(4).Equals(fp.ListOf(4, 3, 1, 2)))
}

func TestListIsEmpty(t *testing.T) {
	assert.True(t, fp.ListOf[int]().IsEmpty())
	assert.False(t, fp.ListOf(1, 2, 3).IsEmpty())
}

func TestListNonEmpty(t *testing.T) {
	assert.False(t, fp.ListOf[int]().NonEmpty())
	assert.True(t, fp.ListOf(1, 2, 3).NonEmpty())
}

func TestListHeadOption(t *testing.T) {
	assert.True(t, fp.ListOf[int]().HeadOption().NonDefined())

	assert.True(t, fp.ListOf(1).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf(1).HeadOption().GetOrElse(0))

	assert.True(t, fp.ListOf(1, 2, 3).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf(1, 2, 3).HeadOption().GetOrElse(0))
}

func TestListLastOption(t *testing.T) {
	assert.True(t, fp.ListOf[int]().LastOption().NonDefined())

	assert.True(t, fp.ListOf(1).LastOption().IsDefined())
	assert.Equal(t, 1, fp.ListOf(1).LastOption().GetOrElse(0))

	assert.True(t, fp.ListOf(1, 2, 3).LastOption().IsDefined())
	assert.Equal(t, 3, fp.ListOf(1, 2, 3).LastOption().GetOrElse(0))
}

func TestListTail(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Tail().IsEmpty())

	assert.True(t, fp.ListOf(1).Tail().IsEmpty())

	assert.True(t, fp.ListOf(1, 2, 3).Tail().NonEmpty())
	assert.Equal(t, 2, fp.ListOf(1, 2, 3).Tail().HeadOption().GetOrElse(0))
}

func TestListEquals(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3).Equals(fp.ListOf(1, 2, 3)))

	assert.False(t, fp.ListOf[int]().Equals(fp.ListOf(1)))
	assert.False(t, fp.ListOf(1, 2).Equals(fp.ListOf(2, 1)))
	assert.False(t, fp.ListOf(1, 2).Equals(fp.ListOf(1, 3)))
	assert.False(t, fp.ListOf(1, 2).Equals(fp.ListOf(1, 2, 3)))
}

func TestListReverse(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Reverse().Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Reverse().Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3).Reverse().Equals(fp.ListOf(3, 2, 1)))
	assert.True(t, fp.ListRange(1, 5).Reverse().Reverse().Equals(fp.ListRange(1, 5)))
}

func TestListAppend(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Append(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1).Append(2).Equals(fp.ListOf(1, 2)))
	assert.True(t, fp.ListOf(1, 2).Append(3).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2).Append(3).Append(4).Equals(fp.ListOf(1, 2, 3, 4)))
}

func TestListConcat(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Concat(fp.ListOf[int]()).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf[int]().Concat(fp.ListOf(1)).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1).Concat(fp.ListOf(2, 3)).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3).Concat(fp.ListOf(4, 5)).Equals(fp.ListOf(1, 2, 3, 4, 5)))
}

func TestListContains(t *testing.T) {
	assert.False(t, fp.ListOf[int]().ContainsElement(1))
	assert.True(t, fp.ListOf(1).ContainsElement(1))
	assert.False(t, fp.ListOf(2, 3, 4).ContainsElement(1))
	assert.True(t, fp.ListOf(1, 2, 3, 4).ContainsElement(1))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).ContainsElement(5))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).ContainsElement(3))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 3, 5).ContainsElement(3))
}

func TestListSize(t *testing.T) {
	assert.Equal(t, 0, fp.ListOf[int]().Size())
	assert.Equal(t, 1, fp.ListOf(1).Size())
	assert.Equal(t, 2, fp.ListOf(1, 2).Size())
	assert.Equal(t, 3, fp.ListOf(1, 2, 3).Size())
}

func TestListExists(t *testing.T) {
	f := func(x int) bool { return x > 5 }
	assert.False(t, fp.ListOf[int]().Exists(f))
	assert.False(t, fp.ListOf(1).Exists(f))
	assert.False(t, fp.ListOf(1, 2, 3).Exists(f))
	assert.True(t, fp.ListOf(6).Exists(f))
	assert.True(t, fp.ListOf(1, 5, 8, 10).Exists(f))
}

func TestListFilter(t *testing.T) {
	f := func(x int) bool { return x > 0 }
	assert.True(t, fp.ListOf[int]().Filter(f).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(-1).Filter(f).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Filter(f).Equals(fp.ListOf[int](1)))
	assert.True(t, fp.ListOf(-5, 6, -7, 8, 9).Filter(f).Equals(fp.ListOf(6, 8, 9)))
	assert.True(t, fp.ListOf(1, 2, 3).Filter(f).Equals(fp.ListOf(1, 2, 3)))
}

func TestListFilterNot(t *testing.T) {
	f := func(x int) bool { return x < 0 }
	assert.True(t, fp.ListOf[int]().FilterNot(f).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(-1).FilterNot(f).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).FilterNot(f).Equals(fp.ListOf[int](1)))
	assert.True(t, fp.ListOf(-5, 6, -7, 8, 9).FilterNot(f).Equals(fp.ListOf(6, 8, 9)))
	assert.True(t, fp.ListOf(1, 2, 3).FilterNot(f).Equals(fp.ListOf(1, 2, 3)))
}

func TestListFind(t *testing.T) {
	f := func(x int) bool { return x > 0 }
	assert.Equal(t, 0, fp.ListOf[int]().Find(f).GetOrElse(0))
	assert.Equal(t, 0, fp.ListOf(-1).Find(f).GetOrElse(0))
	assert.Equal(t, 1, fp.ListOf(1).Find(f).GetOrElse(0))
	assert.Equal(t, 6, fp.ListOf(-5, 6, -7, 8, 9).Find(f).GetOrElse(0))
	assert.Equal(t, 1, fp.ListOf(1, 2, 3).Find(f).GetOrElse(0))

	assert.Equal(t, 1, fp.ListOf(1, 2, 3).Find(func(x int) bool { return x == 1 }).GetOrElse(0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3).Find(func(x int) bool { return x == 2 }).GetOrElse(0))
	assert.Equal(t, 3, fp.ListOf(1, 2, 3).Find(func(x int) bool { return x == 3 }).GetOrElse(0))
}

func TestListDiff(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Diff(fp.ListOf[int]()).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf[int]().Diff(fp.ListOf(1)).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Diff(fp.ListOf[int]()).Equals(fp.ListOf(1)))

	assert.True(t, fp.ListOf(1).Diff(fp.ListOf(1)).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Diff(fp.ListOf(2)).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Diff(fp.ListOf(2, 4)).Equals(fp.ListOf(1, 3, 5)))
	assert.True(t, fp.ListOf(1, 2, 2, 3, 3, 4, 5, 5).Diff(fp.ListOf(2, 4)).Equals(fp.ListOf(1, 3, 3, 5, 5)))
}

func TestListDistinct(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Distinct().Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Distinct().Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Distinct().Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 2, 3, 3, 4, 5, 5).Distinct().Equals(fp.ListOf(1, 2, 3, 4, 5)))
}

func TestListDrop(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Drop(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Drop(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Drop(5).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Drop(6).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Drop(7).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3).Drop(0).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3).Drop(-1).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Drop(2).Equals(fp.ListOf(3, 4, 5)))
}

func TestListDropRight(t *testing.T) {
	assert.True(t, fp.ListOf[int]().DropRight(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).DropRight(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropRight(5).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropRight(6).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropRight(7).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3).DropRight(0).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3).DropRight(-1).Equals(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropRight(2).Equals(fp.ListOf(1, 2, 3)))
}

func TestListDropWhile(t *testing.T) {
	assert.True(t, fp.ListOf[int]().DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(-1, -2, 3, 4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf(3, 4, 5)))
	assert.True(t, fp.ListOf(-1, -2, 3, -4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf(3, -4, 5)))
}

func TestListTake(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Take(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).Take(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Take(5).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Take(6).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Take(7).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3).Take(0).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3).Take(-1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Take(2).Equals(fp.ListOf(1, 2)))
}

func TestListTakeRight(t *testing.T) {
	assert.True(t, fp.ListOf[int]().TakeRight(1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).TakeRight(1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeRight(5).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeRight(6).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeRight(7).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3).TakeRight(0).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3).TakeRight(-1).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeRight(2).Equals(fp.ListOf(4, 5)))
}

func TestListTakeWhile(t *testing.T) {
	assert.True(t, fp.ListOf[int]().TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(-1, -2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf(-1, -2)))
	assert.True(t, fp.ListOf(-1, -2, 3, -4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ListOf(-1, -2)))
}

func TestListForAll(t *testing.T) {
	assert.True(t, fp.ListOf[int]().ForAll(func(x int) bool { return x == 0 }))
	assert.True(t, fp.ListOf(1).ForAll(func(x int) bool { return x == 1 }))
	assert.False(t, fp.ListOf(1).ForAll(func(x int) bool { return x == 0 }))
	assert.True(t, fp.ListOf[int](1, 2, 3).ForAll(func(x int) bool { return x > 0 }))
	assert.False(t, fp.ListOf[int](0, 1, 2).ForAll(func(x int) bool { return x > 0 }))
}

func TestListForEach(t *testing.T) {
	var s = ""
	f := func(x int) fp.Unit {
		s = s + " " + fmt.Sprint(x)
		return fp.GetUnit()
	}

	fp.ListOf[int]().ForEach(f)
	assert.Equal(t, "", s)

	fp.ListOf(1, 2, 3, 4, 5).ForEach(f)
	assert.Equal(t, "1 2 3 4 5", strings.Trim(s, " "))
}

func TestListIndexes(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Indexes().IsEmpty())
	assert.True(t, fp.ListOf(1).Indexes().Equals(fp.ListOf(0)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Indexes().Equals(fp.ListOf(0, 1, 2, 3, 4)))
}

func TestListZipWithIndex(t *testing.T) {
	exp := fp.ListOf(
		fp.PairOf("zero", 0),
		fp.PairOf("one", 1),
		fp.PairOf("two", 2),
	)

	assert.True(t, fp.ListZipWithIndex(fp.ListOf[string]()).IsEmpty())
	assert.True(t, fp.ListZipWithIndex(fp.ListOf("zero", "one", "two")).Equals(exp))
}
