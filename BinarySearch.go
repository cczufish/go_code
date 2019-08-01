package main

import "fmt"

func main() {

	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(BinarySearch(number,2))

}

// 二分法检索（binary search）又称折半检索，二分法检索的基本思想是设字典中的元素从小到大有序地存放在数组（array）中。

/*

首先将给定值key与字典中间位置上元素的关键码(key)比较，如果相等，则检索成功；
否则，若key小，则在字典前半部分中继续进行二分法检索;
若key大，则在字典后半部分中继续进行二分法检索。
这样，经过一次比较就缩小一半的检索区间，如此进行下去，直到检索成功或检索失败。
偶数个取中间2个其中任何一个作为中间元素
二分法检索是一种效率较高的检索方法，要求字典在顺序表中按关键码排序。

*/
func BinarySearch(array []int,number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
		midItem := array[midIndex]
		if number == midItem {
			return midIndex
		}
		if midItem < number {
			midIndex = midIndex + 1
		}else if midItem > number {
			maxIndex = midIndex -1
		}

	}
	return -1
}

