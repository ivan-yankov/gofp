package fp

type Option[T any] interface {
	IsDefined() bool
	NonDefined() bool
	Get() T
	GetOrElse(T) T
}

type option[T any] struct {
	value   T
	defined bool
}

func SomeOf[T any](value T) Option[T] {
	return option[T]{
		value:   value,
		defined: true,
	}
}

func None[T any]() Option[T] {
	return option[T]{
		value:   *new(T),
		defined: false,
	}
}

func MapOption[A, B any](x Option[A], f func(A) B) Option[B] {
	if x.IsDefined() {
		return SomeOf(f(x.Get()))
	}

	return None[B]()
}

func FlatMapOption[A, B any](x Option[A], f func(A) Option[B]) Option[B] {
	if x.IsDefined() {
		return f(x.Get())
	}

	return None[B]()
}

func (this option[T]) IsDefined() bool {
	return this.defined
}

func (this option[T]) NonDefined() bool {
	return !this.defined
}

func (this option[T]) Get() T {
	if this.defined {
		return this.value
	}

	panic("Unable to get value from None")
}

func (this option[T]) GetOrElse(y T) T {
	if this.defined {
		return this.Get()
	}

	return y
}
