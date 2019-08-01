package main

import "fmt"

func main() {


	s := [] int {1,2,3}

	printSlice(s)

	var numbers = make([] int ,3,5)
	printSlice(numbers)


	var num [] int

	printSlice(num)

	if (num == nil) {
		fmt.Println("slice is null")
	}



}



func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}