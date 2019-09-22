package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// arTester := []int{1, 2, 3, 4}

	fmt.Println("1) Enter your own array to test the 4 methodologies")
	fmt.Println("2) Enter a size, to create an array of that size to test select methods")
	fmt.Println("3) Select a method to estimate the runtime for different array sizes for a single method")
	fmt.Println("4) Quit")

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	for i != 4 {
		method := choiceTaker(i)
		method()

		fmt.Println("1) Enter your own array to test the 4 methodologies")
		fmt.Println("2) Enter a size, to create an array of that size to test select methods")
		fmt.Println("3) Select a method to estimate the runtime for different array sizes for a single method")
		fmt.Println("4) Quit")

		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("failed to retrieve data from scanner")
		}
	}

}

func choiceTaker(x int) func() {

	m := map[int]func(){
		1: option1,
		2: option2,
		3: option3,
	}
	return m[x]
}

func option1() {
	fmt.Println("Enter the array seperated by commas:")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the string has failed")
	}
	reply = reply[:len(reply)-1]
	strs := strings.Split(reply, ",")
	ar := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		ar[i], _ = strconv.Atoi(strs[i])
	}

	fmt.Println("Solution 1: ", solution1(ar))
	fmt.Println("Solution 2: ", solution2(ar))
	left := 0
	right := len(ar) - 1
	fmt.Println("Solution 3: ", solution3(ar, left, right))
	fmt.Println("Solution 4: ", solution4(ar))
}

func option2() {
	fmt.Println("We detected a two")
}

func option3() {

}

func solution1(ar []int) int {
	max_sum := 0

	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			this_sum := 0
			for k := i; k <= j; k++ {
				this_sum += ar[k]
			}
			if this_sum > max_sum {
				max_sum = this_sum
			}
		}
	}
	return max_sum
}

func solution2(ar []int) int {
	max_sum := 0
	for i := 0; i < len(ar); i++ {
		this_sum := 0
		for j := 0; j < len(ar); j++ {
			this_sum += ar[j]
			if this_sum > max_sum {
				max_sum = this_sum
			}
		}
	}

	return max_sum
}

func solution3(ar []int, l int, r int) int {
	if r == l {
		return ar[l]
	}

	if r == l+1 {
		return int(math.Max(float64(ar[l]), math.Max(float64(ar[r]), float64(ar[l]+ar[r]))))
	}

	m := (l + r) / 2

	msl := solution3(ar, l, m)
	msr := solution3(ar, m+1, r)

	msm := solution3_mid(ar, l, m, r)

	return int(math.Max(float64(msl), math.Max(float64(msr), float64(msm))))
}

func solution3_mid(ar []int, left int, mid int, right int) int {
	max_left := int(math.Inf(-1))
	sum := 0

	for i := mid; i >= left; i-- {
		sum += ar[i]

		if sum > max_left {
			max_left = sum
		}
	}

	max_right := int(math.Inf(-1))
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += ar[i]

		if sum > max_right {
			max_right = sum
		}
	}
	return max_right + max_left
}

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
