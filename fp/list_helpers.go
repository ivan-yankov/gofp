package fp

func emptyList[T any]() Seq[T] {
	return List[T]{
		head:  *new(T),
		tail:  nil,
		empty: true,
	}
}

func iterateAdd[T any](seq Seq[T], acc Seq[T]) Seq[T] {
	if seq.IsEmpty() {
		return acc
	} else {
		v, _ := seq.HeadOption().Get()
		return iterateAdd(seq.Tail(), acc.Add(v))
	}
}
