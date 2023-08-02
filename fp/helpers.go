package fp

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

func emptyList[T any]() Seq[T] {
	return List[T]{
		head: *new(T),
		tail: nil,
		size: 0,
	}
}

func emptyArray[T any]() Seq[T] {
	return Array[T]{[]T{}}
}

func emptySeq[T any](list bool) Seq[T] {
	if list {
		return emptyList[T]()
	}
	return emptyArray[T]()
}

func findIndex[T any](seq Seq[T], p func(int, T, Option[int]) bool) int {
	f := func(i int, e T, acc Option[int]) Option[int] {
		if p(i, e, acc) {
			return SomeOf(i)
		}
		return acc
	}

	return SeqFoldCount[T, Option[int]](seq, f, None[int]()).GetOrElse(-1)
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

	return SeqFoldCount[T, Seq[T]](seq, f, emptySeq()).Reverse()
}

func minInt(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func maxInt(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
