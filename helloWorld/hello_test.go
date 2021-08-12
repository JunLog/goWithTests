package helloWorld

import (
	"testing"
)

// 测出文件必须是 xxx_test.go 的文件名

// 测试函数命名必须是 Test 开头

// 测试函数必须接收一个 t *testing.T 的参数
// t 变量这是个勾子,测试失败时可以执行 t.Fail() 之类的操作

func TestHello(t *testing.T) {
	// got := Hello("Chris")
	// want := "Hello,Chris"

	// if got != want {
	// 	// t.Errorf 打印一条消息并使测试失败，f 格式化 format
	// 	t.Errorf("got '%q' want '%q'", got, want)
	// 	// %q 带双引号的字符串输出
	// }
	// ================子测试=====================
	// t.Run("saying hello to people", func(t *testing.T) {
	// 	got := Hello("Chris")
	// 	want := "Hello,Chris"

	// 	if got != want {
	// 		t.Errorf("got '%q' want '%q'", got, want)
	// 	}
	// })

	// t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello,World"

	// 	if got != want {
	// 		t.Errorf("got '%q' want '%q'", got, want)
	// 	}
	// })
	// ===================================
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 声明此函数是 辅助函数，当测试失败时报的行号将是函数中的而不是辅助函数内部的
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello,Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello,World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola,Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Lauren", french)
		want := "Bonjour,Lauren"
		assertCorrectMessage(t, got, want)
	})
}
