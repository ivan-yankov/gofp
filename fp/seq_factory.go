package fp

func seqRangeStep(from int, n int, step int, list bool) Seq[int] {
	if n <= 0 || step < 0 {
		return emptySeq[int](list)
	}

	inc := func(x int) int { return x + 1 }
	p := func(x int) bool { return x == n }
	f := func(x int, acc Seq[int]) Seq[int] { return acc.Add(from + x*step) }
	return loop(0, inc, p, f, emptySeq[int](list)).Reverse()
}

func seqRange(from int, n int, list bool) Seq[int] {
	return seqRangeStep(from, n, 1, list)
}

func seqTabulate[T any](n int, f func(int) T, list bool) Seq[T] {
	indexes := seqRange(0, n, list)
	fAcc := func(i int, acc Seq[T]) Seq[T] {
		return acc.Add(f(i))
	}

	return SeqFoldLeft[int, Seq[T]](indexes, fAcc, emptySeq[T](list)).Reverse()
}

func seqFill[T any](n int, e T, list bool) Seq[T] {
	return seqTabulate(n, func(i int) T { return e }, list)
}
