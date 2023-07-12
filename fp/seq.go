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
	// Diff(Seq[T]) Seq[T] // all elements in this which are not in that (including duplicates)
	// Distinct() Seq[T] // check if already added to acc during iteration
	// Drop(int) Seq[T]
	// DropRight(int) Seq[T]
	// DropWhile(func(T) bool) Seq[T]
	// ForAll(func(T) bool) bool
	// ForEach(func(T))
	// IndexOf(T) int
	// IndexOfFrom(T, int) int
	// IndexOfWhere(func(T) bool) int
	// IndexOfWhereFrom(func(T) bool, int) int
	// LastIndexOf(T) int
	// LastIndexOfFrom(T, int) int
	// LastIndexOfWhere(func(T) bool) int
	// LastIndexOfWhereFrom(func(T) bool, int) int
	// IsValidIndex(int) bool
	// Indexes() Seq[int]
	// ContainsSlice(Seq[T]) bool
	// StartsWith(Seq[T]) bool
	// EndsWith(Seq[T]) bool
}
