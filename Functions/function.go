package main
import ("fmt")

func addTwo(a int, b int) int{
	return a+b
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
  }

func main(){

	fmt.Println(addTwo(5, 5))
	fmt.Println(myFunction(5, "Hello"))

}