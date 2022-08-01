package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected : '%s', but got '%s'", expected, repeat)
	}

}
func ExampleRepeat() {
	repeat := Repeat("b")
	fmt.Println(repeat)
	// Output: bbbbb
}
func ExampleClone() {
	clone:= stringsClone()
	fmt.Println(clone)
	// Output: Hello, World
}
func ExampleContain() {
	iscontain := stringsCotains()
	fmt.Println(iscontain)
	// Output: true

}
func ExampleTrimLeft()  {
	res := stringsTrimLeft()
	fmt.Println(res)
	// Output: Hello, Gophers!!!
}
func BenchmarkRepeat(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		Repeat("a")
	}
}
func BenchmarkContains(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		stringsCotains()
	}

}