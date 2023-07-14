package fp

import "reflect"

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func ListZip[A, B any](a Seq[A], b Seq[B]) Seq[Pair[A, B]] {
	type T = Seq[Pair[A, B]]

	var it func(Seq[A], Seq[B], T) T
	it = func(s1 Seq[A], s2 Seq[B], acc T) T {
		if s1.IsEmpty() || s2.IsEmpty() {
			return acc
		}

		return it(s1.Tail(), s2.Tail(), acc.Add(PairOf(s1.HeadOption().Get(), s2.HeadOption().Get())))
	}

	return it(a, b, emptyList[Pair[A, B]]()).Reverse()
}

// implemented not as an interface method due to generic instantiation cycle error
// same is applicable to the functions which use it
func ListZipWithIndex[T any](seq Seq[T]) Seq[Pair[T, int]] {
	return ListZip[T, int](seq, seq.Indexes())
}

func ListFoldLeft[A, B any](seq Seq[A], f func(A, B) B, acc B) B {
	if seq.IsEmpty() {
		return acc
	}

	return ListFoldLeft(seq.Tail(), f, f(seq.HeadOption().Get(), acc))
}

func ListFoldRight[A, B any](seq Seq[A], f func(A, B) B, acc B) B {
	return ListFoldLeft(seq.Reverse(), f, acc)
}

func ListFoldCount[A, B any](seq Seq[A], f func(int, A, B) B, acc B) B {
	var fold func(seq Seq[A], f func(int, A, B) B, acc B, i int) B
	fold = func(seq Seq[A], f func(int, A, B) B, acc B, i int) B {
		if seq.IsEmpty() {
			return acc
		}

		return fold(seq.Tail(), f, f(i, seq.HeadOption().Get(), acc), i+1)
	}

	return fold(seq, f, acc, 0)
}

func ListStartsWith[T any](a Seq[T], b Seq[T]) bool {
	if a.IsEmpty() || b.IsEmpty() || a.Size() < b.Size() {
		return false
	}

	return ListZip[T, T](a, b).
		ForAll(func(x Pair[T, T]) bool {
			return reflect.DeepEqual(x.GetA(), x.GetB())
		})
}

func ListEndsWith[T any](a Seq[T], b Seq[T]) bool {
	return ListStartsWith(a.Reverse(), b.Reverse())
}

func ListContainsSlice[T any](a Seq[T], b Seq[T]) Option[int] {
	if a.IsEmpty() || b.IsEmpty() || (a.Size() < b.Size()) {
		return None[int]()
	}

	var it func(Seq[T], int, Option[int]) Option[int]
	it = func(s Seq[T], i int, acc Option[int]) Option[int] {
		if acc.IsDefined() || s.IsEmpty() {
			return acc
		}
		if ListZip(s, b).ForAll(func(x Pair[T, T]) bool { return reflect.DeepEqual(x.GetA(), x.GetB()) }) {
			return it(s.Tail(), i+1, SomeOf(i))
		}
		return it(s.Tail(), i+1, None[int]())
	}

	return it(a, 0, None[int]())
}
