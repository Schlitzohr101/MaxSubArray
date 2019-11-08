/*
William Murray
John Miner
CECS 328
Project #1 Testing different algorithms runtimes
*/

package main

import (
	"fmt"
)

//main method to run from
func main() {

	//define controller object
	C := Controller{}

	//view the menu
	C.Menu()

	//define var to take in users input
	var i int
	_, err := fmt.Scanf("%d", &i) //input goes to the address of the defined var
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	//loop for program features
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
