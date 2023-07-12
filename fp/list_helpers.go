package fp

func emptyList[T any]() Seq[T] {
	return List[T]{
		head:  *new(T),
		tail:  nil,
		empty: true,
	}
}

func iterate[A, B any](seq Seq[A], f func(A, B) B, acc B) B {
	if seq.IsEmpty() {
		return acc
	}

	v, _ := seq.HeadOption().Get()
	return iterate(seq.Tail(), f, f(v, acc))
}

func iterateCount[A, B any](seq Seq[A], f func(int, A, B) B, acc B) B {
	var it func(seq Seq[A], f func(int, A, B) B, acc B, i int) B
	it = func(seq Seq[A], f func(int, A, B) B, acc B, i int) B {
		if seq.IsEmpty() {
			return acc
		}

		v, _ := seq.HeadOption().Get()
		return it(seq.Tail(), f, f(i, v, acc), i+1)
	}

	return it(seq, f, acc, 0)
}

func add[T any](e T, acc Seq[T]) Seq[T] {
	return acc.Add(e)
}

func loop[T any](
	i int,
	inc func(int) int,
	p func(int) bool,
	f func(int, T) T,
	acc T) T {

	if p(i) {
		return acc
	}
	return loop(inc(i), inc, p, f, f(i, acc))
}
