package main

import (
	"fmt"
)

type circle struct {
	radius float64
	area   float64
}

//methods augments functions by adding an extra parameter section
// just after the func keyword that accepts a single argument
// this argument is a reciever
// method is  afunction with a defined reciever

// format for declaration
// func <reciever> <method_name>(<parameter>)
// <retuen_parameter>{
// 	// code
// }

func (c *circle) calcArea() {
	const pi float64 = 3.14
	c.area = pi * c.radius * c.radius
}

func (c circle) calcArea2() {
	const pi float64 = 3.14
	c.area = pi * c.radius * c.radius
}

// method set are a set of method available to the datatype
// OR the sruct , it encpsulates functionality

type student struct {
	name   string // string
	grades []int  // maps
}

// method1
func (s *student) displayName() {
	fmt.Println(s.name)
}

// method2
func (s *student) calcPercentage() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

// interfaces

// imterfaces is a method set and introuces modularitiies
// its like a blue print for method sets
// they describe all the methods of a method set by providing the function
// signature for each method
// they specify a set of methods but do not implement them

// format
// tyoe <interface_name> interface{
// 	method signatures
// }
// there can only be method set declaration in the {} no variable declaration
// and no method declaration

type FIxedDeposit interface {
	getRateOfInterest() float64
	calcReturn() float64
}

// implememnting interfaces
// interfaces are defined so that they can be implemented by other datatypes
// a type implements an interface by implementign its methods
// they are implemeted implicitlu and do not requrire any specific keyword to be implemented

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	length, breadth float64
}

func (r rect) area() float64 {
	return r.breadth * r.length
}
func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	fmt.Println("This is for methods")
	c := circle{radius: 4}
	c.calcArea()
	fmt.Printf("%+v", c)

	// calcaarea is a method that recieves c which is the pointer for the circle type
	c.calcArea2()
	fmt.Printf("%+v", c)
	// in this second method nothing changes

	// for method sets
	s := student{
		name:   "Joe",
		grades: []int{90, 75, 80},
	}
	s.displayName()
	fmt.Printf("%.2f%%", s.calcPercentage())

	rr := rect{
		length:  3,
		breadth: 4,
	}
	cc := square{
		side: 5,
	}
	printData(rr)
	printData(cc)

}

// problems and scenerios
// type Movie struct {
// 	name    string
// 	summary string
// 	rating  float32
// }

// func (m Movie) getSummary() {
// 	m.summary = "summary"
// }

// func (m *Movie) increaseRating() {
// 	m.rating += 1
// }

// func main() {
// 	mov := Movie{"xyz", "", 2.1}
// 	fmt.Printf("%+v", mov)
// 	mov.increaseRating()
// 	mov.getSummary()
// 	fmt.Printf("%+v", mov)
// }

// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// func (r Rectangle) area() int {
// 	return r.length * r.breadth
// }

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }

// type Rectangle struct {
// 	length  int
// 	breadth int
// }

// func (r Rectangle) area() int {
// 	return r.length * r.breadth
// }

// func (r *Rectangle) incLength(n int) {
// 	for i := 0; i < n; i++ {
// 		r.length += i
// 	}
// }

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// 	r.incLength(7)
// 	fmt.Println(r.area())
// 	fmt.Println(r)
// }

// type Employee struct {
// 	eid int
// 	id  int
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{i, i + 10}
// 			fmt.Println(employees[i])
// 	}
// }

// type Employee struct {
// 	eid int
// 	id  int
// }

// func (e Employee) get_id() int {
// 	return e.eid + 10
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 			employees[i] = Employee{eid: i}
// 			employees[i].id = employees[i].get_id()
// 			fmt.Printf("%+v\n", employees[i])
// 	}
// }


// problem  for interfaces
type Student interface {
	getPercentage() int
	getName()
}

type Undergrad struct {
	name   string
	grades []int
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
		sum += v
	}
	return sum / len(u.grades)
}

func printPercentage(s Student) {
	fmt.Println(s.getPercentage())
}

func main() {
	grades := []int{90, 75, 80}
	u := Undergrad{"Ross", grades}
	printPercentage(u)
}

type Student interface {
	getPercentage() int
	getName() string
}

type Undergrad struct {
	name   string
	grades []int
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
		sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	grades := []int{90, 75, 80}
	u := Undergrad{"Ross", grades}
	printData(u)
}

type Student interface {
	getPercentage() int
	getName() string
}

type Undergrad struct {
	name   string
	grades []int
}

type Postgrad struct {
	name   string
	grades []int
}

func (p Postgrad) getPercentage() int {
	sum := 0
	for _, v := range p.grades {
			sum += v
	}
	return ((sum * 100) / (len(p.grades) * 200))
}
func (p Postgrad) getName() string {
	return p.name
}

func (u Undergrad) getPercentage() int {
	sum := 0
	for _, v := range u.grades {
			sum += v
	}
	return sum / len(u.grades)
}
func (u Undergrad) getName() string {
	return u.name
}

func printData(s Student) {
	fmt.Println(s.getName())
	fmt.Println(s.getPercentage())
}

func main() {
	u := Undergrad{"Ross", []int{90, 75, 80}}
	p := Postgrad{"Joe", []int{150, 190, 185}}
	printData(u)
	printData(p)
}