package fp

type Seq[T any] interface {
	Add(T) Seq[T]
	Get(int) T
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
	ForEachPar(func(T) Unit) Unit
	Indexes() Seq[int]
	IndexOf(T) int
	IndexOfFrom(T, int) int
	IndexOfWhere(func(T) bool) int
	IndexOfWhereFrom(func(T) bool, int) int
	LastIndexOf(T) int
	LastIndexOfFrom(T, int) int
	LastIndexOfWhere(func(T) bool) int
	LastIndexOfWhereFrom(func(T) bool, int) int
	IsValidIndex(int) bool
	Min(func(T, T) bool) Option[T]
	Max(func(T, T) bool) Option[T]
	MkString(string) string
	PrefixLength(func(T) bool) int
	Reduce(func(T, T) T) Option[T]
	Slice(int, int) Seq[T]
	SplitAt(int) Pair[Seq[T], Seq[T]]
	Sort(func(T, T) bool) Seq[T]
	ToList() List[T]
	ToArray() Array[T]
	ToGoSlice() []T
	IsList() bool
}

func SeqZip[A, B any](a Seq[A], b Seq[B]) Seq[Pair[A, B]] {
	type T = Seq[Pair[A, B]]

	var it func(Seq[A], Seq[B], T) T
	it = func(sa Seq[A], sb Seq[B], acc T) T {
		if sa.IsEmpty() || sb.IsEmpty() {
			return acc
		}
		return it(sa.Tail(), sb.Tail(), acc.Append(PairOf(sa.HeadOption().Get(), sb.HeadOption().Get())))
	}

	return it(a, b, emptySeq[Pair[A, B]](a.IsList()))
}

func SeqZipWithIndex[T any](seq Seq[T]) Seq[Pair[T, int]] {
	return SeqZip[T, int](seq, seq.Indexes())
}

func SeqFoldLeft[A, B any](seq Seq[A], f func(A, B) B, acc B) B {
	if seq.IsEmpty() {
		return acc
	}

	return SeqFoldLeft(seq.Tail(), f, f(seq.HeadOption().Get(), acc))
}

func SeqFoldRight[A, B any](seq Seq[A], f func(A, B) B, acc B) B {
	return SeqFoldLeft(seq.Reverse(), f, acc)
}

func SeqFoldCount[A, B any](seq Seq[A], f func(int, A, B) B, acc B) B {
	var fold func(seq Seq[A], f func(int, A, B) B, acc B, i int) B
	fold = func(seq Seq[A], f func(int, A, B) B, acc B, i int) B {
		if seq.IsEmpty() {
			return acc
		}
		return fold(seq.Tail(), f, f(i, seq.HeadOption().Get(), acc), i+1)
	}

	return fold(seq, f, acc, 0)
}

func SeqMap[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	return SeqReverseMap(seq, f).Reverse()
}

func SeqReverseMap[A, B any](seq Seq[A], f func(A) B) Seq[B] {
	var it func(Seq[A], Seq[B]) Seq[B]
	it = func(s Seq[A], acc Seq[B]) Seq[B] {
		if s.IsEmpty() {
			return acc
		}
		return it(s.Tail(), acc.Add(f(s.HeadOption().Get())))
	}
	return it(seq, emptySeq[B](seq.IsList()))
}

func SeqFlatMap[A, B any](seq Seq[A], f func(A) Seq[B]) Seq[B] {
	var it func(Seq[A], Seq[B]) Seq[B]
	it = func(s Seq[A], acc Seq[B]) Seq[B] {
		if s.IsEmpty() {
			return acc
		}
		return it(s.Tail(), acc.Concat(f(s.HeadOption().Get())))
	}
	return it(seq, emptySeq[B](seq.IsList()))
}

func SeqSliding[T any](seq Seq[T], size int, step int) Seq[Seq[T]] {
	if size <= 0 || step <= 0 {
		return emptySeq[Seq[T]](seq.IsList())
	}

	var it func(Seq[T], Seq[Seq[T]]) Seq[Seq[T]]
	it = func(s Seq[T], acc Seq[Seq[T]]) Seq[Seq[T]] {
		if s.Size() <= size {
			if s.NonEmpty() {
				return acc.Add(s.Take(size))
			}
			return acc
		}
		return it(s.Drop(step), acc.Add(s.Take(size)))
	}
	return it(seq, emptySeq[Seq[T]](seq.IsList())).Reverse()
}

func SeqStartsWith[T any](a Seq[T], b Seq[T]) bool {
	return 0 == SeqFindSlice(a, b).GetOrElse(-1)
}

func SeqEndsWith[T any](a Seq[T], b Seq[T]) bool {
	return SeqStartsWith(a.Reverse(), b.Reverse())
}

func SeqFindSlice[T any](a Seq[T], b Seq[T]) Option[int] {
	if a.IsEmpty() || b.IsEmpty() || (a.Size() < b.Size()) {
		return None[int]()
	}

	found := SeqSliding(SeqZipWithIndex(a), b.Size(), 1).Find(
		func(x Seq[Pair[T, int]]) bool {
			return SeqMap(
				x,
				func(y Pair[T, int]) T { return y.GetA() },
			).Equals(b)
		},
	)

	return OptionMap(
		found,
		func(x Seq[Pair[T, int]]) int {
			return x.HeadOption().Get().GetB()
		},
	)
}
