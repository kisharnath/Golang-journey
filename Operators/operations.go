package main
import ("fmt")

/*
+	Addition	Adds together two values	x + y	
-	Subtraction	Subtracts one value from another	x - y	
*	Multiplication	Multiplies two values	x * y	
/	Division	Divides one value by another	x / y	
%	Modulus	Returns the division remainder	x % y	
++	Increment	Increases the value of a variable by 1	x++	
--	Decrement	Decreases the value of a variable by 1	x--
*/

func main(){
	fmt.Println("2+4",2+4);
	fmt.Println("2-5",2-5);
	fmt.Println("2*5",2*5);
	fmt.Println("5/2",5/2);
	fmt.Println("5%2",5%2);

	var x int = 5
	x++
	fmt.Println("5+1",x);
	x--
	fmt.Println("6-1",x);

	// Assignment operators
	fmt.Println(x&3)

}