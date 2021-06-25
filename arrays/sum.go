package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// 使用 range 语法 重构
	// range 迭代返回索引和值，用 _ (空白标识符) 忽略索引
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)

	// for i, number := range numbersToSum {
	// 	sums[i] = Sum(number)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// slice 切片有 容量 和 长度 两个属性
// 长度 len(slice)
// 容量 cap(slice)

// 切片 底层是 数组arrays，就是说 其实切片是内存中的数组的子集
// 切片长度是指切片包含的元素的个数
// 容量是指在切片的第一个元素 在 底层数组的位置到最后一个的元素个数
// s := [9]int{1,2,3,4,5,6,7,8,9} 长度9
// s1 := s[0:5] 长度5，容量5
// s2 := s[5:] 长度4，容量4

// 修改 切片 元素等于修改 数组 元素

// 在没超过底层数组时 append ，就是改变对应位置的值
// 在超过底层数组时，append 就是 新指向一个新的 arrays

// 扩容
// 在切片容量小于 1024 扩容是乘2，大于 1024 扩容是乘 1.25
