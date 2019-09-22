package main

import (
	"fmt"
)

func main() {
	// arTester := []int{1, 2, 3, 4}

	C := Controller{}

	C.Menu()

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	for i != 4 {
		method := C.choiceTaker(i)
		method()

		C.Menu()

		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("failed to retrieve data from scanner")
		}
	}

}
