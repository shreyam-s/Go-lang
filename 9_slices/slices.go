package main

// -> slice is a dynamic, flexible view over an array
// -> most used construct in golang
// -> + useful methods

// // -> for slicing in function
// func change(s []int){
// 	s[1] = 2
// }

func main(){

	// uninitialized slice is nil
	// var nums []int
	// // -> len() = length of a slice && cap() = capacity of a slice.
	// fmt.Println(nums, len(nums), cap(nums))


	// creating slice using make
	// var s = make([]int, 2, 3)
	// fmt.Println(s, len(s), cap(s))


	// slice literal
	// b := []int{10, 20, 30}
	// fmt.Println(b)


	// Slice from an array
	// arr := [5]int{1, 2, 3, 4, 5}
	// // slice syntax -> slice[start:end]
	// s := arr[1:4]
	// fmt.Println(s)


	// Append() -> append may create a new backing array.
	// s := []int{1, 2}
	// // s = append(s, 3)

	// // -> Append Multiple Values
	// s = append(s, 3, 4, 5)
	// fmt.Println(s)


	// Append Slice to Slice
	// a := []int{1, 2}
	// b := []int{3, 4}
	// a = append(a, b...)
	// fmt.Println(a)


	// // Slice in reference type
	// arr := [4]int{1, 2, 3, 4}
	// s := arr[:]
	// s[1] = 99
	// fmt.Println(arr)


	// // Slice in function
	// s := []int{1, 100, 3, 4}
	// change(s)
	// fmt.Println(s)


	// // Copying Slices
	// a := []int{1, 2, 3}
    // b := make([]int, len(a))
    // copy(b, a)

    // b[0] = 99
    // fmt.Println(a)
    // fmt.Println(b)
}
