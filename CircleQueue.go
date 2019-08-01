package main

import "errors"

type CircleQueue struct {
	MaxSize int
	array [5]int
	head int
	tail int
}

func (this *CircleQueue) AddQueue(val int) (err error){
	if (this.tail + 1)%this.MaxSize == this.head {
		return errors.New("full")
	}

	this.array[this.tail] = val
	this.tail ++
	return
}

func (this *CircleQueue) GetQueue(val int) (err error) {
	if this.head == this.tail {
		return errors.New("null")
	}

	for i := this.head; (i+this.MaxSize)%this.MaxSize < this.tail; i++ {

	}
}




func main() {
	
}
