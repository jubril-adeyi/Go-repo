package main

import (
	"fmt"
)

// passing arguments to a function by value and not by reference
// means that the function is called by directly passing the value of
// the variable as an argument
// so the parameter is copied into another location of the memory
// so when accessing or modifying the variable within function
// onkly thee copy is accessed or modified
// all basic datatype can be passed by value

// example here
func modify(a int) int {
	a += 100
	return a
}

// passing by reference in functions
// the parameter in the function is the address of the variable
// so all operations in that fujction is performed on
// the value stored in the address paremeter i.e in the same memory address

func modify2(s *string) string {
	*s = "world"
	return *s
}

// this means if i pass this function to a value boy it will change to my world

// for slice

func modify3(s []int) {
	s[0] = 100
}

// for maps

func modify4(m map[string]int) {
	m["K"] = 34
}

func main() {
	fmt.Println("this is about pointers")

	// what are pointers? whenever a variable is assigned,
	// a section of RAM is allocated to it
	// this mem allocation is done while the program runs
	// so the var may get a diffwerent address anytime the program runs

	// so a pointer is a variable that holds the memory address of another variable
	// so theres a variable and then another variable that stores the pointer of x
	// which stores the vbalue of the mempry address of x in another memory address

	// they point to the location where the memory is allocated and also
	// provide ways to find or even change the value located at the memeory location

	// pointer operators: address and dereference operators

	// address operator: the address of a variable can be gotten by preceding it
	// with a & sign, this is called address-of-operator.  so gets address

	// dereferenmce operator: when placed before an address it returns the value of
	// that address it is the * sign . get the value of the address
	var x int = 77

	// get memory address of x
	fmt.Println(&x)
	// get what is stored in thst memory address, the value in the address
	fmt.Println(*(&x))

	// with their data types
	fmt.Printf("%T %v \n", &x, &x)
	fmt.Printf("%T %v \n", *(&x), *(&x))

	// declaring pointers
	var i *int
	var s *string
	fmt.Println(i)
	fmt.Println(s)

	// initializzing pointer
	// var <pointer_name> *<data_type> = &<variable_name>
	// OR
	// var <pointer_name> = &<variable_name> letting the compiler infer the datatype.

	p := 10
	var ptr_p *int = &p
	fmt.Println(ptr_p)

	//

	a := "hello"
	var b *string = &a
	fmt.Println(b)
	var c = &a
	fmt.Println(c)
	d := &a
	fmt.Println(d)

	// dereferencing pointers
	//pointer d
	*d = "yeah"
	fmt.Println(*d)

	// passing arguments to a function by value

	num := 20
	fmt.Println(num)
	modify(num)
	fmt.Println(num)

	// point is that the num value still remains the same in the
	// address in memory its been allocated
	// another memory space is allocated to modify(num) which is the function applied to num

	word := "boy"
	fmt.Println(word)
	modify2(&word)
	fmt.Println(word)

	// the original word var changes
	// it chnages because its is pointed to the same memory location

	// slices are passed by reference by default
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modify3(slice)
	fmt.Println(slice)

	// as known before, the value in the memory space
	// usually chanbges after any reassigment like the function is doing here

	//in maps too its the same
	maps := make(map[string]int)
	maps["A"] = 65
	maps["B"] = 89
	fmt.Println(maps)
	modify4(maps)
	fmt.Println(maps)
}

// // problem
// func modify(y int) int {
// 	y += 15
// 	return y
// }
// func main() {
// 	y := 20
// 	modify(y)
// 	fmt.Println(y)
// }

// func main() {
// 	y := [3]int{10, 20, 30}
// 	py := &y
// 	fmt.Printf("%T %v \n", py, *py)
// }

// func main() {
// 	y := [3]int{10, 20, 30}
// 	fmt.Printf("%v \n", y)
// 	(*&y)[0] = 100
// 	fmt.Printf("%v \n", y)
// }

// func main() {
// 	var y int
// 	var ptr *int = &y

// 	*ptr = 0
// 	fmt.Println(y)

// 	*ptr += 5
// 	fmt.Println(y)

// }

// func main() {
// 	var y int
// 	var ptr *int = &y
// 	fmt.Println(y)
// 	fmt.Println(*ptr)
// }

// func main() {
// 	s := 100
// 	var ptr *string = &s
// 	fmt.Println(s)
// 	*ptr += 100
// 	fmt.Println(s)
// }

// func main() {
// 	s := "hello"
// 	var ptr *string = &s
// 	fmt.Println(s)
// 	*ptr += strings.ToUpper(s)
// 	fmt.Println(s)
// }

//more problems on passing as value and as reference
// func modify(numbers ...int) {
// 	for i := range numbers {
// 			numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := []int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr...)
// 	fmt.Println(arr)
// }

// func modify(numbers [3]int) {
// 	for i := range numbers {
// 			numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// func modify(numbers *[3]int) {
// 	for i := range numbers {
// 			numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// func modify(s *string) {
// 	*s = strings.ToUpper(*s)
// }
// func main() {
// 	s := "hello"
// 	fmt.Println(s)
// 	modify(&s)
// 	fmt.Println(s)
// }

// func modify(s map[string]int) {
// 	s["A"] = 100
// }
// func main() {
// 	ascii_codes := map[string]int{}
// 	ascii_codes["A"] = 65
// 	fmt.Println(ascii_codes)
// 	modify(ascii_codes)
// 	fmt.Println(ascii_codes)
// }
