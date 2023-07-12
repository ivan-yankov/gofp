package fp

import "reflect"

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func ListOf[T any](elements ...T) Seq[T] {
	i := len(elements) - 1
	inc := func(x int) int { return x - 1 }
	p := func(x int) bool { return x < 0 }
	f := func(x int, acc Seq[T]) Seq[T] { return acc.Add(elements[x]) }
	return loop(i, inc, p, f, emptyList[T]())
}

func ListRangeStep(from int, n int, step int) Seq[int] {
	if n <= 0 || step < 0 {
		return emptyList[int]()
	}

	inc := func(x int) int { return x + 1 }
	p := func(x int) bool { return x == n }
	f := func(x int, acc Seq[int]) Seq[int] { return acc.Add(from + x*step) }
	return loop(0, inc, p, f, emptyList[int]()).Reverse()
}

func ListRange(from int, n int) Seq[int] {
	return ListRangeStep(from, n, 1)
}

func ListTabulate[T any](n int, f func(int) T) Seq[T] {
	indexes := ListRange(0, n)
	fAcc := func(i int, acc Seq[T]) Seq[T] {
		return acc.Add(f(i))
	}

	return iterate[int, Seq[T]](indexes, fAcc, emptyList[T]()).Reverse()
}

func ListFill[T any](n int, e T) Seq[T] {
	return ListTabulate(n, func(i int) T { return e })
}

func (x List[T]) Add(e T) Seq[T] {
	if x.NonEmpty() {
		return List[T]{
			head: e,
			tail: List[T]{
				head:  x.head,
				tail:  x.tail,
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
	return iterate[T, Seq[T]](this, add[T], emptyList[T]())
}

func (this List[T]) Append(e T) Seq[T] {
	return this.Reverse().Add(e).Reverse()
}

func (this List[T]) Concat(that Seq[T]) Seq[T] {
	return iterate(that, add[T], this.Reverse()).Reverse()
}

func (this List[T]) ContainsElement(e T) bool {
	f := func(ei T, acc bool) bool { return acc || reflect.DeepEqual(e, ei) }
	return iterate[T, bool](this, f, false)
}

func (this List[T]) Size() int {
	f := func(_ T, acc int) int { return acc + 1 }
	return iterate[T, int](this, f, 0)
}

func (this List[T]) Exists(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc || p(e) }
	return iterate[T, bool](this, f, false)
}

func (this List[T]) Filter(p func(T) bool) Seq[T] {
	f := func(e T, acc Seq[T]) Seq[T] {
		if p(e) {
			return acc.Add(e)
		}
		return acc
	}

	return iterate[T, Seq[T]](this, f, emptyList[T]()).Reverse()
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
	f := func(e T, acc Seq[T]) Seq[T] {
		if acc.ContainsElement(e) {
			return acc
		}
		return acc.Add(e)
	}
	return iterate[T, Seq[T]](this, f, emptyList[T]()).Reverse()
}
