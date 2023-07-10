package fp

import "errors"

type Option[T any] interface {
	IsDefined() bool
	NonDefined() bool
	Get() (T, error)
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

func (x option[T]) IsDefined() bool {
	return x.defined
}

func (x option[T]) NonDefined() bool {
	return !x.defined
}

func (x option[T]) Get() (T, error) {
	if x.defined {
		return x.value, nil
	} else {
		var x T
		return x, errors.New("Unable to get value from None")
	}
}

func (x option[T]) GetOrElse(y T) T {
	if x.defined {
		v, _ := x.Get()
		return v
	} else {
		return y
	}
}

func MapOption[A, B any](x Option[A], f func(A) B) Option[B] {
	if x.IsDefined() {
		v, _ := x.Get()
		return SomeOf(f(v))
	} else {
		return None[B]()
	}
}

func FlatMapOption[A, B any](x Option[A], f func(A) Option[B]) Option[B] {
	if x.IsDefined() {
		v, _ := x.Get()
		return f(v)
	} else {
		return None[B]()
	}
}
