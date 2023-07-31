package fp_test

import (
	"fmt"
	"testing"

	"github.com/ivan-yankov/gofp/fp"

	"github.com/stretchr/testify/assert"
)

func TestIsDefined(t *testing.T) {
	assert.True(t, fp.SomeOf(1).IsDefined())
	assert.False(t, fp.None[int]().IsDefined())
}

func TestNonDefined(t *testing.T) {
	assert.False(t, fp.SomeOf(1).NonDefined())
	assert.True(t, fp.None[int]().NonDefined())
}

func TestGet(t *testing.T) {
	assert.Equal(t, 1, fp.SomeOf(1).Get())
	assert.Panics(t, func() { fp.None[int]().Get() })
}

func TestGetOrElse(t *testing.T) {
	assert.Equal(t, 1, fp.SomeOf(1).GetOrElse(2))
	assert.Equal(t, 2, fp.None[int]().GetOrElse(2))
}

func TestMapOption_Some(t *testing.T) {
	assert.Equal(
		t,
		fp.SomeOf("1"),
		fp.OptionMap(fp.SomeOf[int](1), func(x int) string { return fmt.Sprint(x) }),
	)
}

func TestMapOption_None(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.OptionMap(fp.None[int](), func(x int) string { return fmt.Sprint(x) }),
	)
}

func TestFlatMapOption_SomeToSome(t *testing.T) {
	assert.Equal(
		t,
		fp.SomeOf("1"),
		fp.OptionFlatMap(fp.SomeOf[int](1), func(x int) fp.Option[string] { return fp.SomeOf(fmt.Sprint(x)) }),
	)
}

func TestFlatMapOption_SomeToNone(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.OptionFlatMap(fp.SomeOf[int](1), func(int) fp.Option[string] { return fp.None[string]() }),
	)
}

func TestFlatMapOption_NoneToNone(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.OptionFlatMap(fp.None[int](), func(x int) fp.Option[string] { return fp.SomeOf(fmt.Sprint(x)) }),
	)
}
