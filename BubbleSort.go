package main

import "fmt"

/*

冒泡排序（Bubble Sort），是一种计算机科学领域的较简单的排序算法。
它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果他们的顺序（如从大到小、首字母从A到Z）错误就把他们交换过来。走访元素的工作是重复地进行直到没有相邻元素需要交换，也就是说该元素已经排序完成。
*/

func main() {

	var num = []int{3,5,12,0,4}
	fmt.Println(num)
	BubbleSort(num)
	fmt.Println(num)
}

func BubbleSort(array []int){
	swapCount := 1
	for swapCount > 0{
		swapCount = 0
		for itemIndex := 0;itemIndex < len(array)-1;itemIndex ++ {
			if array[itemIndex] > array[itemIndex + 1]{
				array[itemIndex],array[itemIndex+1]  = array[itemIndex + 1],array[itemIndex]
				swapCount += 1
			}
		}
	}
}