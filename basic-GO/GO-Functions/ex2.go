package main
import ("fmt")

func familyName(fname string){
	fmt.Println("hello", fname, "Rael")
}

func main(){
	fname := [2]string{"Kirana", "Ranum"}
	for i:=0; i<len(fname); i++ {
		familyName(fname[i])
	}
}