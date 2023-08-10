# gofp

Functional Programming Library for Go

## About

This is a functional programming library for `go`.

The purpose is to propose more features than many of the existing functional programming `go` libraries.

Many of the features are inspired by Scala.

## Limitations

Some of the functions are implemented not as interface methods due to some `go` language limitations:

- Methods cannot define their own generic types, for example for `map` function

- In some cases `go` compiler gives `Instantiation cycle` error, for example method `Seq.Sliding` should return `Seq[Seq[T]]`, which causes the error

Methods that cause `Instantiation cycle` error are implemented as functions.

## Future work

- Parallel sequence operations
    - map
    - flatMap

- Sequence operations
    - GroupBy

Group the sequence to `Seq[Pair[K, Seq[T]]]` based on a given function for creating keys.

- Structures
    - Map
    - LazyList
    - LazyArray
