package array

import "fmt"

func main(){
	Array()
}

func Array(){
	var mobilArray = [4]string{
		"BMW",
		"Ferarri",
		"Mercedes Benz",
		"Toyota Supra",
	}

	for i, car:=range mobilArray{
		fmt.Println("Merk Mobil : ", i, car)
	}
}