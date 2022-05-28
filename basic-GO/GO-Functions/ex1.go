package main
import ("fmt")

func myMessage(){
	fmt.Println("I just go executed")
}

func main(){
	for i:=0; i<3; i++{
		myMessage()
	}
	
}