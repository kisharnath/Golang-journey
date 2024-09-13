package main
import("fmt")

/*
Assignment operators
=	x = 5	x = 5	
+=	x += 3	x = x + 3	
-=	x -= 3	x = x - 3	
*=	x *= 3	x = x * 3	
/=	x /= 3	x = x / 3	
%=	x %= 3	x = x % 3	
&=	x &= 3	x = x & 3	
|=	x |= 3	x = x | 3	
^=	x ^= 3	x = x ^ 3	
>>=	x >>= 3	x = x >> 3	
<<=	x <<= 3	x = x << 3
*/
func main(){

	// &
	x:= 10 //binary of 10 is 1010 
	x &= 3 // binary of 3 is 0011 
	/* If we perform bitwise& operation here  
	10    = 1 0 1 0
		    & & & &
	3     = 0 0 1 1 
	result= 0 0 1 0 = 2
	*/
	fmt.Println(x)

	y:=10
	// | (or) Now you can guess
	y |= 3
	/* If we perform bitwise& operation here  
	10    = 1 0 1 0
		    | | | |
	3     = 0 0 1 1 
	result= 1 0 1 1 = 
	*/
	fmt.Println(y)

	var x2, y2 = 15, 15

    x2 = x2 >> 3
    y2 = y2 << 3
    fmt.Print(x2, y2) // Output?
}