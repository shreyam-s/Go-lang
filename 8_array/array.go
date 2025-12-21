package main

import "fmt"

func main(){

	// Declare an array
	// var arr [5]int
	// fmt.Println(arr)

	// Declare and Initialize an array
	// var arr = [3]int{10, 20, 30}
	// fmt.Println(arr)

	// Short Declaration
	// arr := [3]string{"Go", "Java", "Python"}
	// fmt.Println(arr)

	// Array with inferred size
	// arr := [...]int{10, 20, 30, 40}
	// fmt.Println(arr) 

	// Access and Modify Array element
	// arr := [3]int{10, 99, 30}
	// arr[1] = 20
	// fmt.Println(arr)

	// Loop through an Array
	// arr := [3]int{10, 20, 30}
	
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// Loop using for range
	// arr := [3]int{10, 20, 30}

	// for index, value := range arr{
	// 	fmt.Println(index, value)
	// }

	// Ignore index or value
	// arr := [3]int{10, 20, 30}

	// for _,v := range arr{
	// 	fmt.Println(v)
	// }

	// Multidimensional Array
	arr := [2][3]int{{1, 2, 3},{4, 5, 6}}
	fmt.Println(arr)
}
