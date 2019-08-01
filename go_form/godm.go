package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL = "127.0.0.1:20717"
)


var col = Collection{"testDB","contacts"}
var p = Person{"jack","123456","13261573787"}



type Person struct {
	Name string
	ID string
	Phone string
}

type Collection struct {
	DB string
	Name string
}

// 往通讯录中添加一条记录

func insert(p Person,col Collection){

	session,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()
	c := session.DB(col.DB).C(col.Name)

	err = c.Insert(p)
	if err != nil {
		fmt.Println(err)
	}

}

// 根据姓名从通讯录中查找联系方式

func findByName(name string,col Collection) Person {
	session,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic,true)
	collection := session.DB(col.DB).C(col.Name)

	result := Person{}

	err = collection.Find(bson.M{"name":name}).One(&result)
	if err != nil {
		fmt.Println(err)
	}

	return result

}

// 更新通讯录中某人的联系方式

func update(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Update(bson.M{"name": p.Name}, bson.M{"$set": bson.M{"phone": p.Phone}})
	if err != nil {
		fmt.Println(err)
	}
}


// 将某人的信息从通讯录中删除
func deleteData(p Person, col Collection) {
	session, err := mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	collection := session.DB(col.DB).C(col.Name)
	err = collection.Remove(bson.M{"name": p.Name})
}



func main() {

	TestInsert()

}

// C create/insert
func TestInsert() {
	insert(p, col)
	pb := findByName(p.Name, col)
	if pb.Name != p.Name || pb.Phone != p.Phone {
		fmt.Println("insert failed")
	}
	fmt.Println("Insert Result")
	fmt.Println(pb)
}


// R read/find
func TestFindByName() {
	p := findByName(p.Name, col)
	if p.Name == "" || p.Phone == "" {
		fmt.Println("find by name test failed")
	}
	fmt.Println("Find Result")
	fmt.Println(p)
}

// U update
func TestUpdate() {
	p := Person{"WEW", "2333","'2344444"}
	update(p, col)
	ub := findByName(p.Name, col)
	fmt.Println("Update Result")
	fmt.Println(ub)
}

// D delete
func TestDeleteData() {
	deleteData(p, col)
}



