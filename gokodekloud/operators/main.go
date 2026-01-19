package main

import "fmt"

func main() {
	// in this arithmetic the operators are the + - * ++ --  etc
	// mathematical

	product := 4 * 9
	sum := 4 + 9
	a, b := "foo", "bar"

	fmt.Print("product = ", product, "\n")
	fmt.Print("sum = ", sum, "\n")
	fmt.Println(a + b)
	fmt.Println(product - sum)
	fmt.Println(product / sum)
	//modulus gives the remainder
	fmt.Println(product % 5)

	//increment ++
	// it is a unary operator acts on a single operand decreasing or incresing by 1

	i := 1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	// we have comparison operator
	// == > < etc
	// this retuns true or false
	fmt.Println("comaprison operators \n")

	city1 := "tokyo"
	city2 := "lagos"

	fmt.Println(city1 == city2)
	fmt.Println(city1 != city2)
	fmt.Println(product > sum)
	fmt.Println(product < sum)

	// logical operators determine the logic btw vaiables or values
	// && and  || OR  ! Not(seen in combination with the < or > or ==)
	fmt.Println("logical operators \n")
	x := 10
	fmt.Println((x < 100) && (x < 200))
	fmt.Println((x < 300) && (x < 0))
	fmt.Println((x < 100) || (x < 200))

	// ! is also unary
	fmt.Println(!(x > 20))

	// assignment operartors += -= *= /= %=

	fmt.Println("Assignment operators \n")
	y := 10

	// x+=y means x= x+y
	x += y
	fmt.Println(x)

	x -= y
	fmt.Println(x)

	x *= y
	fmt.Println(x)

	x /= y
	fmt.Println(x)

	x %= 3
	fmt.Println(x)

	// bitwise operastors
	// | or & and ^ XOR >> right shift << left shift
	// & takes  two operands and does AND on every bit of the two numbers i.e in binary

	x = 12
	y = 25

	z := x & y
	fmt.Println(z)

	// | does OR on everybit of the two numbers in binary

	z = x | y
	fmt.Println(z)

	// ^ does XOR on the bits, XOr is the 0/false for similar and 1/true for diff

	z = x ^ y
	fmt.Println(z)

	// << shifts all bit to the left by a certain number of specified bits
	z = x << y
	fmt.Println(z)

	// >> shifts all bit to the right by a certain number of specified bits
	z = 212 >> 2
	fmt.Println(z)

}
