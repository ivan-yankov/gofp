package fp_test

import (
	"fmt"
	"testing"

	"github.com/ivan-yankov/gofp/fp"

	"github.com/stretchr/testify/assert"
)

func TestIsLeft(t *testing.T) {
	x := fp.LeftOf[int, int](1)
	assert.True(t, true, x.IsLeft())
	assert.False(t, false, x.IsRight())
}

func TestLeftIsRight(t *testing.T) {
	x := fp.RightOf[int, int](1)
	assert.False(t, false, x.IsLeft())
	assert.True(t, true, x.IsRight())
}

func TestGetLeft(t *testing.T) {
	assert.Equal(t, 1, fp.LeftOf[int, string](1).GetLeft().GetOrElse(0))
}

func TestGetRight(t *testing.T) {
	assert.Equal(t, "hello", fp.RightOf[int, string]("hello").GetRight().GetOrElse(""))
}

func TestFold_Left(t *testing.T) {
	var left = ""
	var right = 0
	fp.LeftOf[string, int]("1").Fold(
		func(x string) { left = "Error " + x },
		func(x int) { right = right + x },
	)
	assert.Equal(t, "Error 1", left)
	assert.Equal(t, 0, right)
}

func TestFold_Right(t *testing.T) {
	var left = ""
	var right = 1
	fp.RightOf[string, int](10).Fold(
		func(x string) { left = "Error " + x },
		func(x int) { right = right + x },
	)
	assert.Empty(t, left)
	assert.Equal(t, 11, right)
}

func TestSwap_LeftToRight(t *testing.T) {
	e := fp.LeftOf[int, string](1).Swap()
	assert.True(t, e.IsRight())
	assert.Equal(t, 1, e.GetRight().GetOrElse(0))
}

func TestSwap_RightToLeft(t *testing.T) {
	e := fp.RightOf[int, string]("hello").Swap()
	assert.True(t, e.IsLeft())
	assert.Equal(t, "hello", e.GetLeft().GetOrElse(""))
}

func TestMapEither_Right(t *testing.T) {
	e := fp.MapEither[bool, int, string](
		fp.RightOf[bool, int](1),
		func(x int) string { return fmt.Sprint(x * 2) },
	)
	assert.True(t, e.IsRight())
	assert.Equal(t, "2", e.GetRight().GetOrElse(""))
}

func TestMapEither_Left(t *testing.T) {
	e := fp.MapEither[bool, int, string](
		fp.LeftOf[bool, int](true),
		func(x int) string { return fmt.Sprint(x * 2) },
	)
	assert.True(t, e.IsLeft())
	assert.True(t, e.GetLeft().GetOrElse(false))
}

func TestFlatMapEither_Right(t *testing.T) {
	e := fp.FlatMapEither[bool, int, string](
		fp.RightOf[bool, int](1),
		func(x int) fp.Either[bool, string] { return fp.RightOf[bool, string](fmt.Sprint(x * 2)) },
	)
	assert.True(t, e.IsRight())
	assert.Equal(t, "2", e.GetRight().GetOrElse(""))
}

func TestFlatMapEither_Left(t *testing.T) {
	e := fp.FlatMapEither[bool, int, string](
		fp.LeftOf[bool, int](true),
		func(x int) fp.Either[bool, string] { return fp.RightOf[bool, string](fmt.Sprint(x * 2)) },
	)
	assert.True(t, e.IsLeft())
	assert.True(t, e.GetLeft().GetOrElse(false))
}
