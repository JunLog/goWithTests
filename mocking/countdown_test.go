package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &spySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		// 反引号 `` 字符可跨行
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("调用 sleeper 的次数不足，预期是 4 只调用了 %d", spySleeper.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("want calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

// 使用 注入 监视 sleep 调用多少次

// 监视 sleep 调用次数
type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

// ------------
// 单独只统计 sleep 数量的话不能体现 sleep 的顺序
type CountdownOperationsSpy struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

// 这里实现了 Sleeper.Sleep()
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// 这里实现了 io.Writer()
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
