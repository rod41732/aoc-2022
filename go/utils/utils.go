package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string)string {
  var bytes, _ = os.ReadFile(fileName)
  return string(bytes)
}

func Split(text, sep string) []string {
  if (sep == "") { panic("sep cannot be empty!")}
  var parts = make([]string, 0)
  for {
    before, after, found  := strings.Cut(text, sep)
    parts = append(parts, before)
    if !found { break }
    text = after
  }
  return parts
}

func ReadLines(fileName string)[]string {
  var bytes, _ = os.ReadFile(fileName)
  return Split(string(bytes), "\n")
}

func Unwrap[T any](val T, err error) T {
  if err != nil { panic(err)}
  return val
}

func Reverse[T any](slice []T) {
  var n = len(slice)
  for i := 0; i < n/2; i++ {
    slice[n-i-1], slice[i] = slice[i], slice[n-i-1]
  }
}


func AssertEqual[T comparable](a, b T) {
  if a == b {return}
  if a != b {
    log.Fatalf("Assertion failed, %v != %v", a, b)
  }
}
