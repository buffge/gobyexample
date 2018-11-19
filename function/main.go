package main

import "fmt"

func main() {
	fmt.Println(sum0(2, 34))
	fmt.Println(sum1(2, 34))
	fmt.Println(sum1(2, 34, 0x78, 0x12))
	fmt.Println(sum1(2, 34, 34, 234, 234, 234))
	numArr := []int{1, 2, 3, 4}
	//... 结构数组
	fmt.Println(sum1(numArr...))
	//匿名函数
	total := sum2()
	fmt.Println(total(1))
	fmt.Println(total(2))
	fmt.Println(total(3))
	fmt.Println(total(4))

}

/**
变量参数,返回值
*/
func sum0(a, b int) (res int) {
	res = a + b
	return
}

/**
可变参数
*/
func sum1(numArr ...int) (total int) {
	for _, num := range numArr {
		total += num
	}
	return
}

//匿名函数

func sum2() func(b int) int {
	a := 0
	return func(b int) int {
		a = a + b
		return a
	}
}
