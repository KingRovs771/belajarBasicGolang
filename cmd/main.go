package main

import (
	"fmt"
)

func main(){ 

	//basic()
	//For()
	//percabanganIF()
	//Array()
}

func basic(){
	fmt.Println("Hello World")
}


//perulangan
func For(){
	var num1 int=10
	for i:=0;i <= num1;i++{
		fmt.Println(i)
	}
}

//percabanganIF
func percabanganIF(){
	var benar bool = false

	if benar {
		fmt.Println("Hello World")
	}
}

func Array(){
	var mobilArray = [4]string{
		"BMW",
		"Ferarri",
		"Mercedes Benz",
		"Toyota Supra",
	}

	for i, car:=range mobilArray{
		fmt.Println("Merk Mobil ", i ,":", car)
	}
}