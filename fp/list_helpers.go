package fp

func emptyList[T any]() Seq[T] {
	return List[T]{
		head:  *new(T),
		tail:  nil,
		empty: true,
	}
}

func findIndex[T any](seq Seq[T], p func(int, T, Option[int]) bool) int {
	f := func(i int, e T, acc Option[int]) Option[int] {
		if p(i, e, acc) {
			return SomeOf(i)
		}
		return acc
	}

	return ListFoldCount[T, Option[int]](seq, f, None[int]()).GetOrElse(-1)
}
