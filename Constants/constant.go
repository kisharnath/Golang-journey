package main

import ("fmt")
/*
Refer to this document
https://www.w3schools.com/go/go_constants.php
*/
func main(){

	// Declaring one constant
	const A1 int = 1

	// Declaring multiple constant type is inferred
	const (
		A int = 1
  		B = 3.14
  		C = "Hi!"
	)
	fmt.Println(A1)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}