package main

import (
	"errors"
	"fmt"
)

type QueueArr struct {
	 MaxSize int
	 array [5]int
	 front int
	 rear int
}


func (this *QueueArr) AddQueue(val int) (err error){
	if (this.rear + 1 == this.MaxSize){
		return errors.New("full")
	}

	this.rear ++
	this.array[this.rear] = val
	return
}


func (this *QueueArr) GetQueue(val int) (err error) {
	if this.rear == this.front {
		return errors.New("null")
	}
	for i := this.front + 1;i <= this.rear;i ++ {
		fmt.Printf("queue %d=%v\n",i+1,this.array[i])
	}
	return
}

func (this *QueueArr) GetOneQueue() (val int, err error) {
	if this.rear == this.front {
		return -1,errors.New("null")
	}

	this.front ++
	val = this.array[this.front]
	return val,err

}



func main() {
	queueArr := &QueueArr{
		MaxSize:5,
		front:-1,
		rear:-1,
	}

	err := queueArr.AddQueue(0)
	err = queueArr.AddQueue(1)
	err = queueArr.AddQueue(2)

	if err != nil {
		fmt.Println("err")
	}








}
