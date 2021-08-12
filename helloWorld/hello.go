package helloWorld

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("world", ""))
}

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const frenchHelloPrefix = "Bonjour,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// 函数名称 小写开头 ，私有函数
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default: //如果没匹配到的话就执行本分支
		prefix = englishHelloPrefix
	}
	return // 直接return，不用指定变量，函数声明时已经定义了返回变量
}
