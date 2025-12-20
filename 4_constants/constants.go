package main

import "fmt"

//const age = 30

func main(){

	//const name = "golang"
	//const age = 20

	//fmt.Println(age)

	const(
		port = 5000
		host = "locathost"
	)

	fmt.Println(port, host)

}