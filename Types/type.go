package main
import ("fmt")

/*
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.
*/

type Car struct {
	modelName string
	age int
	ownerName string
	milage float32
}


func main(){
	var car1 Car
	car1.modelName = "Volvo"
	car1.age = 5
	car1.ownerName = "Johnny"
	car1.milage = 32.5

	fmt.Println(car1.modelName)
	fmt.Println(car1.age)
	fmt.Println(car1.ownerName)
	fmt.Println(car1.milage)

}