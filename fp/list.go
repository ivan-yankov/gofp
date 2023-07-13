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

func ListZip[A, B any](sa Seq[A], sb Seq[B]) Seq[Pair[A, B]] {
	type T = Seq[Pair[A, B]]

	var it func(Seq[A], Seq[B], T) T
	it = func(s1 Seq[A], s2 Seq[B], acc T) T {
		if s1.IsEmpty() || s2.IsEmpty() {
			return acc
		}

		a, _ := s1.HeadOption().Get()
		b, _ := s2.HeadOption().Get()
		return it(s1.Tail(), s2.Tail(), acc.Add(PairOf(a, b)))
	}

	return it(sa, sb, emptyList[Pair[A, B]]()).Reverse()
}

// implemented not as an interface method due to generic instantiation cycle error
func ListZipWithIndex[T any](seq Seq[T]) Seq[Pair[T, int]] {
	return ListZip[T, int](seq, seq.Indexes())
}

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

func (this List[T]) Drop(n int) Seq[T] {
	return iterateWhile[T](this, n, true)
}

func (this List[T]) DropRight(n int) Seq[T] {
	return this.Reverse().Drop(n).Reverse()
}

func (this List[T]) DropWhile(p func(e T) bool) Seq[T] {
	f := func(e T, acc Seq[T]) Seq[T] {
		if acc.NonEmpty() || !p(e) {
			return acc.Add(e)
		}
		return acc
	}

	return iterate[T, Seq[T]](this, f, emptyList[T]()).Reverse()
}

func (this List[T]) Take(n int) Seq[T] {
	return iterateWhile[T](this, n, false)
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

	r := iterate[T, Acc](this, f, Acc{emptyList[T](), true})
	return r.result.Reverse()
}

func (this List[T]) ForAll(p func(T) bool) bool {
	f := func(e T, acc bool) bool { return acc && p(e) }
	return iterate[T, bool](this, f, true)
}

func (this List[T]) ForEach(f func(T) Unit) Unit {
	fi := func(e T, acc Unit) Unit { f(e); return GetUnit() }
	return iterate[T, Unit](this, fi, GetUnit())
}

func (this List[T]) Indexes() Seq[int] {
	f := func(i int, _ T, acc Seq[int]) Seq[int] { return acc.Add(i) }
	return iterateCount[T, Seq[int]](this, f, emptyList[int]()).Reverse()
}

func (this List[T]) IndexOf(e T) int {
	f := func(i int, ei T, acc Option[int]) Option[int] {
		if reflect.DeepEqual(e, ei) && acc.NonDefined() {
			return SomeOf(i)
		}
		return acc
	}
	return iterateCount[T, Option[int]](this, f, None[int]()).GetOrElse(-1)
}
