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

func add[T any](e T, acc Seq[T]) Seq[T] {
	return acc.Add(e)
}

func collect[T any](
	seq Seq[T],
	appendCondition func(int, T, Seq[T]) bool,
	emptySeq func() Seq[T]) Seq[T] {

	f := func(i int, e T, acc Seq[T]) Seq[T] {
		if appendCondition(i, e, acc) {
			return acc.Add(e)
		}
		return acc
	}

	return ListFoldCount[T, Seq[T]](seq, f, emptySeq()).Reverse()
}
