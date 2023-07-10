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

func TestGet_Some(t *testing.T) {
	v, e := fp.SomeOf(1).Get()
	assert.Equal(t, 1, v)
	assert.Nil(t, e)
}

func TestGet_None(t *testing.T) {
	_, e := fp.None[int]().Get()
	assert.NotNil(t, e)
	assert.Equal(t, "Unable to get value from None", e.Error())
}

func TestGetOrElse(t *testing.T) {
	assert.Equal(t, 1, fp.SomeOf(1).GetOrElse(2))
	assert.Equal(t, 2, fp.None[int]().GetOrElse(2))
}

func TestMapOption_Some(t *testing.T) {
	assert.Equal(
		t,
		fp.SomeOf("1"),
		fp.MapOption(fp.SomeOf[int](1), func(x int) string { return fmt.Sprint(x) }),
	)
}

func TestMapOption_None(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.MapOption(fp.None[int](), func(x int) string { return fmt.Sprint(x) }),
	)
}

func TestFlatMapOption_SomeToSome(t *testing.T) {
	assert.Equal(
		t,
		fp.SomeOf("1"),
		fp.FlatMapOption(fp.SomeOf[int](1), func(x int) fp.Option[string] { return fp.SomeOf(fmt.Sprint(x)) }),
	)
}

func TestFlatMapOption_SomeToNone(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.FlatMapOption(fp.SomeOf[int](1), func(int) fp.Option[string] { return fp.None[string]() }),
	)
}

func TestFlatMapOption_NoneToNone(t *testing.T) {
	assert.Equal(
		t,
		fp.None[string](),
		fp.FlatMapOption(fp.None[int](), func(x int) fp.Option[string] { return fp.SomeOf(fmt.Sprint(x)) }),
	)
}
