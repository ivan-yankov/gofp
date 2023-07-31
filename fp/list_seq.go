package fp

import (
	"fmt"
	"reflect"
)

func (this List[T]) Add(e T) Seq[T] {
	if this.NonEmpty() {
		return List[T]{
			head: e,
			tail: List[T]{
				head:  this.head,
				tail:  this.tail,
				empty: false,
			},
			empty: false,
		}
	}

	return List[T]{
		head:  e,
		tail:  nil,
		empty: false,
	}
}

func (this List[T]) IsEmpty() bool {
	return this.empty
}

func (this List[T]) NonEmpty() bool {
	return !this.empty
}

func (this List[T]) HeadOption() Option[T] {
	if this.NonEmpty() {
		return SomeOf(this.head)
	}
	return None[T]()
}

func (this List[T]) LastOption() Option[T] {
	return this.Reverse().HeadOption()
}

func (this List[T]) Tail() Seq[T] {
	if this.tail == nil {
		return emptyList[T]()
	}
	return this.tail
}

func (this List[T]) Equals(that Seq[T]) bool {
	return reflect.DeepEqual(this, that)
}

func (this List[T]) Reverse() Seq[T] {
	return ListFoldLeft[T, Seq[T]](this, add[T], emptyList[T]())
}

func (this List[T]) Append(e T) Seq[T] {
	return this.Reverse().Add(e).Reverse()
}

func (this List[T]) Concat(that Seq[T]) Seq[T] {
	a := ListFoldRight[T, Seq[T]](that, add[T], this)
	fmt.Print(a)
	return ListFoldRight[T, Seq[T]](this, add[T], that)
}

func (this List[T]) ContainsElement(e T) bool {
	f := func(ei T, acc bool) bool { return acc || reflect.DeepEqual(e, ei) }
	return ListFoldLeft[T, bool](this, f, false)
}

func (this List[T]) Size() int {
	f := func(_ T, acc int) int { return acc + 1 }
	return ListFoldLeft[T, int](this, f, 0)
}

func (this List[T]) Exists(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc || p(e) }
	return ListFoldLeft[T, bool](this, f, false)
}

func (this List[T]) Filter(p func(T) bool) Seq[T] {
	return collect[T](
		this,
		func(_ int, e T, _ Seq[T]) bool { return p(e) },
		emptyList[T],
	)
}

func (this List[T]) FilterNot(p func(T) bool) Seq[T] {
	f := func(e T) bool { return !p(e) }
	return this.Filter(f)
}

func (this List[T]) Find(p func(T) bool) Option[T] {
	return this.Filter(p).HeadOption()
}

func (this List[T]) Diff(that Seq[T]) Seq[T] {
	f := func(e T) bool { return that.ContainsElement(e) }
	return this.FilterNot(f)
}

func (this List[T]) Distinct() Seq[T] {
	return collect[T](
		this,
		func(_ int, e T, acc Seq[T]) bool { return !acc.ContainsElement(e) },
		emptyList[T],
	)
}

func (this List[T]) Drop(n int) Seq[T] {
	if n <= 0 {
		return this
	}

	return collect[T](
		this,
		func(i int, _ T, _ Seq[T]) bool { return i >= n },
		emptyList[T],
	)
}

func (this List[T]) DropRight(n int) Seq[T] {
	return this.Reverse().Drop(n).Reverse()
}

func (this List[T]) DropWhile(p func(e T) bool) Seq[T] {
	return collect[T](
		this,
		func(_ int, e T, acc Seq[T]) bool { return acc.NonEmpty() || !p(e) },
		emptyList[T],
	)
}

func (this List[T]) Take(n int) Seq[T] {
	if n <= 0 {
		return emptyList[T]()
	}

	return collect[T](
		this,
		func(i int, _ T, _ Seq[T]) bool { return i < n },
		emptyList[T],
	)
}

func (this List[T]) TakeRight(n int) Seq[T] {
	return this.Reverse().Take(n).Reverse()
}

func (this List[T]) TakeWhile(p func(e T) bool) Seq[T] {
	type Acc struct {
		result Seq[T]
		flag   bool
	}

	f := func(e T, acc Acc) Acc {
		if acc.flag && p(e) {
			return Acc{acc.result.Add(e), true}
		}
		return Acc{acc.result, false}
	}

	r := ListFoldLeft[T, Acc](this, f, Acc{emptyList[T](), true})
	return r.result.Reverse()
}

func (this List[T]) ForAll(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc && p(e) }
	return ListFoldLeft[T, bool](this, f, true)
}

func (this List[T]) ForEach(f func(T) Unit) Unit {
	fi := func(e T, acc Unit) Unit { f(e); return GetUnit() }
	return ListFoldLeft[T, Unit](this, fi, GetUnit())
}

func (this List[T]) Indexes() Seq[int] {
	f := func(i int, _ T, acc Seq[int]) Seq[int] { return acc.Add(i) }
	return ListFoldCount[T, Seq[int]](this, f, emptyList[int]()).Reverse()
}

func (this List[T]) IndexOf(e T) int {
	p := func(_ int, ei T, acc Option[int]) bool {
		return acc.NonDefined() && reflect.DeepEqual(e, ei)
	}
	return findIndex[T](this, p)
}

func (this List[T]) IndexOfFrom(e T, from int) int {
	p := func(i int, ei T, acc Option[int]) bool {
		return i >= from && acc.NonDefined() && reflect.DeepEqual(e, ei)
	}
	return findIndex[T](this, p)
}

func (this List[T]) IndexOfWhere(p func(T) bool) int {
	f := func(i int, e T, acc Option[int]) bool {
		return acc.NonDefined() && p(e)
	}
	return findIndex[T](this, f)
}

func (this List[T]) IndexOfWhereFrom(p func(T) bool, from int) int {
	f := func(i int, e T, acc Option[int]) bool {
		return i >= from && acc.NonDefined() && p(e)
	}
	return findIndex[T](this, f)
}

func (this List[T]) LastIndexOf(e T) int {
	p := func(_ int, ei T, acc Option[int]) bool {
		return reflect.DeepEqual(e, ei)
	}
	return findIndex[T](this, p)
}

func (this List[T]) LastIndexOfFrom(e T, from int) int {
	p := func(i int, ei T, acc Option[int]) bool {
		return i >= from && reflect.DeepEqual(e, ei)
	}
	return findIndex[T](this, p)
}

func (this List[T]) LastIndexOfWhere(p func(T) bool) int {
	f := func(i int, e T, acc Option[int]) bool {
		return p(e)
	}
	return findIndex[T](this, f)
}

func (this List[T]) LastIndexOfWhereFrom(p func(T) bool, from int) int {
	f := func(i int, e T, acc Option[int]) bool {
		return i >= from && p(e)
	}
	return findIndex[T](this, f)
}

func (this List[T]) IsValidIndex(i int) bool {
	return i >= 0 && i < this.Size()
}

func (this List[T]) StartsWith(b Seq[T]) bool {
	return 0 == this.ContainsSlice(b).GetOrElse(-1)
}

func (this List[T]) EndsWith(b Seq[T]) bool {
	return this.Reverse().StartsWith(b.Reverse())
}

func (this List[T]) ContainsSlice(that Seq[T]) Option[int] {
	if this.IsEmpty() || that.IsEmpty() || (this.Size() < that.Size()) {
		return None[int]()
	}

	var eq func(sa Seq[T], sb Seq[T], acc bool) bool
	eq = func(sa Seq[T], sb Seq[T], acc bool) bool {
		if sa.IsEmpty() || sb.IsEmpty() {
			return acc
		}
		return eq(
			sa.Tail(),
			sb.Tail(),
			acc && reflect.DeepEqual(
				sa.HeadOption().Get(),
				sb.HeadOption().Get(),
			),
		)
	}

	var it func(Seq[T], int, Option[int]) Option[int]
	it = func(s Seq[T], i int, acc Option[int]) Option[int] {
		if acc.IsDefined() || s.IsEmpty() {
			return acc
		}
		if eq(s, that, true) {
			return it(s.Tail(), i+1, SomeOf(i))
		}
		return it(s.Tail(), i+1, None[int]())
	}

	return it(this, 0, None[int]())
}

func (this List[T]) ToList() List[T] {
	return this
}
