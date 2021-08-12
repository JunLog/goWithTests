package int

import (
	"fmt"
	"testing"
)

// 测试函数命名以 Test 开头
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// 示例，反映代码的实际功能 函数名 Example 开头，
// 会有一行 Output 的注释行,删除这行的话这个函数不会执行
// 这个函数也会出现在 godoc 中
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6

}

// 命令 ： go doc
// 查看包的所有函数

// 命令(未调通) ： godoc -http:=6060 -play
// 从 1.13 版本后被移除，安装 go get golang.org/x/tools/cmd/godoc

// 命令 go test -v
// -v 输出测试详细信息
