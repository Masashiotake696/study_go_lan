package main

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
)

type MyInt int

type (
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
)

type Callback func(i int) int

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

type T0 int
type T1 int

type Point struct {
	X, Y int
}

type T struct {
	int
	float64
	string
}

type Feed struct {
	Name   string
	Amount uint
}

type Animal struct {
	Name string
	Feed
}

type S0 struct {
	Name1 string
}

type S1 struct {
	S0
	Name2 string
}

type S2 struct {
	S1
	Name3 string
}

type Base struct {
	Id    int
	Owner string
}

type A struct {
	Base
	Name string
	Area string
}

type B struct {
	Base
	Title  string
	Bodies []string
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

func swap(p Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func swapPointer(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

type Person struct {
	Id   int
	Name string
	Area string
}

func (p *Point) Render() {
	fmt.Println(p.X, p.Y)
}

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

type IntPoint struct {
	X, Y int
}

type FloatPoint struct {
	X, Y float64
}

func (p *IntPoint) Distance(dp *IntPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

func (s Strings) Join(d string) (sum string) {
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return
}

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

type Points []*Point

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d %d]", p.X, p.Y)
}

func (p Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Point) PointerSet(x, y int) {
	p.X = x
	p.Y = y
}

func (ps Points) ToString() (str string) {
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return
}

type Man struct {
	Id   int    `ID`
	Name string `名前`
	Age  uint   `年齢`
}

type Woman struct {
	Id   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  uint   `json:"user_age"`
}

func main() {
	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1)
	fmt.Println(n2)

	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)
	fmt.Println(pair)
	fmt.Println(strs)
	fmt.Println(amap)
	fmt.Println(ich)

	fmt.Println(
		Sum(
			[]int{1, 2, 3, 4, 5},
			func(i int) int {
				return i * 2
			},
		),
	)

	t0 := T0(5)
	i0 := int(t0)
	t1 := T1(8)
	i1 := int(t1)
	fmt.Printf("%T\n", t0)
	fmt.Printf("%T\n", i0)
	fmt.Printf("%T\n", t1)
	fmt.Printf("%T\n", i1)

	var pt1 Point
	fmt.Println(pt1.X, pt1.Y)
	pt1.X = 10
	pt1.Y = 8
	fmt.Println(pt1.X, pt1.Y)

	pt2 := Point{1, 2}
	fmt.Println(pt2.X, pt2.Y)

	pt3 := Point{X: 1, Y: 2}
	fmt.Println(pt3.X, pt3.Y)

	pt4 := Point{Y: 10}
	fmt.Println(pt4.X, pt4.Y)

	t := T{1, 3.14, "文字列"}
	fmt.Println(t.int, t.float64, t.string)

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Amount)
	a.Amount = 15
	fmt.Println(a.Amount)

	s := S2{
		S1: S1{
			S0: S0{
				Name1: "X",
			},
			Name2: "Y",
		},
		Name3: "Z",
	}
	fmt.Println(s.Name1)
	fmt.Println(s.Name2)
	fmt.Println(s.Name3)

	aBase := A{
		Base: Base{
			Id:    17,
			Owner: "Taro",
		},
		Name: "Taro",
		Area: "Tokyo",
	}
	bBase := B{
		Base: Base{
			Id:    81,
			Owner: "Hanako",
		},
		Title:  "No Title",
		Bodies: []string{"A", "B"},
	}
	fmt.Println(aBase.Id)
	fmt.Println(aBase.Owner)
	fmt.Println(bBase.Id)
	fmt.Println(bBase.Owner)

	st := struct{ X, Y int }{X: 1, Y: 2}
	showStruct(st)

	p := Point{X: 7, Y: 8}
	showStruct(p)

	p1 := Point{X: 1, Y: 2}
	fmt.Println(p1)
	swap(p1)
	fmt.Println(p1)

	p2 := Point{X: 1, Y: 2}
	fmt.Println(p2)
	fmt.Println(p2.X, p2.Y)
	swapPointer(&p2)
	fmt.Println(p2)
	fmt.Println(&p2)

	person := new(Person)
	fmt.Println(person)
	fmt.Println(person.Id)
	fmt.Println(person.Name)
	fmt.Println(person.Area)

	i := new(int)
	fmt.Println(*i)

	strings := new([]string)
	fmt.Println(*strings)

	p3 := &Point{X: 5, Y: 12}
	p3.Render()
	fmt.Println(p3.Distance(&Point{X: 1, Y: 1}))

	fmt.Println(MyInt(3).Plus(3))

	ip := IntPair{1, 2}
	fmt.Println(ip.First())
	fmt.Println(ip.Last())

	fmt.Println(Strings{"A", "B", "C"}.Join(","))

	fmt.Println(NewUser(1, "Taro"))

	f := (*Point).ToString
	fmt.Println(f(&Point{X: 7, Y: 11}))
	fmt.Println(((*Point).ToString)(&Point{X: 1, Y: 2}))

	p4 := &Point{X: 2, Y: 3}
	f1 := p4.ToString
	fmt.Println(f1())

	p5 := Point{}
	fmt.Println(p5.X, p5.Y)
	p5.Set(1, 2)
	fmt.Println(p5.X, p5.Y)

	p6 := &Point{}
	fmt.Println(p6.X, p6.Y)
	p6.Set(1, 2)
	fmt.Println(p6.X, p6.Y)

	p7 := Point{}
	fmt.Println(p7.X, p7.Y)
	p7.PointerSet(1, 2)
	fmt.Println(p7.X, p7.Y)

	p8 := &Point{}
	fmt.Println(p8.X, p8.Y)
	p8.PointerSet(1, 2)
	fmt.Println(p8.X, p8.Y)

	ps1 := make([]Point, 5)
	for _, p := range ps1 {
		fmt.Println(p.X, p.Y)
	}

	ps2 := Points{}
	ps2 = append(ps2, &Point{X: 1, Y: 2})
	ps2 = append(ps2, nil)
	ps2 = append(ps2, &Point{X: 3, Y: 4})
	fmt.Println(ps2.ToString())

	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Tokyo",
		{Id: 2, Name: "Jiro"}: "Osaka",
	}
	fmt.Println(m1)
	m2 := map[int]User{
		1: {Id: 1, Name: "Tara"},
		2: {Id: 2, Name: "Jiro"},
	}
	fmt.Println(m2)

	ms := map[int][]string{
		1: {"A", "B", "C"},
	}
	fmt.Println(ms)
	mm := map[int]map[int]string{
		1: {
			1: "Apple",
			2: "Banana",
			3: "Cherry",
		},
	}
	fmt.Println(mm)

	man := Man{Id: 1, Name: "Taro", Age: 32}
	mant := reflect.TypeOf(man)
	for i := 0; i < mant.NumField(); i++ {
		f := mant.Field(i)
		fmt.Println(f.Name, f.Tag)
	}

	woman := Woman{Id: 1, Name: "Hanako", Age: 20}
	bs, _ := json.Marshal(woman)
	fmt.Println(string(bs))
}
