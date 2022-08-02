package Select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func TestRacer(t *testing.T)  {
//	slowURL := "http://www.zhihu.com"
//	fastURL := "http://baidu.com"
//
//	want := fastURL
//	got := Racer(slowURL, fastURL)
//
//	if got != want{
//		t.Errorf("got %s, want %s", got, want)
//	}
//
//}

func TestRacer2(t *testing.T)  {
	t.Run("common test", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("want %v, got %v", got ,want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(21 * time.Second)
		serverB := makeDelayedServer(22 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfiguableRacer(serverA.URL, serverB.URL, 20 * time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server  {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}