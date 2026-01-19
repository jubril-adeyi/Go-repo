package main

import (
	"fmt"
)

func main() {
	fmt.Println("This section is about maps in Go")

	// maos are unordered collcectio of key/value pairs in
	// dictionaries in python
	// values are stored based on the key, its like an index that is used to retieve
	// data from the map except it is not regular number index by the keyword/ key

	// it is implemented by hash tables
	// so they have the add delete get operations /

	// Decalring maps

	// var <map-name> map[key_data_type]<value_data_type\
	var map1 map[string]int

	// with values, initialize it this is the only way to add values,
	// cannot append to this map map1 or map2 (if created to be empty)
	// it has to be reinitialized
	// values can be added to maps created with make function though
	// <map_name> := map[<key_data_type]value_data_type{key_value_pairs}
	map2 := map[string]int{"one": 1}
	map3 := map[string]string{"en": "english", "fr": "french", "esp": "spanish", "rus": "russian"}

	// declaring and initializing using make() function
	// <map_name> := make(map[<key_data_type>]<value_data_type, <initial_capacity)
	// initial_capacity is an optional argument

	// map4 := make(map[string]string{})

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	// accessing values in a map
	fmt.Println(map3["en"])

	// using value and boolean; here found is the boolean
	value, found := map3["fr"]
	fmt.Println(found, value)

	value1, found1 := map3["man"]
	fmt.Println(value1, found1)

	// adding a new index key and value to a map

	map3["ita"] = "italia"
	fmt.Println(map3)
	// how do i specify the position of this new key value pair??? a;phabetical ordder

	// updating the value of a key
	map3["ita"] = "italian"
	fmt.Println(map3)

	// deleting a keyvalue pair from a map
	// delete function takes two parameters the map name and the key name
	delete(map3, "en")
	fmt.Println(map3)

	// length of map
	fmt.Println(len(map3))

	// when iterating over a map using thr for loop maybe combined with range
	// the two assigment elements used are for the key and the value
	// as opposed to the value and boolean that is the normal convention

	// for key, value := range map3 {
	// 	fmt.Println(key, "=>", value)
	// }

	// truncating map / shorteneing the map
	// using an for and range interation style

	for key2, value2 := range map3 {
		delete(map3, key2)
		fmt.Println(key2, "=>", value2)
		fmt.Println(map3)
	}
	fmt.Println(map3)

	// OR just reinitialize it with zero or less values like

	map3 = make(map[string]string)
	fmt.Println(map3)
}
