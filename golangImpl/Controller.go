//William Murray
//John Miner
package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//In go, only the attributes need to be defined in the body of the type(class) declaration
//so in this case it has an empty body
type Controller struct{}

//Methods bound to a type are done so with the prefix operators
//   |<---------->|
func (C Controller) Menu() {
	fmt.Print("MENU \n\n")
	fmt.Println("1) Enter your own array to test the 4 methodologies")
	fmt.Println("2) Enter a size, to create an array of that size to test select methods")
	fmt.Println("3) Select a method to estimate the runtime for different array sizes")
	fmt.Println("4) Quit")
}

//Method takes in the choice of the user as an int, and returns the specified func from the
//map (hashtable,dictionary,etc)
func (C Controller) choiceTaker(x int) func() {

	m := map[int]func(){
		1: option1,
		2: option2,
		3: option3,
	}
	return m[x]
}

//generates a random int between the min and max provided
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//creates a random array via the arguement array, and size
func randArray(ar []int, i int) {
	for j := 0; j < i; j++ {
		ar[j] = randInt(-50, 50)
	}
}

//another return map to give the specified measure func back
func measureTaker(x int) func([]int) time.Duration {
	m := map[int]func([]int) time.Duration{
		1: measure1,
		2: measure2,
		3: measure3,
		4: measure4,
	}
	return m[x]
}

/*
first option of the project
User enters in an array in the format of ints followed by ','
given the array, the program will run each of solutions for the given array and return the
Maximum Sub-Array Sum
*/
func option1() {
	fmt.Println("Input elements of your array followed by ',' i.e. : 1,2,...\n:")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the string has failed")
	}
	//the read in string will contain a \n so truncate that from the string
	reply = reply[:len(reply)-1]
	//devy up the string based on the commas
	strs := strings.Split(reply, ",")
	//build the array from the entered values
	ar := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		ar[i], _ = strconv.Atoi(strs[i])
	}
	//run each of the given solution with the entered array
	fmt.Println("Solution 1: ", solution1(ar))
	fmt.Println("Solution 2: ", solution2(ar))
	left := 0
	right := len(ar) - 1
	fmt.Println("Solution 3: ", solution3(ar, left, right))
	fmt.Println("Solution 4: ", solution4(ar))
}

/*
Second option of the project
User enters a size, of which an array of random numbers will be made from
Then user selects which solution they want to test on the array
	Solutions are taken as a string i.e. ("124","34", etc) where in ommitions are to be expected
Program will then run the selected solutions on the random array, displaying the result and time taken to run
*/
func option2() {
	fmt.Print("Enter the size of the array to test\n:")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	//Create the space for the array in memory
	ar := make([]int, i)

	//populate the array
	randArray(ar, i)

	fmt.Print("Enter each Method to test i.e.(1234)\n:")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the string has failed")
	}
	//again the newline is annoying so get rid of it
	reply = reply[:len(reply)-1]
	//check to see which of the methods has been selected!
	if strings.Contains(reply, "1") {
		for index := 0; index < 10; index++ {
			measure1(ar)
		}
	}
	if strings.Contains(reply, "2") {
		measure2(ar)
	}
	if strings.Contains(reply, "3") {
		measure3(ar)
	}
	if strings.Contains(reply, "4") {
		measure4(ar)
	}

}

/*
Option 3 of the project
User will select one solution to test
Program will ask for a baseline array size, and an experiment array size
Program will run with the baseline size, and using the runtime of the baseline, estimate the experiment size
After which the program will run the experiment size, and the user can compare the estimate to the actual
*/
func option3() {
	fmt.Print("Which method would you like to test\n:")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	methodNum := i
	//here the method returned from the measureTaker method is stored in method, and can be called at a later time
	method := measureTaker(methodNum)

	fmt.Print("Enter the baseline array size\n:")
	_, err = fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	//create array based on the baseline size
	m := i
	ar := make([]int, m)
	randArray(ar, m)

	fmt.Print("Enter an array size to estimate\n:")
	_, err = fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}

	//After we got both array sizes we can run the method we stored and pass in the array
	//store the runtime of the baseline in diff
	diff := method(ar)

	//create array based on the experimental size
	n := i
	ar = make([]int, n)
	randArray(ar, n)

	//create var to hold the float val of the estimated runtime
	estimate := float64(0)

	switch methodNum {
	case 1:
		estimate = calculate1(m, n, diff)
		break
	case 2:
		estimate = calculate2(m, n, diff)
		break
	case 3:
		estimate = calculate3(m, n, diff)
		break
	case 4:
		estimate = calculate4(m, n, diff)
		break
	default:
		estimate = 0
	}

	fmt.Print("\nThe expectied time for Method #", methodNum, " at array size ", n, " ")

	//estimate is based on nanoseconds, thus simple calculations to refactor the prefix
	if estimate/math.Pow(10.0, 9.0) > 1 {
		fmt.Printf("is %.2f (s)", estimate/math.Pow(10.0, 9.0))
	} else if estimate/math.Pow(10.0, 6.0) > 1 {
		fmt.Printf("is %.2f (ms)", estimate/math.Pow(10.0, 6.0))
	} else if estimate/math.Pow(10.0, 3.0) > 1 {
		fmt.Printf("is %.2f (µs)", estimate/math.Pow(10.0, 3.0))
	} else {
		fmt.Printf("is %.2f (ns)", estimate/math.Pow(10.0, 3.0))
	}
	fmt.Print("\n\n")

	method(ar)

}

//first solution
//basic implimentation to derive the MSS
//triple nested loops means O(n^3) runtime
//very inefficient
func solution1(ar []int) int {
	max_sum := 0

	//start at beggining of array
	for i := 0; i < len(ar); i++ {
		//set second loop at first loops iterator
		for j := i; j < len(ar); j++ {
			//set current sum to 0
			this_sum := 0
			//set third loop to first iterator, stop when passing second iterator
			for k := i; k <= j; k++ {
				this_sum += ar[k]
			}
			//if the current sum is greater than the max, then replace
			if this_sum > max_sum {
				max_sum = this_sum
			}
		}
	}
	return max_sum
}

//second implimentation
//only two nested loops for O(n^2) runtime
//refactored out the unecessary looping of the third loop, basically used second loop as a stopping point
//and had the third loop do calulations, thus deprecated in this implimentation
func solution2(ar []int) int {
	max_sum := 0

	for i := 0; i < len(ar); i++ {
		this_sum := 0
		for j := i; j < len(ar); j++ {
			this_sum += ar[j]
			if this_sum > max_sum {
				max_sum = this_sum
			}
		}
	}
	return max_sum
}

//divide and conquer implimentation
//Using recursive approach
//recursive methods give time complexity of O(nlogn) runtime
// l and r are indexes in this case
func solution3(ar []int, l int, r int) int {
	//if the indexes are the same return the value in the index
	if r == l {
		return ar[l]
	}

	//this case, l and r are right next to each other,
	//so return either the larger of the two, or both if their sum is largest
	if r == l+1 {
		return int(math.Max(float64(ar[l]), math.Max(float64(ar[r]), float64(ar[l]+ar[r]))))
	}

	//neither of the common cases are present so divide the array by integer division
	m := (l + r) / 2

	//calculate the maximum sum for either side of the array
	msl := solution3(ar, l, m)
	msr := solution3(ar, m+1, r)

	//calculate the sum for the middle case of the array
	msm := solution3_mid(ar, l, m, r)

	//return the max of each of the array sections
	return int(math.Max(float64(msl), math.Max(float64(msr), float64(msm))))
}

//sub recursive step to calculate the maximum sum across the left and right sub arrays
func solution3_mid(ar []int, left int, mid int, right int) int {
	//set baseline max to an impossibly small value, such that nothing can be smaller
	//as we are summing possibly negative numbers in sequence, giving a extremely small
	//negative num can be overwritten, if the array size is large enough
	max_left := int(math.Inf(-1))
	sum := 0
	//iterate from the middle outwards to collect the largest left half of the sum
	for i := mid; i >= left; i-- {
		sum += ar[i]

		if sum > max_left {
			max_left = sum
		}
	}
	// '                                                                               '
	max_right := int(math.Inf(-1))
	sum = 0
	//iterate from the middle outwards to the right to find largest right hand sum
	for i := mid + 1; i <= right; i++ {
		sum += ar[i]

		if sum > max_right {
			max_right = sum
		}
	}
	// return the sum of the max sub array from the left and the right, both echo out from the middle, so
	//this sum is effectively the total sum of the array inbetween left and right halves
	return max_right + max_left
}

//simple forloop that deprecates the current sum when it dips below 0
//Single for loop yields O(n) runtime
func solution4(ar []int) int {
	max_sum := 0
	this_sum := 0

	for i := 0; i < len(ar); i++ {
		this_sum += ar[i]
		if this_sum > max_sum {
			max_sum = this_sum
		} else if this_sum < 0 {
			this_sum = 0
		}
	}

	return max_sum
}

func measure1(ar []int) time.Duration {
	fmt.Println("Method 1 is being measured")
	//fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution1(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
	return diff
}

func measure2(ar []int) time.Duration {
	fmt.Println("Method 2 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution2(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
	return diff
}

func measure3(ar []int) time.Duration {
	l := 0
	r := len(ar) - 1
	fmt.Println("Method 3 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution3(ar, l, r))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
	return diff
}

func measure4(ar []int) time.Duration {
	fmt.Println("Method 4 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution4(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
	return diff
}

func calculate1(m int, n int, diff time.Duration) float64 {
	return float64(diff.Nanoseconds()) * math.Pow(float64(n/m), 3.0)
}

func calculate2(m int, n int, diff time.Duration) float64 {
	return float64(diff.Nanoseconds()) * math.Pow(float64(n/m), 2.0)
}

func calculate3(m int, n int, diff time.Duration) float64 {
	a := float64(diff.Nanoseconds()) / (float64(m) * math.Log(float64(m)))
	return a * float64(n) * math.Log(float64(n))
}

func calculate4(m int, n int, diff time.Duration) float64 {
	return float64(diff.Nanoseconds()) * (float64(n) / float64(m))
}
