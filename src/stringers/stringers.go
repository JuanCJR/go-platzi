package stringers

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB Ram, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

// StringersFunc is a function
func StringersFunc() {
	myPC := pc{ram: 16, brand: "msi", disk: 200}

	fmt.Println(myPC)
}
