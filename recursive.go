package main

import (
	"fmt"
	"errors"
)

func fib(n int,m int) int {

	
}

type error interface {
	Error() string
}


func sqrt(f float64) (float64,error){
	if f < 0{
		return 0,errors.New("math:square root of negative number")
	}

}






func main() {
	//n := 100
	//fmt.Println(fib(n))

	result,err = sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}

}


