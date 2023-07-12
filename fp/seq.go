package fp

type Seq[T any] interface {
	Add(T) Seq[T]
	IsEmpty() bool
	NonEmpty() bool
	HeadOption() Option[T]
	LastOption() Option[T]
	Tail() Seq[T]
	Equals(Seq[T]) bool
	Reverse() Seq[T]
	Append(T) Seq[T]
	Concat(Seq[T]) Seq[T]
	ContainsElement(T) bool
	Size() int
	Exists(func(T) bool) bool
	Filter(func(T) bool) Seq[T]
	FilterNot(func(T) bool) Seq[T]
	Find(func(T) bool) Option[T]
	Diff(Seq[T]) Seq[T]
	Distinct() Seq[T]
	Drop(int) Seq[T]
	DropRight(int) Seq[T]
	DropWhile(func(T) bool) Seq[T]
	Take(int) Seq[T]
	TakeRight(int) Seq[T]
	TakeWhile(func(T) bool) Seq[T]
	ForAll(func(T) bool) bool
	ForEach(func(T) Unit) Unit
	Indexes() Seq[int]
	// ZipWithIndex() Seq[Pair[T, int]]
	// IndexOf(T) int
	// IndexOfFrom(T, int) int
	// IndexOfWhere(func(T) bool) int
	// IndexOfWhereFrom(func(T) bool, int) int
	// LastIndexOf(T) int
	// LastIndexOfFrom(T, int) int
	// LastIndexOfWhere(func(T) bool) int
	// LastIndexOfWhereFrom(func(T) bool, int) int
	// IsValidIndex(int) bool
	// ContainsSlice(Seq[T]) bool
	// StartsWith(Seq[T]) bool
	// EndsWith(Seq[T]) bool
}

func seqOf[T any](emptySeq func() Seq[T], elements []T) Seq[T] {
	i := len(elements) - 1
	inc := func(x int) int { return x - 1 }
	p := func(x int) bool { return x < 0 }
	f := func(x int, acc Seq[T]) Seq[T] { return acc.Add(elements[x]) }
	return loop(i, inc, p, f, emptySeq())
}

func seqRangeStep(emptySeq func() Seq[int], from int, n int, step int) Seq[int] {
	if n <= 0 || step < 0 {
		return emptySeq()
	}

	inc := func(x int) int { return x + 1 }
	p := func(x int) bool { return x == n }
	f := func(x int, acc Seq[int]) Seq[int] { return acc.Add(from + x*step) }
	return loop(0, inc, p, f, emptySeq()).Reverse()
}

func seqRange(emptySeq func() Seq[int], from int, n int) Seq[int] {
	return seqRangeStep(emptySeq, from, n, 1)
}

func seqTabulate[T any](emptySeqInt func() Seq[int], emptySeq func() Seq[T], n int, f func(int) T) Seq[T] {
	indexes := seqRange(emptySeqInt, 0, n)
	fAcc := func(i int, acc Seq[T]) Seq[T] {
		return acc.Add(f(i))
	}

	return iterate[int, Seq[T]](indexes, fAcc, emptySeq()).Reverse()
}

func seqFill[T any](emptySeqInt func() Seq[int], emptySeq func() Seq[T], n int, e T) Seq[T] {
	return seqTabulate(emptySeqInt, emptySeq, n, func(i int) T { return e })
}

func seqZip[A, B any](emptySeq func() Seq[Pair[A, B]], sa Seq[A], sb Seq[B]) Seq[Pair[A, B]] {
	type T = Seq[Pair[A, B]]

	var it func(Seq[A], Seq[B], T) T
	it = func(s1 Seq[A], s2 Seq[B], acc T) T {
		if s1.IsEmpty() || s2.IsEmpty() {
			return acc
		}

		a, _ := s1.HeadOption().Get()
		b, _ := s2.HeadOption().Get()
		return it(s1.Tail(), s2.Tail(), acc.Add(PairOf(a, b)))
	}

	return it(sa, sb, emptySeq()).Reverse()
}
