package fp

func emptyList[T any]() Seq[T] {
	return List[T]{
		head:  *new(T),
		tail:  nil,
		empty: true,
	}
}

func iterate[A, B any](seq Seq[A], acc B, f func(A, B) B) B {
	if seq.IsEmpty() {
		return acc
	}

	v, _ := seq.HeadOption().Get()
	return iterate(seq.Tail(), f(v, acc), f)
}

func add[T any](e T, acc Seq[T]) Seq[T] {
	return acc.Add(e)
}
