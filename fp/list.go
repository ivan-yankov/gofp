package fp

import "reflect"

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func ListTabulate[T any](n int, f func(int) T) Seq[T] {
	var loop func(int, Seq[T]) Seq[T]
	loop = func(i int, acc Seq[T]) Seq[T] {
		if i == n {
			return acc
		} else {
			return loop(i+1, acc.Add(f(i)))
		}
	}

	return loop(0, emptyList[T]()).Reverse()
}

func ListOf[T any](elements ...T) Seq[T] {
	return ListTabulate(len(elements), func(i int) T { return elements[i] })
}

func ListFill[T any](n int, e T) Seq[T] {
	return ListTabulate(n, func(i int) T { return e })
}

func ListRange(from int, n int) Seq[int] {
	return ListTabulate(n, func(i int) int { return from + i })
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
	} else {
		return None[T]()
	}
}

func (this List[T]) LastOption() Option[T] {
	return this.Reverse().HeadOption()
}

func (this List[T]) Tail() Seq[T] {
	if this.tail == nil {
		return emptyList[T]()
	} else {
		return this.tail
	}
}

func (this List[T]) Equals(that Seq[T]) bool {
	return reflect.DeepEqual(this, that)
}

func (this List[T]) Reverse() Seq[T] {
	return iterateAdd[T](this, emptyList[T]())
}

func (this List[T]) Append(e T) Seq[T] {
	return this.Reverse().Add(e).Reverse()
}

func (this List[T]) Concat(that Seq[T]) Seq[T] {
	return iterateAdd(that, this.Reverse()).Reverse()
}
