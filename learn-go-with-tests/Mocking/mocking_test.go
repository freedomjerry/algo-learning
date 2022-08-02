package  main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T)  {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spysleeper := &SpySleeper{}

		Countdown(buffer, spysleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s, want %s", got ,want)
		}
		if spysleeper.Calls != 4 {
			t .Errorf("not enough calls to sleeper, want 4 got %d", spysleeper.Calls)
		}
	})
	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}
