package  main
import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T)  {
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
}
