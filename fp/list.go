package fp

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func ListZip[A, B any](a Seq[A], b Seq[B]) Seq[Pair[A, B]] {
	type T = Seq[Pair[A, B]]

	var it func(Seq[A], Seq[B], T) T
	it = func(sa Seq[A], sb Seq[B], acc T) T {
		if sa.IsEmpty() || sb.IsEmpty() {
			return acc
		}
		return it(sa.Tail(), sb.Tail(), acc.Add(PairOf(sa.HeadOption().Get(), sb.HeadOption().Get())))
	}

	return it(a, b, emptyList[Pair[A, B]]()).Reverse()
}

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

func ListMap[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	return ListReverseMap(seq, f).Reverse()
}

func ListReverseMap[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	var it func(Seq[A], Seq[B]) Seq[B]
	it = func(s Seq[A], acc Seq[B]) Seq[B] {
		if s.IsEmpty() {
			return acc
		}
		return it(s.Tail(), acc.Add(f(s.HeadOption().Get())))
	}
	return it(seq, emptyList[B]())
}

func ListFlatMap[A, B any](seq Seq[A], f func(A) Seq[B]) Seq[B] {
	var it func(Seq[A], Seq[B]) Seq[B]
	it = func(s Seq[A], acc Seq[B]) Seq[B] {
		if s.IsEmpty() {
			return acc
		}
		return it(s.Tail(), acc.Concat(f(s.HeadOption().Get())))
	}
	return it(seq, emptyList[B]())
}

func ListSliding[T any](seq Seq[T], size int, step int) Seq[Seq[T]] {
	if size <= 0 || step <= 0 {
		return emptyList[Seq[T]]()
	}

	var it func(Seq[T], Seq[Seq[T]]) Seq[Seq[T]]
	it = func(s Seq[T], acc Seq[Seq[T]]) Seq[Seq[T]] {
		if s.Size() <= size {
			if s.NonEmpty() {
				return acc.Add(s.Take(size))
			}
			return acc
		}
		return it(s.Drop(step), acc.Add(s.Take(size)))
	}
	return it(seq, emptyList[Seq[T]]()).Reverse()
}
