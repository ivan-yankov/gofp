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
	IndexOf(T) int
	IndexOfFrom(T, int) int
	IndexOfWhere(func(T) bool) int
	IndexOfWhereFrom(func(T) bool, int) int
	// LastIndexOf(T) int
	// LastIndexOfFrom(T, int) int
	// LastIndexOfWhere(func(T) bool) int
	// LastIndexOfWhereFrom(func(T) bool, int) int
	// IsValidIndex(int) bool
	// ContainsSlice(Seq[T]) bool
	// StartsWith(Seq[T]) bool
	// EndsWith(Seq[T]) bool
}

func add[T any](e T, acc Seq[T]) Seq[T] {
	return acc.Add(e)
}

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
