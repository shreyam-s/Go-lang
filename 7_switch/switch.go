package main

import "fmt"

func main(){

	//simple switch
	// day := 3

	// switch day{
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	println("Thursday")
	// }

	//Multiple Case
	// n := 2

	// switch n{
	// case 1,3,5:
	// 	fmt.Println("Odd")
	// case 2,4,6:
	// 	println("Even")
	// }

	// switch without an expression
	// x:=75
	// switch{
	// case x>=90:
	// 	fmt.Println("Grade A")
	// case x >= 60:
    //     fmt.Println("Grade B")
    // case x >= 40:
    //     fmt.Println("Grade C")
    // default:
    //     fmt.Println("Fail")
	// }

	// switch with short hand statement
	// switch x:=10; x{
	// case 5:
	// 	fmt.Println("Odd")
	// case 10:
	// 	println("Even")
	// }

	// Nested switch
	// x:=1
	// y:=2

	// switch x{
	// case 1:
	// 	switch y{
	// 	case 2:
	// 		fmt.Println("x=1,y=2")
	// 	}
	// }

	// type switch
	var i interface{} = "Golang"

	switch v:=i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		println("String", v)
	default:
		println("Unknown")
	}
}
