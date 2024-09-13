/* 
Go is statically typing language like java
We have to declare the varible's datatype
*/
package main

import ("fmt")

func main(){
	/* 
	int- stores integers (whole numbers), such as 123 or -123
	float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
	string - stores text, such as "Hello World". String values are surrounded by double quotes
	bool- stores values with two states: true or false 
	*/

	var name string = "Elon Musk"
	var height float32 = 1.88 //m
	var children int = 12
	var inIndian bool = true

	
	name = "Steve jobs" // you can change the name here also
	
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(children)
	fmt.Println(inIndian)
	
	// Another way to declare variable
	company := "Tesla"
	/*
	variablename := value
	In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value)
	*/
	fmt.Println(name,"'s company is ",company)

	//Constant variable
	const birthYear int = 1971 // You can not change this
	fmt.Print(birthYear)
	
	// Multiple Variable Declaration
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)








}