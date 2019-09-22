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

type Controller struct{}

func (C Controller) Menu() {
	fmt.Println("MENU \n")
	fmt.Println("1) Enter your own array to test the 4 methodologies")
	fmt.Println("2) Enter a size, to create an array of that size to test select methods")
	fmt.Println("3) Select a method to estimate the runtime for different array sizes for a single method")
	fmt.Println("4) Quit")
}

func (C Controller) choiceTaker(x int) func() {

	m := map[int]func(){
		1: option1,
		2: option2,
		3: option3,
	}
	return m[x]
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// func measureTaker
// 	}
// }

func option1() {
	fmt.Println("Input elements of your array followed by ',' i.e. : 1,2,...\n:")
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

	// ar := []int{1, -3, 2, 4, -5, 9}

	fmt.Println("Solution 1: ", solution1(ar))
	fmt.Println("Solution 2: ", solution2(ar))
	left := 0
	right := len(ar) - 1
	fmt.Println("Solution 3: ", solution3(ar, left, right))
	fmt.Println("Solution 4: ", solution4(ar))
}

func option2() {
	fmt.Print("Enter the size of the array to test\n:")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("failed to retrieve data from scanner")
	}
	ar := make([]int, i)

	for j := 0; j < i; j++ {
		ar[j] = randInt(-50, 50)
	}

	fmt.Print("Enter each Method to test i.e.(1234)\n:")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the string has failed")
	}
	reply = reply[:len(reply)-1]
	if strings.Contains(reply, "1") {
		measure1(ar)
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
		for j := i; j < len(ar); j++ {
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

func measure1(ar []int) {
	fmt.Println("Method 1 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution1(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
}

func measure2(ar []int) {
	fmt.Println("Method 2 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution2(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
}

func measure3(ar []int) {
	l := 0
	r := len(ar) - 1
	fmt.Println("Method 3 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution3(ar, l, r))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
}

func measure4(ar []int) {
	fmt.Println("Method 4 is being measured")
	fmt.Println("Array to be tested: ", ar)
	t1 := time.Now()
	fmt.Println("MSS found: ", solution4(ar))
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Elapsed time: ", diff)
}
