package main
import("fmt")

func familyName(fname string, age int){
	fmt.Println("Nama:", fname, "Rael", "Usia:", age )
}

func main(){
	fname := [2]string{"Kirana", "Ranum"}
	age := [2]int{4,1}
	for i:= 0; i<len(fname); i++{
		familyName(fname[i],age[i])
	}
}