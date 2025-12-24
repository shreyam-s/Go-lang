package main

// A map stores data as a key -> value pairs
// -> map[keyType]valueType

func main(){

	// // Declare map
	// var m map[string]int
	// fmt.Println(m)
	// // You cannot insert in a nil map.


	// // Initialize a map using make
	// m := make(map[string]int)
	// m["Apple"] = 10
	// m["Banana"] = 20

	// fmt.Println(m)


	// // Map literal
	// m := map[string]int{
	// 	"go" : 1,
	// 	"java" : 2,
	// 	"python" : 3,
	// }
	// fmt.Println(m)


	// // Accessing map values
	// m := map[string]int{
	// 	"go" : 1,
	// 	"java" : 2,
	// 	"python" : 3,
	// }

	// fmt.Println(m["go"])
	// fmt.Println(m["java"])
	// fmt.Println(m["python"])


	// // Check if a key exists
	// m := map[string]int{
	// 	"go" : 1,
	// 	"java" : 2,
	// 	"python" : 3,
	// }

	// if val, ok := m["go"]; ok {
	// 	fmt.Println(val, ok)
	// }else{
	// 	fmt.Println("Key does not exist")
	// }


	// // Update map value
	// m := map[string]int{"a":1}
	// m["a"]=100
	// fmt.Println(m)


	// // Delete from map
	// m := map[string]int{"a":1, "b":2, "c":3}
	// delete(m, "a")
	// fmt.Println(m)


	// // Loop through map
	// m := map[string]int{"a":1, "b":2, "c":3}
	// for key, value := range m{
	// 	fmt.Println(key, value)
	// }


	// // Loop through map keys
	// m := map[string]int{"a":1, "b":2, "c":3}
	// for key := range m{
	// 	fmt.Println(key)
	// }


	// // Loop through map values
	// m := map[string]int{"a":1, "b":2, "c":3}
	// for _, value := range m{
	// 	fmt.Println(value)
	// }


	// // Get length of map
	// m := map[string]int{"a":1, "b":2, "c":3}
	// fmt.Println(len(m))


	// // Get all keys
	// m := map[string]int{"a":1, "b":2, "c":3}
	// keys := make([]string, len(m))
	// i := 0
	// for key := range m{
	// 	keys[i] = key
	// 	i++
	// }
	// fmt.Println(keys)
}


// Maps are unordered, reference-type key-value stores with fast lookup.
