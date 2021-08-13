package mocking

import (
	"fmt"
	"io"

	// "os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	// sleep
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

const finalWord = "Go!"
const CountdownStart = 3

// bytes.Buffer 可以看作是一个 数据的容器
// func Countdown(out *bytes.Buffer) {
// 使用 接口 替代
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		// time.Sleep(1 * time.Second)
		// 依赖注入
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	// time.Sleep(1 * time.Second)
	// 依赖注入
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// func main() {
// 	// Countdown(os.Stdout)
// 	sleeper := &ConfigurableSleeper{1 * time.Second}
// 	Countdown(os.Stdout, sleeper)
// }
