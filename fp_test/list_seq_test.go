package fp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

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

func TestListIndexOf(t *testing.T) {
	assert.Equal(t, -1, fp.ListOf[int]().IndexOf(1))
	assert.Equal(t, 0, fp.ListOf(1).IndexOf(1))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).IndexOf(1))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).IndexOf(1))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).IndexOf(5))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).IndexOf(3))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOf(3))
}

func TestListIndexOfFrom(t *testing.T) {
	assert.Equal(t, -1, fp.ListOf[int]().IndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ListOf(1).IndexOfFrom(1, 0))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).IndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).IndexOfFrom(1, 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).IndexOfFrom(5, 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).IndexOfFrom(3, 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfFrom(3, 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfFrom(3, 2))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5, 6, 3, 9).IndexOfFrom(3, 3))
	assert.Equal(t, -1, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfFrom(2, 2))
}

func TestListIndexOfWhere(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ListOf[int]().IndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ListOf(1).IndexOfWhere(p(1)))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).IndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).IndexOfWhere(p(1)))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).IndexOfWhere(p(5)))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).IndexOfWhere(p(3)))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfWhere(p(3)))
}

func TestListIndexOfWhereFrom(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ListOf[int]().IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ListOf(1).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).IndexOfWhereFrom(p(5), 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).IndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(3), 2))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5, 6, 3, 9).IndexOfWhereFrom(p(3), 3))
	assert.Equal(t, -1, fp.ListOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(2), 2))
}

func TestListLastIndexOf(t *testing.T) {
	assert.Equal(t, -1, fp.ListOf[int]().LastIndexOf(1))
	assert.Equal(t, 0, fp.ListOf(1).LastIndexOf(1))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).LastIndexOf(1))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).LastIndexOf(1))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).LastIndexOf(5))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).LastIndexOf(3))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOf(3))
}

func TestListLastIndexOfFrom(t *testing.T) {
	assert.Equal(t, -1, fp.ListOf[int]().LastIndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ListOf(1).LastIndexOfFrom(1, 0))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).LastIndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).LastIndexOfFrom(1, 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfFrom(5, 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfFrom(3, 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(3, 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(3, 2))
	assert.Equal(t, 7, fp.ListOf(1, 2, 3, 4, 3, 5, 6, 3, 9).LastIndexOfFrom(3, 3))
	assert.Equal(t, -1, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(2, 2))
}

func TestListLastIndexOfWhere(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ListOf[int]().LastIndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ListOf(1).LastIndexOfWhere(p(1)))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).LastIndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).LastIndexOfWhere(p(1)))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfWhere(p(5)))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfWhere(p(3)))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfWhere(p(3)))
}

func TestListLastIndexOfWhereFrom(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ListOf[int]().LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ListOf(1).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, -1, fp.ListOf(2, 3, 4).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfWhereFrom(p(5), 0))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).LastIndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 4, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(3), 2))
	assert.Equal(t, 7, fp.ListOf(1, 2, 3, 4, 3, 5, 6, 3, 9).LastIndexOfWhereFrom(p(3), 3))
	assert.Equal(t, -1, fp.ListOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(2), 2))
}

func TestListIsValidIndex(t *testing.T) {
	assert.False(t, fp.ListOf[int]().IsValidIndex(-1))
	assert.False(t, fp.ListOf[int]().IsValidIndex(0))
	assert.False(t, fp.ListOf[int]().IsValidIndex(1))
	assert.False(t, fp.ListOf(1, 2, 3).IsValidIndex(-1))
	assert.False(t, fp.ListOf(1, 2, 3).IsValidIndex(3))
	assert.False(t, fp.ListOf(1, 2, 3).IsValidIndex(4))
	assert.True(t, fp.ListOf(1, 2, 3).IsValidIndex(0))
	assert.True(t, fp.ListOf(1, 2, 3).IsValidIndex(1))
	assert.True(t, fp.ListOf(1, 2, 3).IsValidIndex(2))
}

func TestListStartsWith(t *testing.T) {
	assert.False(t, fp.ListOf[int]().StartsWith(fp.ListOf[int]()))
	assert.False(t, fp.ListOf[int]().StartsWith(fp.ListOf(1)))
	assert.False(t, fp.ListOf(1).StartsWith(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).StartsWith(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3).StartsWith(fp.ListOf(1, 2, 3)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).StartsWith(fp.ListOf(1, 2, 3)))
	assert.False(t, fp.ListOf(1, 2, 3, 4, 5).StartsWith(fp.ListOf(4, 5, 6)))
	assert.False(t, fp.ListOf(1, 2, 3).StartsWith(fp.ListOf(1, 2, 3, 4, 5)))
}

func TestListEndsWith(t *testing.T) {
	assert.False(t, fp.ListOf[int]().EndsWith(fp.ListOf[int]()))
	assert.False(t, fp.ListOf[int]().EndsWith(fp.ListOf(1)))
	assert.False(t, fp.ListOf(1).EndsWith(fp.ListOf[int]()))
	assert.True(t, fp.ListOf(1).EndsWith(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3).EndsWith(fp.ListOf(1, 2, 3)))
	assert.False(t, fp.ListOf(1, 2, 3, 4, 5).EndsWith(fp.ListOf(4, 5, 6)))
	assert.False(t, fp.ListOf(1, 2, 3).EndsWith(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).EndsWith(fp.ListOf(3, 4, 5)))
	assert.False(t, fp.ListOf(1, 2, 3, 4, 5).EndsWith(fp.ListOf(1, 2, 3, 4, 5, 6)))
	assert.False(t, fp.ListOf(1, 2, 3, 4, 5).EndsWith(fp.ListOf(6, 1, 2, 3, 4, 5)))
}

func TestListFindSlice(t *testing.T) {
	assert.True(t, fp.ListOf[int]().FindSlice(fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.ListOf(1).FindSlice(fp.ListOf[int]()).NonDefined())
	assert.True(t, fp.ListOf[int]().FindSlice(fp.ListOf(1)).NonDefined())
	assert.True(t, fp.ListOf(1, 2, 3).FindSlice(fp.ListOf(1, 2, 3, 4, 5)).NonDefined())

	assert.Equal(t, 0, fp.ListOf(1, 2, 3, 4, 5).FindSlice(fp.ListOf(1, 2, 3, 4, 5)).GetOrElse(-1))
	assert.Equal(t, 3, fp.ListOf(1, 2, 3, 4, 5).FindSlice(fp.ListOf(4, 5)).GetOrElse(-1))
	assert.Equal(t, 2, fp.ListOf(1, 2, 3, 4, 5).FindSlice(fp.ListOf(3, 4, 5)).GetOrElse(-1))
}

func TestListMin(t *testing.T) {
	f := func(x int, y int) bool { return x > y }
	assert.True(t, fp.ListOf[int]().Min(f).NonDefined())
	assert.Equal(t, 1, fp.ListOf(1).Min(f).Get())
	assert.Equal(t, 1, fp.ListOf(1, 2, 3, 4, 5).Min(f).Get())
}

func TestListMax(t *testing.T) {
	f := func(x int, y int) bool { return x > y }
	assert.True(t, fp.ListOf[int]().Max(f).NonDefined())
	assert.Equal(t, 1, fp.ListOf(1).Max(f).Get())
	assert.Equal(t, 5, fp.ListOf(1, 2, 3, 4, 5).Max(f).Get())
}

func TestListMkString(t *testing.T) {
	assert.Equal(t, "", fp.ListOf[int]().MkString(""))
	assert.Equal(t, "", fp.ListOf[int]().MkString("-"))
	assert.Equal(t, "12345", fp.ListOf(1, 2, 3, 4, 5).MkString(""))
	assert.Equal(t, "1-2-3-4-5", fp.ListOf(1, 2, 3, 4, 5).MkString("-"))

	type person struct {
		name string
		age  int
	}

	persons := fp.ListOf(
		person{"P1", 24},
		person{"P2", 28},
		person{"P3", 34},
	)

	assert.Equal(t, "{name:P1 age:24};{name:P2 age:28};{name:P3 age:34}", persons.MkString(";"))
}

func TestListPrefixLength(t *testing.T) {
	p := func(x int) bool { return x > 0 }
	assert.Equal(t, 0, fp.ListOf[int]().PrefixLength(p))
	assert.Equal(t, 0, fp.ListOf[int]().PrefixLength(p))
	assert.Equal(t, 0, fp.ListOf(0, 1, 2, 3).PrefixLength(p))
	assert.Equal(t, 3, fp.ListOf(1, 2, 3, -1, 4, 5).PrefixLength(p))
}

func TestListReduce(t *testing.T) {
	sum := func(x int, y int) int { return x + y }
	assert.True(t, fp.ListOf[int]().Reduce(sum).NonDefined())
	assert.Equal(t, 1, fp.ListOf(1).Reduce(sum).Get())
	assert.Equal(t, 15, fp.ListOf(1, 2, 3, 4, 5).Reduce(sum).Get())

	product := func(x int, y int) int { return x * y }
	assert.True(t, fp.ListOf[int]().Reduce(product).NonDefined())
	assert.Equal(t, 1, fp.ListOf(1).Reduce(product).Get())
	assert.Equal(t, 120, fp.ListOf(1, 2, 3, 4, 5).Reduce(product).Get())
}

func TestListSlice(t *testing.T) {
	assert.True(t, fp.ListOf[int]().Slice(0, 1).IsEmpty())
	assert.True(t, fp.ListOf[int]().Slice(1, 0).IsEmpty())
	assert.True(t, fp.ListOf(1, 2, 3).Slice(1, 0).IsEmpty())

	assert.True(t, fp.ListOf(1, 2, 3).Slice(0, 1).Equals(fp.ListOf(1)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Slice(0, 4).Equals(fp.ListOf(1, 2, 3, 4)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Slice(0, 5).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Slice(0, 6).Equals(fp.ListOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Slice(1, 4).Equals(fp.ListOf(2, 3, 4)))
	assert.True(t, fp.ListOf(1, 2, 3, 4, 5).Slice(2, 9).Equals(fp.ListOf(3, 4, 5)))
}

func TestListToList(t *testing.T) {
	assert.True(t, fp.ListOf(1, 2, 3).ToList().Equals(fp.ListOf(1, 2, 3)))
}

func TestListToGoSlice(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, s, fp.ListOf(1, 2, 3).ToGoSlice())
}
