package utils_test

import (
	"aoc/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNormal(t *testing.T) {
  res := utils.Split("foo_bar_baz", "_")
  assert.ElementsMatch(t, res, []string{"foo", "bar", "baz"}, "xd")
}

func TestSplitMultichars(t *testing.T) {
  res := utils.Split("fooXDDbarXXDDbaz", "XD")
  assert.ElementsMatch(t, res, []string{"foo", "DbarX", "Dbaz"}, "xd")
}

func TestSplitMultichars2(t *testing.T) {
  res := utils.Split("foo___bar____baz", "__")
  assert.ElementsMatch(t, res, []string{"foo", "_bar", "", "baz"}, "xd")
}

func TestSplitConsecutiveSeps(t *testing.T) {
  res := utils.Split("foo\nbar\n\nbaz\n\n", "\n")
  assert.ElementsMatch(t, res, []string{"foo", "bar", "", "baz", "", ""}, "xd")
}

// func TestSplitWhenNotFound(t *testing.T) {
//   res := utils.Split("foobar", "xddd")
//   assert.ElementsMatch(t, res, []string{"foobar"}, "xd")
// }
//
// func TestSplitWhenSepAtFirst(t *testing.T) {
//   res := utils.Split("XDfoobar", "XD")
//   assert.ElementsMatch(t, res, []string{"", "foobar"}, "xd")
// }
//
// func TestSplitWhenSepAtLast(t *testing.T) {
//   res := utils.Split("foobarXD", "XD")
//   assert.ElementsMatch(t, res, []string{"foobar", ""}, "xd")
// }
