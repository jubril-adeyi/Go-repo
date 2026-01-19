package main

import "fmt"

func main() {
	fmt.Println("This chapter is all about arrays")
	// arrays are collection of similar data elements stored in contiguos memory location
	// collection of strings/ numbers
	//memory allocation is allocated to the arrays
	// an array takes homogenous datatypes
	// arrays have fixed length the lenght cannot be changed, no addition or removal of elements
	//length = capacity in arrays

	// declaration
	var grades [5]int
	var fruits [8]string

	fmt.Println("at this point the content of the arrays is usually zeros an empty string because no values in []")
	fmt.Println(grades)
	fmt.Println(fruits)

	// adding the value to array using curly braces {}
	grades = [5]int{55, 66, 77, 88, 99}
	fruits = [8]string{"orange", "banana", "mango"}
	fmt.Println(grades)
	fmt.Println(fruits)

	// using elipses i.e using the ... in [] basically says that the compiler
	// determines the length implicitly by itself

	nums := [...]int{6, 7, 8, 9, 0}
	fmt.Println(nums)

	// using len()
	l1 := len(nums)
	l2 := len(grades)
	l3 := len(fruits)
	lsum := l1 + l2 + l3
	fmt.Println(lsum)

	// indexing in arrays is from 0,1,2,3 fromt the first to the fourth in that list
	// the elemnts in the arrays can be called and accessed from the array usign the indexes

	// fmt.Printf("I had %v %vs and %v %vs", nums[2], fruits[0], nums[1], fruits[1])

	// or change the value

	fruits[1] = "soursop"
	fmt.Printf("I had %v %vs and %v %vs", nums[2], fruits[0], nums[1], fruits[1])

	// looping through array

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	// using range keywords

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}

	// multidimensional array
	// an array of arrays

	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr)
	fmt.Println(arr[1][0])

	// Slice
	// continuous segment of an underlying array
	// so like a section/ segment of an array

	// an elemnt can be added and removed from the slice
	// i.e  it can increase in length
	// three component pointer length and capacity
	// pointer points to the first element in the array
	// this decides the start of the slice
	// declare a slice

	// slice_name:= []datatype{values}
	grade1 := []int{10, 20, 30}

	fmt.Println(grade1)

	// to create a slice we have the start index and end index
	// the last index is not included in a slice

	// basically to get a slice from an array its array[start_index: end_index]
	// these indexes are numbers that indicate indexes

	arr2 := [10]int{10, 34, 44, 45, 67, 77, 33, 66, 77, 87}

	slice_7 := arr2[1:8]
	sub_slice := slice_7[0:5]

	fmt.Println(slice_7)
	fmt.Println(sub_slice)

	// modifying elements in a slice it modifies the underlying array too
	slice_7[4] = 9000
	fmt.Println(slice_7)
	fmt.Println(arr2)

	// declaring and initializing a slice using the make function
	// format : slice := make([]<data_type>, length, capacity)
	// but its an empty slice just declared but not filled

	make_slice := make([]int, 5, 10)
	slice_4 := make_slice[0:3]

	fmt.Println(make_slice)
	fmt.Println(len(make_slice))
	fmt.Println(cap(make_slice))

	// appndding a slice
	append_slice := append(make_slice, 4)
	// to another slice
	append_slice2 := append(slice_4, append_slice...)
	// append_slice[0:4] = [6,5,6,6,7]
	fmt.Println(append_slice)
	// fmt.Println(len(append_slice))
	// fmt.Println(cap(append_slice))
	fmt.Println(append_slice2)

	// dleeting elements in a slice
	// append manipulation

	// copying from a slice to another
	// same datatype

	// copy (destslice, srcslice)
	//  looping through a slice

	for index, value := range append_slice2 {
		fmt.Println(index, "=>", value)
	}
}
