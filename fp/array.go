package fp

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

type Array[T any] struct {
	data []T
}

func (this Array[T]) Add(e T) Seq[T] {
	return ArrayOfGoSlice(append([]T{e}, this.data...))
}

func (this Array[T]) Get(i int) T {
	if this.IsValidIndex(i) {
		return this.data[i]
	}
	panic("Index " + fmt.Sprint(i) + " out of bounds " + fmt.Sprint(this.Size()))
}

func (this Array[T]) IsEmpty() bool {
	return this.Size() == 0
}

func (this Array[T]) NonEmpty() bool {
	return !this.IsEmpty()
}

func (this Array[T]) HeadOption() Option[T] {
	if this.NonEmpty() {
		return SomeOf(this.data[0])
	}
	return None[T]()
}

func (this Array[T]) LastOption() Option[T] {
	if this.NonEmpty() {
		return SomeOf(this.data[this.Size()-1])
	}
	return None[T]()
}

func (this Array[T]) Tail() Seq[T] {
	if this.NonEmpty() {
		return ArrayOfGoSlice(this.data[1:this.Size()])
	}
	return emptyArray[T]()
}

func (this Array[T]) Equals(that Seq[T]) bool {
	return reflect.DeepEqual(this, that)
}

func (this Array[T]) Reverse() Seq[T] {
	return SeqFoldLeft[T, Seq[T]](this, add[T], emptyArray[T]())
}

func (this Array[T]) Append(e T) Seq[T] {
	return ArrayOfGoSlice(append(this.data, e))
}

func (this Array[T]) Concat(that Seq[T]) Seq[T] {
	return ArrayOfGoSlice(append(this.data, that.ToArray().data...))
}

func (this Array[T]) ContainsElement(e T) bool {
	for i := 0; i < this.Size(); i++ {
		if reflect.DeepEqual(e, this.data[i]) {
			return true
		}
	}
	return false
}

func (this Array[T]) Size() int {
	return len(this.data)
}

func (this Array[T]) Exists(p func(T) bool) bool {
	for i := 0; i < this.Size(); i++ {
		if p(this.data[i]) {
			return true
		}
	}
	return false
}

func (this Array[T]) Filter(p func(T) bool) Seq[T] {
	return collect[T](
		this,
		func(_ int, e T, _ Seq[T]) bool { return p(e) },
		emptyArray[T],
	)
}

func (this Array[T]) FilterNot(p func(T) bool) Seq[T] {
	f := func(e T) bool { return !p(e) }
	return this.Filter(f)
}

func (this Array[T]) Find(p func(T) bool) Option[T] {
	for i := 0; i < this.Size(); i++ {
		if p(this.data[i]) {
			return SomeOf(this.data[i])
		}
	}
	return None[T]()
}

func (this Array[T]) Diff(that Seq[T]) Seq[T] {
	f := func(e T) bool { return that.ContainsElement(e) }
	return this.FilterNot(f)
}

func (this Array[T]) Distinct() Seq[T] {
	return collect[T](
		this,
		func(_ int, e T, acc Seq[T]) bool { return !acc.ContainsElement(e) },
		emptyArray[T],
	)
}

func (this Array[T]) Drop(n int) Seq[T] {
	if n <= 0 {
		return this
	}
	index := minInt(n, this.Size())
	return ArrayOfGoSlice(this.data[index:])
}

func (this Array[T]) DropRight(n int) Seq[T] {
	if n <= 0 {
		return this
	}
	index := this.Size() - minInt(n, this.Size())
	return ArrayOfGoSlice(this.data[:index])
}

func (this Array[T]) DropWhile(p func(T) bool) Seq[T] {
	index := this.IndexOfWhere(func(e T) bool { return !p(e) })
	if index < 0 {
		return emptyArray[T]()
	}
	return ArrayOfGoSlice(this.data[index:])
}

func (this Array[T]) Take(n int) Seq[T] {
	if n <= 0 {
		return emptyArray[T]()
	}
	index := minInt(n, this.Size())
	return ArrayOfGoSlice(this.data[:index])
}

func (this Array[T]) TakeRight(n int) Seq[T] {
	if n <= 0 {
		return emptyArray[T]()
	}
	index := this.Size() - minInt(n, this.Size())
	return ArrayOfGoSlice(this.data[index:])
}

func (this Array[T]) TakeWhile(p func(T) bool) Seq[T] {
	index := -1
	for i := 0; i < this.Size(); i++ {
		if p(this.data[i]) {
			index = i
		} else {
			break
		}
	}

	if index < 0 {
		return emptyArray[T]()
	}

	return ArrayOfGoSlice(this.data[:index+1])
}

func (this Array[T]) ForAll(p func(T) bool) bool {
	for i := 0; i < this.Size(); i++ {
		if !p(this.data[i]) {
			return false
		}
	}
	return true
}

func (this Array[T]) ForAllPar(p func(T) bool) bool {
	var wg sync.WaitGroup
	ch := make(chan bool)
	for i := 0; i < this.Size(); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			ch <- p(this.data[index])
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	acc := true
	for r := range ch {
		acc = acc && r
	}

	return acc
}

func (this Array[T]) ForEach(f func(T) Unit) Unit {
	for i := 0; i < this.Size(); i++ {
		f(this.data[i])
	}
	return GetUnit()
}

func (this Array[T]) ForEachPar(f func(T) Unit) Unit {
	var wg sync.WaitGroup
	for i := 0; i < this.Size(); i++ {
		wg.Add(1)
		go func(index int) Unit {
			defer wg.Done()
			return f(this.data[index])
		}(i)
	}
	wg.Wait()
	return GetUnit()
}

func (this Array[T]) Indexes() Seq[int] {
	return ArrayRange(0, this.Size())
}

func (this Array[T]) IndexOf(e T) int {
	for i := 0; i < this.Size(); i++ {
		if reflect.DeepEqual(e, this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) IndexOfFrom(e T, from int) int {
	for i := from; i < this.Size(); i++ {
		if reflect.DeepEqual(e, this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) IndexOfWhere(p func(T) bool) int {
	for i := 0; i < this.Size(); i++ {
		if p(this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) IndexOfWhereFrom(p func(T) bool, from int) int {
	for i := from; i < this.Size(); i++ {
		if p(this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) LastIndexOf(e T) int {
	for i := this.Size() - 1; i >= 0; i-- {
		if reflect.DeepEqual(e, this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) LastIndexOfFrom(e T, from int) int {
	for i := this.Size() - 1; i >= from; i-- {
		if reflect.DeepEqual(e, this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) LastIndexOfWhere(p func(T) bool) int {
	for i := this.Size() - 1; i >= 0; i-- {
		if p(this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) LastIndexOfWhereFrom(p func(T) bool, from int) int {
	for i := this.Size() - 1; i >= from; i-- {
		if p(this.data[i]) {
			return i
		}
	}
	return -1
}

func (this Array[T]) IsValidIndex(i int) bool {
	return i >= 0 && i < this.Size()
}

func (this Array[T]) Min(less func(T, T) bool) Option[T] {
	if this.IsEmpty() {
		return None[T]()
	}
	return this.Sort(less).HeadOption()
}

func (this Array[T]) Max(less func(T, T) bool) Option[T] {
	if this.IsEmpty() {
		return None[T]()
	}
	return this.Sort(less).LastOption()
}

func (this Array[T]) MkString(sep string) string {
	return mkString[T](this, sep)
}

func (this Array[T]) PrefixLength(p func(T) bool) int {
	return prefixLength[T](this, p)
}

func (this Array[T]) Reduce(f func(T, T) T) Option[T] {
	return reduce[T](this, f)
}

func (this Array[T]) Slice(from int, until int) Seq[T] {
	lo := maxInt(from, 0)
	up := minInt(until, this.Size())
	if up <= lo || this.IsEmpty() {
		return emptyArray[T]()
	}
	return ArrayOfGoSlice(this.data[lo:up])
}

func (this Array[T]) FindSlice(that Seq[T]) Option[int] {
	if this.IsEmpty() || that.IsEmpty() || (this.Size() < that.Size()) {
		return None[int]()
	}

	for i := 0; i < this.Size(); i++ {
		if this.Slice(i, i+that.Size()).Equals(that) {
			return SomeOf(i)
		}
	}
	return None[int]()
}

func (this Array[T]) StartsWith(that Seq[T]) bool {
	return 0 == this.FindSlice(that).GetOrElse(-1)
}

func (this Array[T]) EndsWith(that Seq[T]) bool {
	return this.Reverse().StartsWith(that.Reverse())
}

func (this Array[T]) SplitAt(i int) Pair[Seq[T], Seq[T]] {
	if this.IsEmpty() {
		return PairOf(emptyArray[T](), emptyArray[T]())
	}
	if !this.IsValidIndex(i) {
		return PairOf(emptyList[T]().Concat(this), emptyArray[T]())
	}
	return PairOf(this.Take(i), this.Drop(i))
}

func (this Array[T]) Sort(less func(T, T) bool) Seq[T] {
	s := copyGoSlice(this.ToGoSlice())
	sort.Slice(
		s,
		func(i int, j int) bool { return less(s[i], s[j]) },
	)
	return ArrayOfGoSlice(s)
}

func (this Array[T]) ToList() List[T] {
	return ListOfGoSlice(this.data).ToList()
}

func (this Array[T]) ToArray() Array[T] {
	return this
}

func (this Array[T]) ToGoSlice() []T {
	return this.data
}

func (this Array[T]) IsList() bool {
	return false
}
