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
type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountdownOperationSpy) Write(p []byte)(n int, err error)  {
	s.Calls = append(s.Calls, write)
	return
}
func (o *ConfigurableSleeper) Sleep()  {
	time.Sleep(o.duraion)
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper){
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
	//for i := countdownStart; i > 0; i-- {
	//	sleeper.Sleep()
	//}
	//
	//for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	//}
	//
	//sleeper.Sleep()
	//fmt.Fprint(out, finalWord)
}

func main()  {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)

}