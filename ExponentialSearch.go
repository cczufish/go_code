package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello,world"[:5]
	s2 := "hello,world"[7:]

	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(s1)
	fmt.Println(s2)


	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5


}


func goType() {
	var a [3]int // 定义一个数组变量的最基本的方式，数组的长度明确指定，数组中的每个元素都以零值初始化

	fmt.Println(a)

	var b = [...]int{1,2,3} // 定义数组，可以在定义的时候顺序指定全部元素的初始化值，数组的长度根据初始化元素的数目自动计算

	fmt.Println(b)

	var c = [...]int{1:2,3:2}

	fmt.Println(c)

	var d = [...]int{1,2,4:5,6}

	fmt.Println(d)

	var e = &b
	fmt.Println(e)

	fmt.Println(b[0],b[1])

	fmt.Println(e[0],e[1])


	for i,v := range e {
		fmt.Println(i,v)
	}


	for i := range b {
		fmt.Printf("a[%d]: %d\n", i, b[i])
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}



}

func ExponentialSearch(array []int,number int) int {
	boundValue := 1
	for boundValue < len(array) && array[boundValue] < number {
		boundValue *= 2
	}
	if boundValue > len(array) {
		boundValue = len(array) - 1
	}
	return BinarySearch(array, boundValue+1, number)

}
func BinarySearch(array []int,bound,number int) int {
	minIndex := 0
	maxIndex := bound - 1
	for minIndex <= maxIndex {
		midIndex := int( (minIndex + maxIndex) / 2)
		midItem := array[midIndex]
		if number == midItem{
			return midIndex
		}
		if midItem < number{
			minIndex = minIndex + 1
		}else if midItem > number {
			maxIndex = midItem - 1
		}
	}
	return -1
}
