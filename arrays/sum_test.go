package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("获取 %d, 想得到 %d, 输入的数组 %v", got, want, numbers)
			// %v 默认输出格式，结构体
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("得到 %d 预期 %d ，输入 %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// 切片的比较不能使用 等号 直接比较
	// 使用 reflect.DeepEqual 可以比较两个变量是否相等
	// 但 reflect.DeepEqual 这个不是 类型安全 的，将 string 类型 和 slice 类型进行比较也可以通过编译
	if !reflect.DeepEqual(got, want) {
		t.Errorf("获取 %v,预期 %v", got, want)
	}
}

func TestAllTails(t *testing.T) {

	// 这部分代码重复里，提取复用
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("得到 %v , 预期 %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

// 测试覆盖率
// go test -cover

// vscode 快捷键
// 折叠所有区域代码	Ctrl+k,Ctrl+0
// 展开所有区域代码	Ctrl+k,Ctrl+j
// 折叠本级代码	Ctrl+k,Ctrl+[
// 展开本级代码	Ctrl+k,Ctrl+]
