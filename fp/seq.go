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
	ToList() List[T]
	ToGoSlice() []T
}
