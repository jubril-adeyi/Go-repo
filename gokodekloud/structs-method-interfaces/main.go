package main

import (
	"fmt"
)

// structs are structures of data elements
// user defined datatype this means I Name the datatype which is crazy
// provide a way to reference a series of grouped
// values through a single variable name
// used when it makes sense to group or associate two/more data variables

// Declaring a structs:

// type <struc_name> struct{
// 	//list of feilds
// }
// struts can be defined outside the main func

type circle struct {
	x      float64
	y      float64
	z      float64
	radius float64
	area   float64
}

type student struct {
	name   string         // string
	rollNo int            // integer
	marks  []int          // slice
	grades map[string]int // maps
}

// passing structs to functions

//	func calcArea(c circle) {
//		const pi float64 = 3.14
//		var area float64
//		area = (pi * c.radius * c.radius)
//		c.area = area
//	}
func calcArea(c *circle) {
	const pi float64 = 3.14
	var area float64
	area = (pi * c.radius * c.radius)
	(*c).area = area
}

// comparing structs tosee if theu have the same tyoeor same values
type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	fmt.Println("This is for Structs")

	// Struct initialization
	var s student
	fmt.Printf("%+v", s)

	// initializing a struct using the new keyword
	// <variable_name> := new(<struct_name>)
	c := new(circle)
	fmt.Printf("\n %+v", c)

	// initializing a struct with values
	// 	<variable_name>:= <struct_name> {
	// 		<field_name> : <value>,
	// 		<field_name> : <value>,
	// 	}
	// not neccesary to fill in all the keys with values

	c1 := circle{
		x: 4.00,
		y: 5.00,
		// z: 6.00,
	}
	// fmt.Println("\n", c1)
	fmt.Printf("\n %+v", c1)

	// accessing fields in a struct
	// <variable_name>.<field_name>

	fmt.Println("\n", c1.y)

	c3 := circle{
		x:      5,
		y:      6,
		z:      8,
		radius: 9,
		area:   0,
	}
	// fmt.Printf("\n %+v", c3)
	// calcArea(c3)
	// fmt.Printf("\n %+v", c3)

	// no change
	// // so we use pssing bu reference here modified in the funtion

	fmt.Printf("\n %+v", c3)
	calcArea(&c3)
	fmt.Printf("\n %+v", c3)

	// comparing structs tosee if theu have the same tyoeor same values
	// type s1 struct {
	// 	x int
	// }

	// type s2 struct {
	// 	x int
	// }

	d1 := s1{x: 4}
	d2 := s1{x: 4}
	if d1 == d2 {
		fmt.Println("yes")
	}
}

// // problem
// type Movie struct {
// 	name   string
// 	rating float32
// }

// func main() {
// 	m := Movie{name: "ABCD"}
// 	var m2 Movie
// 	fmt.Printf("%+v", m)
// 	fmt.Printf("%+v", m2)
// }
// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func main() {
// 	fmt.Printf("%+v", getMovie("xyz", 3.5))
// }

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func increaseRating(m *Movie) {
// 	m.rating += 1.0
// }

// func main() {
// 	mov := getMovie("xyz", 2.0)
// 	increaseRating(mov)
// 	fmt.Printf("%+v", mov)
// }

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func main() {
// 	mov := getMovie("xyz", 2.0)
// 	fmt.Println(mov.name)
// 	fmt.Println(mov.ratings)
// }

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func getMovie(s string, r float32) (m Movie) {
// 	m = Movie{
// 			name:   s,
// 			rating: r,
// 	}
// 	return
// }

// func main() {
// 	mov := getMovie("xyz", 2.1)
// 	mov1 := getMovie("abc", 3.3)
// 	movies := make([]Movie, 5)
// 	movies = append(movies, mov)
// 	movies = append(movies, mov1)
// 	for _, value := range movies {
// 			fmt.Println(value)
// 	}
// }

// type Movie struct {
// 	name   string
// 	rating float32
// }

// func main() {
// 	mov := Movie{"xyz", 2.1}
// 	mov1 := Movie{"abc", 2.1}
// 	if mov.rating == mov1.rating || mov != mov1 {
// 			fmt.Println("condition met")
// 	} else if mov.rating == mov1.rating {
// 			fmt.Println("condition_2 met")
// 	}
// }
