package main

import "fmt"

// for loop -> only construct in golang for looping.
// no other looping method is there in golang.
func main(){

	//while loop
	// i := 1
	// for i<5{
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for{
	// 	fmt.Println("Hello golang!")
	// }

	//classic for loop
	// for i:=0 ; i<=5; i++ {
	// 	//if i==2{
	// 	//	break
	// 	//}
	//  //fmt.Println(i)

	// 	if i==2{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//range
	// for i := range 5 {
	// 	fmt.Println(i)
	// }

	//range
	for i, v := range []int{1,2,3,4,5} {
		fmt.Println(i, v)
	}
}
