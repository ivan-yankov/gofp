package fp_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestArrayAdd(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Add(1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1).Add(2).Equals(fp.ArrayOf(2, 1)))
	assert.True(t, fp.ArrayOf(1, 2).Add(3).Equals(fp.ArrayOf(3, 1, 2)))
	assert.True(t, fp.ArrayOf(1, 2).Add(3).Add(4).Equals(fp.ArrayOf(4, 3, 1, 2)))
}

func TestArrayIsEmpty(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().IsEmpty())
	assert.False(t, fp.ArrayOf(1, 2, 3).IsEmpty())
}

func TestArrayNonEmpty(t *testing.T) {
	assert.False(t, fp.ArrayOf[int]().NonEmpty())
	assert.True(t, fp.ArrayOf(1, 2, 3).NonEmpty())
}

func TestArrayHeadOption(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().HeadOption().NonDefined())

	assert.True(t, fp.ArrayOf(1).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).HeadOption().GetOrElse(0))

	assert.True(t, fp.ArrayOf(1, 2, 3).HeadOption().IsDefined())
	assert.Equal(t, 1, fp.ArrayOf(1, 2, 3).HeadOption().GetOrElse(0))
}

func TestArrayLastOption(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().LastOption().NonDefined())

	assert.True(t, fp.ArrayOf(1).LastOption().IsDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).LastOption().GetOrElse(0))

	assert.True(t, fp.ArrayOf(1, 2, 3).LastOption().IsDefined())
	assert.Equal(t, 3, fp.ArrayOf(1, 2, 3).LastOption().GetOrElse(0))
}

func TestArrayTail(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Tail().IsEmpty())

	assert.True(t, fp.ArrayOf(1).Tail().IsEmpty())

	assert.True(t, fp.ArrayOf(1, 2, 3).Tail().NonEmpty())
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3).Tail().HeadOption().GetOrElse(0))
}

func TestArrayEquals(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Equals(fp.ArrayOf(1, 2, 3)))

	assert.False(t, fp.ArrayOf[int]().Equals(fp.ArrayOf(1)))
	assert.False(t, fp.ArrayOf(1, 2).Equals(fp.ArrayOf(2, 1)))
	assert.False(t, fp.ArrayOf(1, 2).Equals(fp.ArrayOf(1, 3)))
	assert.False(t, fp.ArrayOf(1, 2).Equals(fp.ArrayOf(1, 2, 3)))
}

func TestArrayReverse(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Reverse().Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Reverse().Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Reverse().Equals(fp.ArrayOf(3, 2, 1)))
	assert.True(t, fp.ListRange(1, 5).Reverse().Reverse().Equals(fp.ListRange(1, 5)))
}

func TestArrayAppend(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Append(1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1).Append(2).Equals(fp.ArrayOf(1, 2)))
	assert.True(t, fp.ArrayOf(1, 2).Append(3).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2).Append(3).Append(4).Equals(fp.ArrayOf(1, 2, 3, 4)))
}

func TestArrayConcat(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Concat(fp.ArrayOf[int]()).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf[int]().Concat(fp.ArrayOf(1)).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1).Concat(fp.ArrayOf(2, 3)).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Concat(fp.ArrayOf(4, 5)).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
}

func TestArrayContains(t *testing.T) {
	assert.False(t, fp.ArrayOf[int]().ContainsElement(1))
	assert.True(t, fp.ArrayOf(1).ContainsElement(1))
	assert.False(t, fp.ArrayOf(2, 3, 4).ContainsElement(1))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4).ContainsElement(1))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).ContainsElement(5))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).ContainsElement(3))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 3, 5).ContainsElement(3))
}

func TestArraySize(t *testing.T) {
	assert.Equal(t, 0, fp.ArrayOf[int]().Size())
	assert.Equal(t, 1, fp.ArrayOf(1).Size())
	assert.Equal(t, 2, fp.ArrayOf(1, 2).Size())
	assert.Equal(t, 3, fp.ArrayOf(1, 2, 3).Size())
}

func TestArrayExists(t *testing.T) {
	f := func(x int) bool { return x > 5 }
	assert.False(t, fp.ArrayOf[int]().Exists(f))
	assert.False(t, fp.ArrayOf(1).Exists(f))
	assert.False(t, fp.ArrayOf(1, 2, 3).Exists(f))
	assert.True(t, fp.ArrayOf(6).Exists(f))
	assert.True(t, fp.ArrayOf(1, 5, 8, 10).Exists(f))
}

func TestArrayFilter(t *testing.T) {
	f := func(x int) bool { return x > 0 }
	assert.True(t, fp.ArrayOf[int]().Filter(f).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(-1).Filter(f).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Filter(f).Equals(fp.ArrayOf[int](1)))
	assert.True(t, fp.ArrayOf(-5, 6, -7, 8, 9).Filter(f).Equals(fp.ArrayOf(6, 8, 9)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Filter(f).Equals(fp.ArrayOf(1, 2, 3)))
}

func TestArrayFilterNot(t *testing.T) {
	f := func(x int) bool { return x < 0 }
	assert.True(t, fp.ArrayOf[int]().FilterNot(f).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(-1).FilterNot(f).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).FilterNot(f).Equals(fp.ArrayOf[int](1)))
	assert.True(t, fp.ArrayOf(-5, 6, -7, 8, 9).FilterNot(f).Equals(fp.ArrayOf(6, 8, 9)))
	assert.True(t, fp.ArrayOf(1, 2, 3).FilterNot(f).Equals(fp.ArrayOf(1, 2, 3)))
}

func TestArrayFind(t *testing.T) {
	f := func(x int) bool { return x > 0 }
	assert.Equal(t, 0, fp.ArrayOf[int]().Find(f).GetOrElse(0))
	assert.Equal(t, 0, fp.ArrayOf(-1).Find(f).GetOrElse(0))
	assert.Equal(t, 1, fp.ArrayOf(1).Find(f).GetOrElse(0))
	assert.Equal(t, 6, fp.ArrayOf(-5, 6, -7, 8, 9).Find(f).GetOrElse(0))
	assert.Equal(t, 1, fp.ArrayOf(1, 2, 3).Find(f).GetOrElse(0))

	assert.Equal(t, 1, fp.ArrayOf(1, 2, 3).Find(func(x int) bool { return x == 1 }).GetOrElse(0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3).Find(func(x int) bool { return x == 2 }).GetOrElse(0))
	assert.Equal(t, 3, fp.ArrayOf(1, 2, 3).Find(func(x int) bool { return x == 3 }).GetOrElse(0))
}

func TestArrayDiff(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Diff(fp.ArrayOf[int]()).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf[int]().Diff(fp.ArrayOf(1)).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Diff(fp.ArrayOf[int]()).Equals(fp.ArrayOf(1)))

	assert.True(t, fp.ArrayOf(1).Diff(fp.ArrayOf(1)).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Diff(fp.ArrayOf(2)).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Diff(fp.ArrayOf(2, 4)).Equals(fp.ArrayOf(1, 3, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 2, 3, 3, 4, 5, 5).Diff(fp.ArrayOf(2, 4)).Equals(fp.ArrayOf(1, 3, 3, 5, 5)))
}

func TestArrayDistinct(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Distinct().Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Distinct().Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Distinct().Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 2, 3, 3, 4, 5, 5).Distinct().Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
}

func TestArrayDrop(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Drop(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Drop(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Drop(5).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Drop(6).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Drop(7).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3).Drop(0).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Drop(-1).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Drop(2).Equals(fp.ArrayOf(3, 4, 5)))
}

func TestArrayDropRight(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().DropRight(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).DropRight(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropRight(5).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropRight(6).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropRight(7).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3).DropRight(0).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2, 3).DropRight(-1).Equals(fp.ArrayOf(1, 2, 3)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropRight(2).Equals(fp.ArrayOf(1, 2, 3)))
}

func TestArrayDropWhile(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(-1, -2, 3, 4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf(3, 4, 5)))
	assert.True(t, fp.ArrayOf(-1, -2, 3, -4, 5).DropWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf(3, -4, 5)))
}

func TestArrayTake(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Take(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).Take(1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Take(5).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Take(6).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Take(7).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3).Take(0).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3).Take(-1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Take(2).Equals(fp.ArrayOf(1, 2)))
}

func TestArrayTakeRight(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().TakeRight(1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).TakeRight(1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeRight(5).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeRight(6).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeRight(7).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3).TakeRight(0).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3).TakeRight(-1).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeRight(2).Equals(fp.ArrayOf(4, 5)))
}

func TestArrayTakeWhile(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(1).TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf(1)))
	assert.Equal(t, fp.ArrayOf(1), fp.ArrayOf(1).TakeWhile(func(x int) bool { return x > 0 }))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x > 0 }).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf[int]()))
	assert.True(t, fp.ArrayOf(-1, -2, 3, 4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf(-1, -2)))
	assert.True(t, fp.ArrayOf(-1, -2, 3, -4, 5).TakeWhile(func(x int) bool { return x < 0 }).Equals(fp.ArrayOf(-1, -2)))
}

func TestArrayForAll(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().ForAll(func(x int) bool { return x == 0 }))
	assert.True(t, fp.ArrayOf(1).ForAll(func(x int) bool { return x == 1 }))
	assert.False(t, fp.ArrayOf(1).ForAll(func(x int) bool { return x == 0 }))
	assert.True(t, fp.ArrayOf[int](1, 2, 3).ForAll(func(x int) bool { return x > 0 }))
	assert.False(t, fp.ArrayOf[int](0, 1, 2).ForAll(func(x int) bool { return x > 0 }))
}

func TestArrayForEach(t *testing.T) {
	var s = ""
	f := func(x int) fp.Unit {
		s = s + " " + fmt.Sprint(x)
		return fp.GetUnit()
	}

	fp.ArrayOf[int]().ForEach(f)
	assert.Equal(t, "", s)

	fp.ArrayOf(1, 2, 3, 4, 5).ForEach(f)
	assert.Equal(t, "1 2 3 4 5", strings.Trim(s, " "))
}

func TestArrayIndexes(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Indexes().IsEmpty())
	assert.True(t, fp.ArrayOf(1).Indexes().Equals(fp.ArrayOf(0)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Indexes().Equals(fp.ArrayOf(0, 1, 2, 3, 4)))
}

func TestArrayIndexOf(t *testing.T) {
	assert.Equal(t, -1, fp.ArrayOf[int]().IndexOf(1))
	assert.Equal(t, 0, fp.ArrayOf(1).IndexOf(1))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).IndexOf(1))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).IndexOf(1))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).IndexOf(5))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).IndexOf(3))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOf(3))
}

func TestArrayIndexOfFrom(t *testing.T) {
	assert.Equal(t, -1, fp.ArrayOf[int]().IndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ArrayOf(1).IndexOfFrom(1, 0))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).IndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).IndexOfFrom(1, 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfFrom(5, 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfFrom(3, 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfFrom(3, 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfFrom(3, 2))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5, 6, 3, 9).IndexOfFrom(3, 3))
	assert.Equal(t, -1, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfFrom(2, 2))
}

func TestArrayIndexOfWhere(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ArrayOf[int]().IndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ArrayOf(1).IndexOfWhere(p(1)))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).IndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).IndexOfWhere(p(1)))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfWhere(p(5)))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfWhere(p(3)))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfWhere(p(3)))
}

func TestArrayIndexOfWhereFrom(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ArrayOf[int]().IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ArrayOf(1).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).IndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfWhereFrom(p(5), 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).IndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(3), 2))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5, 6, 3, 9).IndexOfWhereFrom(p(3), 3))
	assert.Equal(t, -1, fp.ArrayOf(1, 2, 3, 4, 3, 5).IndexOfWhereFrom(p(2), 2))
}

func TestArrayLastIndexOf(t *testing.T) {
	assert.Equal(t, -1, fp.ArrayOf[int]().LastIndexOf(1))
	assert.Equal(t, 0, fp.ArrayOf(1).LastIndexOf(1))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).LastIndexOf(1))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).LastIndexOf(1))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOf(5))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOf(3))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOf(3))
}

func TestArrayLastIndexOfFrom(t *testing.T) {
	assert.Equal(t, -1, fp.ArrayOf[int]().LastIndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ArrayOf(1).LastIndexOfFrom(1, 0))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).LastIndexOfFrom(1, 0))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).LastIndexOfFrom(1, 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfFrom(5, 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfFrom(3, 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(3, 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(3, 2))
	assert.Equal(t, 7, fp.ArrayOf(1, 2, 3, 4, 3, 5, 6, 3, 9).LastIndexOfFrom(3, 3))
	assert.Equal(t, -1, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfFrom(2, 2))
}

func TestArrayLastIndexOfWhere(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ArrayOf[int]().LastIndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ArrayOf(1).LastIndexOfWhere(p(1)))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).LastIndexOfWhere(p(1)))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).LastIndexOfWhere(p(1)))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfWhere(p(5)))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfWhere(p(3)))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfWhere(p(3)))
}

func TestArrayLastIndexOfWhereFrom(t *testing.T) {
	p := func(i int) func(int) bool { return func(x int) bool { return i == x } }

	assert.Equal(t, -1, fp.ArrayOf[int]().LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ArrayOf(1).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, -1, fp.ArrayOf(2, 3, 4).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 0, fp.ArrayOf(1, 2, 3, 4).LastIndexOfWhereFrom(p(1), 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfWhereFrom(p(5), 0))
	assert.Equal(t, 2, fp.ArrayOf(1, 2, 3, 4, 5).LastIndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(3), 0))
	assert.Equal(t, 4, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(3), 2))
	assert.Equal(t, 7, fp.ArrayOf(1, 2, 3, 4, 3, 5, 6, 3, 9).LastIndexOfWhereFrom(p(3), 3))
	assert.Equal(t, -1, fp.ArrayOf(1, 2, 3, 4, 3, 5).LastIndexOfWhereFrom(p(2), 2))
}

func TestArrayIsValidIndex(t *testing.T) {
	assert.False(t, fp.ArrayOf[int]().IsValidIndex(-1))
	assert.False(t, fp.ArrayOf[int]().IsValidIndex(0))
	assert.False(t, fp.ArrayOf[int]().IsValidIndex(1))
	assert.False(t, fp.ArrayOf(1, 2, 3).IsValidIndex(-1))
	assert.False(t, fp.ArrayOf(1, 2, 3).IsValidIndex(3))
	assert.False(t, fp.ArrayOf(1, 2, 3).IsValidIndex(4))
	assert.True(t, fp.ArrayOf(1, 2, 3).IsValidIndex(0))
	assert.True(t, fp.ArrayOf(1, 2, 3).IsValidIndex(1))
	assert.True(t, fp.ArrayOf(1, 2, 3).IsValidIndex(2))
}

func TestArrayMin(t *testing.T) {
	f := func(x int, y int) bool { return x < y }
	assert.True(t, fp.ArrayOf[int]().Min(f).NonDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).Min(f).Get())
	assert.Equal(t, 1, fp.ArrayOf(1, 2, 3, 4, 5).Min(f).Get())
}

func TestArrayMax(t *testing.T) {
	f := func(x int, y int) bool { return x < y }
	assert.True(t, fp.ArrayOf[int]().Max(f).NonDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).Max(f).Get())
	assert.Equal(t, 5, fp.ArrayOf(1, 2, 3, 4, 5).Max(f).Get())
}

func TestArrayMkString(t *testing.T) {
	assert.Equal(t, "", fp.ArrayOf[int]().MkString(""))
	assert.Equal(t, "", fp.ArrayOf[int]().MkString("-"))
	assert.Equal(t, "12345", fp.ArrayOf(1, 2, 3, 4, 5).MkString(""))
	assert.Equal(t, "1-2-3-4-5", fp.ArrayOf(1, 2, 3, 4, 5).MkString("-"))

	type person struct {
		name string
		age  int
	}

	persons := fp.ArrayOf(
		person{"P1", 24},
		person{"P2", 28},
		person{"P3", 34},
	)

	assert.Equal(t, "{name:P1 age:24};{name:P2 age:28};{name:P3 age:34}", persons.MkString(";"))
}

func TestArrayPrefixLength(t *testing.T) {
	p := func(x int) bool { return x > 0 }
	assert.Equal(t, 0, fp.ArrayOf[int]().PrefixLength(p))
	assert.Equal(t, 0, fp.ArrayOf[int]().PrefixLength(p))
	assert.Equal(t, 0, fp.ArrayOf(0, 1, 2, 3).PrefixLength(p))
	assert.Equal(t, 3, fp.ArrayOf(1, 2, 3, -1, 4, 5).PrefixLength(p))
}

func TestArrayReduce(t *testing.T) {
	sum := func(x int, y int) int { return x + y }
	assert.True(t, fp.ArrayOf[int]().Reduce(sum).NonDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).Reduce(sum).Get())
	assert.Equal(t, 15, fp.ArrayOf(1, 2, 3, 4, 5).Reduce(sum).Get())

	product := func(x int, y int) int { return x * y }
	assert.True(t, fp.ArrayOf[int]().Reduce(product).NonDefined())
	assert.Equal(t, 1, fp.ArrayOf(1).Reduce(product).Get())
	assert.Equal(t, 120, fp.ArrayOf(1, 2, 3, 4, 5).Reduce(product).Get())
}

func TestArraySlice(t *testing.T) {
	assert.True(t, fp.ArrayOf[int]().Slice(0, 1).IsEmpty())
	assert.True(t, fp.ArrayOf[int]().Slice(1, 0).IsEmpty())
	assert.True(t, fp.ArrayOf(1, 2, 3).Slice(1, 0).IsEmpty())

	assert.True(t, fp.ArrayOf(1, 2, 3).Slice(0, 1).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Slice(0, 4).Equals(fp.ArrayOf(1, 2, 3, 4)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Slice(0, 5).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Slice(0, 6).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Slice(1, 4).Equals(fp.ArrayOf(2, 3, 4)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Slice(2, 9).Equals(fp.ArrayOf(3, 4, 5)))
}

func TestArraySplitAt(t *testing.T) {
	assert.Equal(t, fp.PairOf(fp.ArrayOf[int](), fp.ArrayOf[int]()), fp.ArrayOf[int]().SplitAt(0))
	assert.Equal(t, fp.PairOf(fp.ArrayOf(1, 2, 3), fp.ArrayOf[int]()), fp.ArrayOf(1, 2, 3).SplitAt(-1))
	assert.Equal(t, fp.PairOf(fp.ArrayOf[int](), fp.ArrayOf(1, 2, 3, 4, 5)), fp.ArrayOf(1, 2, 3, 4, 5).SplitAt(0))
	assert.Equal(t, fp.PairOf(fp.ArrayOf[int](1, 2), fp.ArrayOf(3, 4, 5)), fp.ArrayOf(1, 2, 3, 4, 5).SplitAt(2))
	assert.Equal(t, fp.PairOf(fp.ArrayOf(1, 2, 3, 4), fp.ArrayOf[int](5)), fp.ArrayOf(1, 2, 3, 4, 5).SplitAt(4))
}

func TestArraySort(t *testing.T) {
	f := func(x int, y int) bool { return x < y }
	assert.True(t, fp.ArrayOf[int]().Sort(f).IsEmpty())
	assert.True(t, fp.ArrayOf(1).Sort(f).Equals(fp.ArrayOf(1)))
	assert.True(t, fp.ArrayOf(1, 2, 3, 4, 5).Sort(f).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
	assert.True(t, fp.ArrayOf(3, 2, 1, 4, 5).Sort(f).Equals(fp.ArrayOf(1, 2, 3, 4, 5)))
}

func TestArrayToList(t *testing.T) {
	assert.Equal(t, fp.ListOf(1, 2, 3), fp.ArrayOf(1, 2, 3).ToList())
}

func TestArrayToArray(t *testing.T) {
	assert.Equal(t, fp.ArrayOf(1, 2, 3), fp.ArrayOf(1, 2, 3).ToArray())
}

func TestArrayToGoSlice(t *testing.T) {
	s := []int{1, 2, 3}
	assert.Equal(t, s, fp.ArrayOf(1, 2, 3).ToGoSlice())
}
