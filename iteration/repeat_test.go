package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("预期 '%q' 但得到 '%q'", expected, repeated)
	}
}

// 基准测试（benchmarks）
// Benchmark 开头，参数为 *testing.B(可以访问隐形命名 b.N)
// 基准测试会运行 b.N 次
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// 测试命令
// go test -bench=.
// Windows Powershell 环境下 go test -bench="."

// 基准测试 benchmark
// 基准测试是一种测试代码性能的方法
// 基准测试 不能 有 返回值
// 可以在 for 循环前加上 b.ResetTimer() ，重置计时器，避免前面的代码耗时干扰
// 测试的代码放在 for 循环里面

// 输出内容
// goos: linux
// goarch: amd64
// pkg: iteration
// cpu: Intel(R) Xeon(R) CPU E5-4620 0 @ 2.20GHz
// BenchmarkRepeat-16       4050237               286.8 ns/op
// PASS
// ok      iteration       1.473s


// benchmark名称——CPU数(GOMAXPROCS)  循环次数  平均每次耗时