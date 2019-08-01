package main

import "testing"

type Person struct {
	Name string
	ID string
	Phone string
}

type Collection struct {
	DB string
	Name string
}

var col = Collection{"testDB","contacts"}
var p = Person{"jack","123456","13261573787"}


func TestInsert(t *testing.T){
	//insert
}

func main() {

	
}
