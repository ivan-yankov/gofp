package fp

type Seq[T any] interface {
	IsEmpty() bool
	NonEmpty() bool
	HeadOption() Option[T]
	LastOption() Option[T]
	Tail() Seq[T]
	Add(T) Seq[T]
	// Append(T) Seq[T]
	// Contains(T) bool
	// ContainsSlice(Seq[T]) bool
	// Count() int
	// Diff(Seq[T]) Seq[T]
	// Distinct() Seq[T]
	// Drop() Seq[T]
	// DropRight() Seq[T]
	// DropWhile(func(T) bool) Seq[T]
	// StartsWith(Seq[T]) bool
	// EndsWith(Seq[T]) bool
	// Equals(Seq[T]) bool
	// Exists(T) bool
	// Filter(func(T) bool) Seq[T]
	// FilterNot(func(T) bool) Seq[T]
	// Find(func(T) bool) Option[T]
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
}
