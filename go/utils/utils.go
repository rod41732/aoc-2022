package utils
import "os"
import "log"

func ReadFile(fileName string)string {
  var bytes, _ = os.ReadFile(fileName)
  return string(bytes)
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
