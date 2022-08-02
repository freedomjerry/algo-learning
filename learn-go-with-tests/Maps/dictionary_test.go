package Maps

import (
	"testing"
)


func TestSearch(t *testing.T)  {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("Known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)

	})
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper() //标记为辅助函数，错误提醒不会是辅助函数内部

	if got != want {
		t.Errorf("got %s, want %s .", got, want)
	}
}
func assertError(t *testing.T, got, want error)  {
	t.Helper()

	if got != want {
		t.Errorf("got error %s, want error %s" , got, want)
	}

}

func TestAdd(t *testing.T)  {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string)  {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word :", err)
	}
	if definition != got {
		t.Errorf("got %s, want %s", got, definition)
	}


}

func TestUpdate(t *testing.T)  {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dicitonary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dicitonary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dicitonary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T)  {
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %s, to be deleted", word)
	}

}