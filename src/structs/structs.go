package structs

import (
	"fmt"
	"go-platzi/src/mypackage"
)

// StructsFunc is a function
func StructsFunc() {

	myCar := mypackage.CarPublic{Brand: "Ford", Year: 2020}

	fmt.Println(myCar)

	//Otra manera
	var otherCar mypackage.CarPublic
	otherCar.Brand = "Ferrari"
	fmt.Println(otherCar)
}
