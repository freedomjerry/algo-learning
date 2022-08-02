package Select

import (
	"fmt"
	"net/http"
	"time"
)
var tenSecondTimeOut = 10 * time.Second
//func Racer(a, b string) (winner string) {
//	aDuration := measureReponseTime(a)
//	bDuration := measureReponseTime(b)
//	if aDuration < bDuration {
//		return a
//	}
//	return b
//}
//func measureReponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}
func Racer(a, b string) (winner string, err error) {
	return ConfiguableRacer(a, b, tenSecondTimeOut)
}
func ConfiguableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting or %s and %s", a, b)
	}
}
func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}