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

func ListOf[T any](elements ...T) Seq[T] {
	i := len(elements) - 1
	inc := func(x int) int { return x - 1 }
	p := func(x int) bool { return x < 0 }
	f := func(x int, acc Seq[T]) Seq[T] { return acc.Add(elements[x]) }
	return loop(i, inc, p, f, emptyList[T]())
}

func ListRangeStep(from int, n int, step int) Seq[int] {
	if n <= 0 || step < 0 {
		return emptyList[int]()
	}

	inc := func(x int) int { return x + 1 }
	p := func(x int) bool { return x == n }
	f := func(x int, acc Seq[int]) Seq[int] { return acc.Add(from + x*step) }
	return loop(0, inc, p, f, emptyList[int]()).Reverse()
}

func ListRange(from int, n int) Seq[int] {
	return ListRangeStep(from, n, 1)
}

func ListTabulate[T any](n int, f func(int) T) Seq[T] {
	indexes := ListRange(0, n)
	fAcc := func(i int, acc Seq[T]) Seq[T] {
		return acc.Add(f(i))
	}

	return ListFoldLeft[int, Seq[T]](indexes, fAcc, emptyList[T]()).Reverse()
}

func ListFill[T any](n int, e T) Seq[T] {
	return ListTabulate(n, func(i int) T { return e })
}
