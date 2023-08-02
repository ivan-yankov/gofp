package fp

func ArrayOfGoSlice[T any](elements []T) Seq[T] {
	if len(elements) == 0 {
		return emptyArray[T]()
	}
	return Array[T]{elements}
}

func ArrayOf[T any](elements ...T) Seq[T] {
	return ArrayOfGoSlice(elements)
}

func ArrayRangeStep(from int, n int, step int) Seq[int] {
	return seqRangeStep(from, n, step, false)
}

func ArrayRange(from int, n int) Seq[int] {
	return seqRange(from, n, false)
}

func ArrayTabulate[T any](n int, f func(int) T) Seq[T] {
	return seqTabulate(n, f, false)
}

func ArrayFill[T any](n int, e T) Seq[T] {
	return seqFill(n, e, false)
}
