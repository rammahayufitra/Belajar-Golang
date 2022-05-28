package main 
import ("fmt")

func myFunction( x int, y string)(result int, text string){
	result = x + x 
	text = y + "world"
	return result, text
}

func main(){
	a,b := myFunction(15, "hello")
	fmt.Println(a,b)
}