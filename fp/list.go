package fp

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

type List[T any] struct {
	head T
	tail Seq[T]
	size int
}

func (this List[T]) Add(e T) Seq[T] {
	if this.NonEmpty() {
		return List[T]{
			head: e,
			tail: List[T]{
				head: this.head,
				tail: this.tail,
				size: this.size,
			},
			size: this.size + 1,
		}
	}

	return List[T]{
		head: e,
		tail: nil,
		size: 1,
	}
}

func (this List[T]) Get(i int) T {
	if this.IsValidIndex(i) {
		var it func(Seq[T], int) T
		it = func(seq Seq[T], ind int) T {
			if i == ind {
				return seq.HeadOption().Get()
			}
			return it(seq.Tail(), ind+1)
		}
		return it(this, 0)
	}
	panic("Index " + fmt.Sprint(i) + " out of bounds " + fmt.Sprint(this.Size()))
}

func (this List[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this List[T]) NonEmpty() bool {
	return this.Size() != 0
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
	return SeqFoldLeft[T, Seq[T]](this, add[T], emptyList[T]())
}

func (this List[T]) Append(e T) Seq[T] {
	return this.Reverse().Add(e).Reverse()
}

func (this List[T]) Concat(that Seq[T]) Seq[T] {
	return SeqFoldRight[T, Seq[T]](this, add[T], that)
}

func (this List[T]) ContainsElement(e T) bool {
	f := func(ei T, acc bool) bool { return acc || reflect.DeepEqual(e, ei) }
	return SeqFoldLeft[T, bool](this, f, false)
}

func (this List[T]) Size() int {
	return this.size
}

func (this List[T]) Exists(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc || p(e) }
	return SeqFoldLeft[T, bool](this, f, false)
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

func (this List[T]) DropWhile(p func(T) bool) Seq[T] {
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

func (this List[T]) TakeWhile(p func(T) bool) Seq[T] {
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

	r := SeqFoldLeft[T, Acc](this, f, Acc{emptyList[T](), true})
	return r.result.Reverse()
}

func (this List[T]) ForAll(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc && p(e) }
	return SeqFoldLeft[T, bool](this, f, true)
}

func (this List[T]) ForEach(f func(T) Unit) Unit {
	fi := func(e T, acc Unit) Unit { return f(e) }
	return SeqFoldLeft[T, Unit](this, fi, GetUnit())
}

func (this List[T]) ForEachPar(f func(T) Unit) Unit {
	var wg sync.WaitGroup
	fi := func(e T, acc Unit) Unit {
		wg.Add(1)
		go func() Unit {
			defer wg.Done()
			return f(e)
		}()
		return GetUnit()
	}
	SeqFoldLeft[T, Unit](this, fi, GetUnit())
	wg.Wait()
	return GetUnit()
}

func (this List[T]) Indexes() Seq[int] {
	return ListRange(0, this.Size())
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

func (this List[T]) Min(less func(T, T) bool) Option[T] {
	if this.IsEmpty() {
		return None[T]()
	}

	f := func(x T, y T) T {
		if less(x, y) {
			return x
		}
		return y
	}

	r := SeqFoldLeft[T, T](this, f, this.head)
	return SomeOf(r)
}

func (this List[T]) Max(less func(T, T) bool) Option[T] {
	if this.IsEmpty() {
		return None[T]()
	}

	f := func(x T, y T) T {
		if less(x, y) {
			return y
		}
		return x
	}

	r := SeqFoldLeft[T, T](this, f, this.head)
	return SomeOf(r)
}

func (this List[T]) MkString(sep string) string {
	return mkString[T](this, sep)
}

func (this List[T]) PrefixLength(p func(T) bool) int {
	return prefixLength[T](this, p)
}

func (this List[T]) Reduce(f func(T, T) T) Option[T] {
	return reduce[T](this, f)
}

func (this List[T]) Slice(from int, until int) Seq[T] {
	lo := maxInt(from, 0)
	if until <= lo || this.IsEmpty() {
		return emptyList[T]()
	}

	return this.Drop(lo).Take(until - lo)
}

func (this List[T]) FindSlice(that Seq[T]) Option[int] {
	if this.IsEmpty() || that.IsEmpty() || (this.Size() < that.Size()) {
		return None[int]()
	}

	var it func(Seq[T], int, Option[int]) Option[int]
	it = func(seq Seq[T], i int, acc Option[int]) Option[int] {
		if acc.IsDefined() || seq.IsEmpty() {
			return acc
		} else if seq.Take(that.Size()).Equals(that) {
			return SomeOf(i)
		}
		return it(seq.Tail(), i+1, acc)
	}
	return it(this, 0, None[int]())
}

func (this List[T]) StartsWith(that Seq[T]) bool {
	if this.IsEmpty() || that.IsEmpty() || (this.Size() < that.Size()) {
		return false
	}
	return this.Take(that.Size()).Equals(that)
}

func (this List[T]) EndsWith(that Seq[T]) bool {
	return this.Reverse().StartsWith(that.Reverse())
}

func (this List[T]) SplitAt(i int) Pair[Seq[T], Seq[T]] {
	if this.IsEmpty() {
		return PairOf(emptyList[T](), emptyList[T]())
	}
	if !this.IsValidIndex(i) {
		return PairOf(emptyList[T]().Concat(this), emptyList[T]())
	}
	return PairOf(this.Take(i), this.Drop(i))
}

func (this List[T]) Sort(less func(T, T) bool) Seq[T] {
	s := this.ToGoSlice()
	sort.Slice(
		s,
		func(i int, j int) bool { return less(s[i], s[j]) },
	)
	return ListOfGoSlice(s)
}

func (this List[T]) ToList() List[T] {
	return this
}

func (this List[T]) ToArray() Array[T] {
	return Array[T]{this.ToGoSlice()}
}

func (this List[T]) ToGoSlice() []T {
	var it func(Seq[T], []T) []T
	it = func(seq Seq[T], acc []T) []T {
		if seq.IsEmpty() {
			return acc
		}
		return it(seq.Tail(), append(acc, seq.HeadOption().Get()))
	}
	return it(this, make([]T, 0))
}

func (this List[T]) IsList() bool {
	return true
}
