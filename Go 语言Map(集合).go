package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	capital,ok := countryCapitalMap ["美国"]

	if (ok) {
		fmt.Println("美国的首都是",capital)
	}else {
		fmt.Println("美国的首都是22333",capital)

	}

	delete(countryCapitalMap,"France")
	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	r := Rect{width:10,height:5}
	fmt.Println("area:",r.Area())


	y := &Rect{width:10,height:5}
	x := &Rect{width:10,height:5}
	rs := Rects{y,x}

	fmt.Println("r's area: ", r.Area())
	fmt.Println("x's area: ", x.Area())
	fmt.Println("total area: ", rs.Area())

	var z Foo
	z = func() int {return 1}

	fmt.Println(z())
	fmt.Println(z.Add(3))

	p := Person1{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gotham",
			State:  "NY",
			Zip:    "01313",
		},
	}

	fmt.Println(p.String())

}


type Rect struct {
	width int
	height int
}

func (r *Rect) Area() int {
	return r.width * r.height
}

type Rects []*Rect

func (rs Rects) Area() int {
	var a int
	for _,r := range rs {
		a += r.Area()
	}
	return a
}





type Foo func() int

func (f Foo) Add(x int) int {
	return f() + x
}


type Person1 struct {
	Name string
	Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (a *Address) String() string {
	return a.Number + " " + a.Street + "\n" + a.City + ", " + a.State + " " + a.Zip + "\n"
}

func (a *Address) String1() string {
	return a.Number + " " + a.Street + "\n" + a.City + ", " + a.State + " " + a.Zip + "\n"
}

func (p *Person) String() string {
	return p.Name + "\n" + p.Address.String()
}

type Shaper interface{
	Area() int
}

func Describe(s Shaper) {
	fmt.Println("Area is: ", s.Area())
}





