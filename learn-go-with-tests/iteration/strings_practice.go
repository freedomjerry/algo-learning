package iteration

import (
	"strings"
)

func stringsClone() string {
	str := "Hello, World"
	copy := strings.Clone(str)
	return copy
}
func stringsCotains() bool {
	str:= "Hello, World"
	sub := "Hello"
	return strings.Contains(str, sub)

}
func stringsTrimLeft() string {
	return strings.TrimLeft("¡¡!!Hello, Gophers!!!", "!¡")
}