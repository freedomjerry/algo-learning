package array_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {
	//
	//t.Run("collection of 5 numbers", func(t *testing.T) {
	//	nums := []int{1, 2, 3, 4, 5}
	//	got := Sum(nums)
	//	want := 15
	//
	//	if want != got {
	//		t.Errorf("got %d, want %d, given, %v", got, want, nums)
	//	}
	//})
	t.Run("collection of any size ", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got !=want {
			t.Errorf("got %d, want %d, given %v", got, want ,nums)
		}

	})
}
func TestSumAll(t *testing.T)  {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}

	//if got != want {
	if !reflect.DeepEqual(got, want){
		t.Errorf("want %v want %v", got, want)
	}
}
func TestSumTail(t *testing.T)  {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("sfely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSums(t, got, want)

	})
}