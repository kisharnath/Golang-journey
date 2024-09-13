package main
/*
This file gives you a brief idea of how to define array .
*/
import ("fmt")

func main(){

	var names = [4]string{"John","Ram","David","Parag"}
	fmt.Println(names);//You can also give semicolon

	//Other ways to define array in go [...]
	var ceoName =[...]string{"Elon","Tim","Satya"} //Length is inferred
	fmt.Println(ceoName)

	//You can also use := 
	var ceoAge = [...]int{50,60,50}
	fmt.Println(ceoAge)

	// Getting array size
	fmt.Println(len(ceoAge))

	/*
	Iterating through the array using for loop
	for i:=0; i < 5; i++ {
       fmt.Println(i)
    }
	*/
	fmt.Println("-----Iterating through arrays-----")
	for i:=0; i<len(ceoAge);i++{
		fmt.Println(ceoAge[i])
	}
	g := 0
	for ; g < len(ceoAge); {
		fmt.Println(ceoAge[g])
		g++
	}

}