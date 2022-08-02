package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}
type ConfigurableSleeper struct {
	duraion time.Duration
}

func (o *ConfigurableSleeper) Sleep()  {
	time.Sleep(o.duraion)
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}


const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main()  {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)

}