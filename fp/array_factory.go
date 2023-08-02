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
	if n <= 0 || step < 0 {
		return emptyArray[int]()
	}

	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, from+step*i)
	}
	return ArrayOfGoSlice(result)
}

func ArrayRange(from int, n int) Seq[int] {
	return ArrayRangeStep(from, n, 1)
}

func ArrayTabulate[T any](n int, f func(int) T) Seq[T] {
	result := []T{}
	for i := 0; i < n; i++ {
		result = append(result, f(i))
	}
	return ArrayOfGoSlice(result)
}

func ArrayFill[T any](n int, e T) Seq[T] {
	return ArrayTabulate(n, func(i int) T { return e })
}
