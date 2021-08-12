package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	// t.Run("键 存在", func(t *testing.T) {
	// 	got, _ := dictionary.Search("test")
	// 	want := "this is just a test"

	// 	assertString(t, got, want)
	// })

	// t.Run("键 不存在", func(t *testing.T) {
	// 	_, err := dictionary.Search("unknow")
	// 	want := "找不到您要找的词"

	// 	if err == nil {
	// 		// t.Fatal 被调用，将停止测试
	// 		t.Fatal("预计会出错")
	// 	}
	// 	// Error 我类型使用 .Error() 方法转换成字符串
	// 	assertString(t, err.Error(), want)

	// })

	t.Run("键 不存在", func(t *testing.T) {
		_, got := dictionary.Search("unknow")

		assertError(t, got, ErrNotFound)
	})
}

// func assertString(t *testing.T, got, want string) {
// 	// 声明此函数是 辅助函数，当测试失败时报的行号将是函数中的而不是辅助函数内部的
// 	t.Helper()

// 	if got != want {
// 		t.Errorf("获得 '%s' 预期 '%s'", got, want)
// 	}
// }

func assertString(t *testing.T, got, want string) {
	// 声明此函数是 辅助函数，当测试失败时报的行号将是函数中的而不是辅助函数内部的
	t.Helper()

	if got != want {
		t.Errorf("获得 '%s' 预期 '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func TestAdd(t *testing.T) {
	// dictionary := Dictionary{}
	// dictionary.Add("test", "this is just a test")

	// want := "this is just a test"
	// got, err := dictionary.Search("test")
	// if err != nil {
	// 	t.Fatal("应找到添加的单词：", err)
	// }

	// if want != got {
	// 	t.Errorf("got '%s' want '%s'", got, want)
	// }
	t.Run("新词", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("已存在词", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("应找到添加单词：", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("已存在", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("新词", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoedNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("预期 '%s' 已经删除了", word)
	}

}
