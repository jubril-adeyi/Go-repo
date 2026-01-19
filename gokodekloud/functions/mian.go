package main

import "fmt"

// Okay i know what functions are
// functions takes input and then returns output
// great for reuseablility and abstract inner implementations
// FUNCTIONS CANNOT BE DEFINED OUTSIDE THE MAIN FUNCTION IN GOLANG-
// ITS LIKE LAMBDA FUNCTIONS HAVING FUNCTIONS AND THE LAMBDA HANDLER FUNCTION WITHIN IT

// hers is the syntax of the functions

// func <function_name> (<params>) <return data_type> {
// 	// body of the function
// }

// parameters and arguments
// parameters are initialized with the function
// aregument are suppliced when function is called

func addNumbers(a int, b int) int {
	// body of the function
	sum := a + b
	return sum
}

func printgreeting(str string) string {
	fmt.Println("Hey there", str)
	return ""
}

// returning multiple values
// note how sum and diff values do not need := initialization within func /
// because they have been initialized within the output parameter definitions
// and slso return statement doesnt require any aregument
func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = b - a
	return
}

// variadic functions
// functions that take a variable number of arguments
// pass varying number of arguments of the same type as referenced in func signature
// when declaring these kind of functions the type of the final parameter
// is preceded by ... (elipisi)

// func <func_name> (par1 type, par2 type, par3 ...type) <return_type> ...const
// this case the variadicity lies on the par3
// variabdic parameters also have to be the last

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
		// sum = sum + value
	}
	return sum
}
func printDetails(student string, subjects ...string) string {
	fmt.Println("hey, ", student, ", here are your subjects- ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
	return ""
}

// blank identifier '_' in use it reprsent a placeholder that we want ignored
//  in the case above its the index ofthe slice for the ranges in the loops

func f() (int, int) {
	return 42, 53
}

// Recusive function where the functions calls itself until it reaches base case
// useful to solve problems where the solutyion is dependent on a smaller instance of the same problem

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// this way the factorial function is used within itself
// base case is the first conditional with n ==1

// High Order Functions
// functions that recieves a function as argument or retuens a fucntions as outut
// smaller function can do small logic while bigger high order
// ones handle bigger code using the
// smaller ones inside

func calcArea(r float64) float64 {
	return 3.14 * r
}

func calcCircumfrence(r float64) float64 {
	return 2.00 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return r * 2.00
}

// // high order intervention
// func printResult(radius float64, calcFunction func(r float64) float64) {
// 	result := calcFunction(radius)
// 	fmt.Println("Result: ", result)
// 	fmt.Println("Thank you")
// }

// func getFunction(query int) func(r float64) float64 {
// 	query_to_func := map[int]func(r float64) float64{
// 		1: calcArea,
// 		2: calcCircumfrence,
// 		3: calcDiameter,
// 	}
// 	return query_to_func[query]
// }

// Defer statement : it delays the execution of a function until the surrounding function returns
// the deffered call args are evaluated immediately but the function call is not executed until the surrounding
// function returns

func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {

	fmt.Println("This section is about fuctions")

	fmt.Println(addNumbers(4, 5))

	printgreeting("joe")

	sum, difference := operation(4, 5)
	fmt.Println(sum, difference)

	variadic := sumNumbers(4, 5, 6, 7, 7, 7, 6, 5, 5)
	fmt.Println(variadic)
	fmt.Println(sumNumbers(6, 7, 78, 8, 8))

	printDetails("kemi", "english", "soc", "maths")

	a, b := f()
	fmt.Println(a, b)

	a, _ = f()
	fmt.Println(a)

	fmt.Println(factorial(5))

	// Anonymous function
	// no name identifier
	// act like normal functions with taking inputs and giving outputs
	// they are used to contain functionality that does not need to be named and maybe
	// used for short term use \

	x := func(l int, b int) int {
		return l * b
	}
	// no name after func and saved in var x
	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 24))

	// define and call directly
	y := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Printf("%T \n", y)
	fmt.Println(y)

	// // HighOrder USage
	// var query int
	// var radius float64

	// fmt.Print("enter the radius of the circles: ")
	// fmt.Scanf("%f", &radius)
	// fmt.Println(("enter \n 1 - area \n 2 - cricumfrence \n 3- diameter: "))
	// fmt.Scanf("%d", &query)

	// if query == 1 {
	// 	fmt.Println("result: ", calcArea(radius))
	// } else if query == 2 {
	// 	fmt.Println("result: ", calcCircumfrence(radius))
	// } else if query == 3 {
	// 	fmt.Println("result: ", calcDiameter(radius))
	// } else {
	// 	fmt.Println("invalid query")
	// }
	// // functions are called based on conditions here

	// // there may be issues if we add more shapes or more properties to the shape
	// // high order solves this
	// // liek this
	// printResult(radius, getFunction(query))

	// usinf Defer statement
	printName(("Joe"))
	defer printRollNo(23)
	printAddress("Odo-Arawa street")
}

// 	problem for high order function
// 	package main

// import "fmt"

// func addHundred(x int) int {
//         return x + 100
// }
// func partialSum(x ...int) func() {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
//         return func() {
//                 fmt.Println(addHundred(sum))
//         }
// }
// func main() {
//         partial := partialSum(1, 2, 3, 4, 5)
//         partial()
// }

// package main

// func addHundred(x int) int {
//         return x + 100
// }
// func partialSum(x ...int) func() int {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
//         return func() int {
//                 return addHundred(sum)
//         }
// }
// func main() {
//         partial := partialSum(1, 2, 3)
//         partial()
// //

// package main

// import "fmt"

// func addHundred(x int) int {
//         return x + 100
// }
// func partialSum(add100 func(x int) int, x ...int) int {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
//         return add100(sum)

// }
// func main() {
//         partial := partialSum(addHundred, 1, 2, 3)
//         fmt.Println(partial)
// }

// 	package main

// import "fmt"

// func addHundred(x int) {
//         fmt.Println(x + 100)
// }
// func partialSum(add100 func(x int), x ...int) int {
//         sum := 0
//         for _, value := range x {
//                 sum += value
//         }
//         add100(sum)
//         return 0
// }
// func main() {
//         partial := partialSum(addHundred, 1, 2, 3)
//         fmt.Println(partial)
// }

// problem for deferred statements

// package main

// import "fmt"

// func printString(str string){
//         fmt.Printf("%q ", str)
// }

// func printInt(i int){
//         fmt.Printf("%d ", i)
// }

// func printFloat(f float64){
//         fmt.Printf("%.2f ", f)
// }
// func main() {
//         printString("browser")
//         defer printInt(32)
//         defer printFloat(0.24)
//         printString("chrome")
//         printInt(90)
//         defer printFloat(89)
//         printInt(900)
// }

// package main

// import (
//         "fmt"
//         "strings"
// )

// func getString(str string) (string, string) {
//         return strings.ToLower(str), strings.ToUpper(str)
// }

// func main() {
//         _, lower := getString("BROWSER")
//         fmt.Println(lower)
// }

// func greetings() (x, y string) {
// 	x := "hello "
// 	y := "world"
// }

// func main() {
// 	fmt.Print(greetings())
// }

// package main

// import (
//         "fmt"
// )

// func main() {
//         fmt.Println(f1())
// }

// func f1() int {
//         return f2()
// }

// func f2() int {
//         return 1
// }
