package main 
import ("fmt")

func testcount(x int) (result int){
	if x == 11{ 
		return 0
	}
	fmt.Println(x)
	result = testcount(x+1)
	return result
}

func main(){
	testcount(1)
}