package fp

import "reflect"

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func ListOf[T any](elements ...T) Seq[T] {
	return seqOf(emptyList[T], elements)
}

func ListRangeStep(from int, n int, step int) Seq[int] {
	return seqRangeStep(emptyList[int], from, n, step)
}

func ListRange(from int, n int) Seq[int] {
	return seqRange(emptyList[int], from, n)
}

func ListTabulate[T any](n int, f func(int) T) Seq[T] {
	return seqTabulate(emptyList[int], emptyList[T], n, f)
}

func ListFill[T any](n int, e T) Seq[T] {
	return seqFill(emptyList[int], emptyList[T], n, e)
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
	if n <= 0 {
		return this
	}

	f := func(i int, e T, acc Seq[T]) Seq[T] {
		if i < n {
			return acc
		}
		return acc.Add(e)
	}

	return iterateCount[T, Seq[T]](this, f, emptyList[T]()).Reverse()
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
	if n <= 0 {
		return emptyList[T]()
	}

	f := func(i int, e T, acc Seq[T]) Seq[T] {
		if i < n {
			return acc.Add(e)
		}
		return acc
	}

	return iterateCount[T, Seq[T]](this, f, emptyList[T]()).Reverse()
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
