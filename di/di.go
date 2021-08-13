package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// func Greet(writer *bytes.Buffer, name string) {
// 	// fmt.Printf("Hello,%s", name)
// 	// 一个普通的 打印，捕获输出进行测试 很困难
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// 依赖注入
// 就是 原来函数需要一堆的变量参数，变成 实例参数 传入
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")

	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		fmt.Println(err)
	}
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
