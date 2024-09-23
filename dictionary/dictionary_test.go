package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("search known word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
	t.Run("search unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		assertError(t, err)
		assertString(t, err.Error(), want, "unknown")
	})
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")
		if err != nil {
			t.Fatal("should add new word")
		}
		want := "this is just a test"
		assertDefinition(t, dictionary, "test", want)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update", func(t *testing.T) {
		dictionary := Dictionary{"test": "test"}
		err := dictionary.Update("test", "update test")
		if err != nil {
			t.Fatal("should update definition")
		}
		assertDefinition(t, dictionary, "test", "update test")
	})
	t.Run("update none existed word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is a test")
		assertError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		err := dictionary.Delete("test")
		if err != nil {
			t.Fatal("should delete word")
		}
		_, err = dictionary.Search("search")
		if err == nil {
			t.Fatal("should delete word")
		}
	})
	t.Run("delete none existed word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
		assertError(t, err)
	})
}
func assertString(t *testing.T, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, given)

	}
}
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected to get error.")
	}
}
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
