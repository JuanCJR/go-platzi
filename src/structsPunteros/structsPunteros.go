package structsPunteros

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

// StructsPunteros is a function
func StructsPunteros() {

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()
	myPC.duplicateRam()
	fmt.Println(myPC)
}
