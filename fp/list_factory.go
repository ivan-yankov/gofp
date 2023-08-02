package fp

func ListOfGoSlice[T any](elements []T) Seq[T] {
	i := len(elements) - 1
	inc := func(x int) int { return x - 1 }
	p := func(x int) bool { return x < 0 }
	f := func(x int, acc Seq[T]) Seq[T] { return acc.Add(elements[x]) }
	return loop(i, inc, p, f, emptyList[T]())
}

func ListOf[T any](elements ...T) Seq[T] {
	return ListOfGoSlice(elements)
}

func ListRangeStep(from int, n int, step int) Seq[int] {
	return seqRangeStep(from, n, step, true)
}

func ListRange(from int, n int) Seq[int] {
	return seqRange(from, n, true)
}

func ListTabulate[T any](n int, f func(int) T) Seq[T] {
	return seqTabulate(n, f, true)
}

func ListFill[T any](n int, e T) Seq[T] {
	return seqFill(n, e, true)
}
