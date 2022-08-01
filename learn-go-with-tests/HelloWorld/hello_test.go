package HelloWorld

import "testing"

func TestHello(t *testing.T) {//测试，清楚说明代码需要什么

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q', want '%q'", got ,want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "") // 加入指定问候的接收者
		want := "Hello, Chris"

		//if got != want {
		//	t.Errorf("got '%q' want '%q'", got, want)
		//}
		assertCorrectMessage(t, got ,want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		//if got != want {
		//	t.Errorf("got '%q', want '%q'", got, want)
		//}
		assertCorrectMessage(t, got ,want) // 断言重构为函数
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got ,want)

	})
}