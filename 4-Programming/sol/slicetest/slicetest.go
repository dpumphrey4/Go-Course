package slicetest

import (
	"fmt"
	"slices"
)

// This function will start writing to the same slice, but then once you apeend to the slice
// a new slice will be made
func SlicesOnSlices1() {
	fmt.Println("SlicesOnSlices1==========================================================")

	x := []int{1, 2, 3, 4, 5}
	y := x

	// Can also allocate like so
	// x := make([]int, len)

	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

	fmt.Println("Setting y[1] to 50")
	y[1] = 50

	fmt.Println("")
	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

	fmt.Println("Appending 5435 to y")
	y = append(y, 5435)

	fmt.Println("")
	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

}

// Subslices will still write to the original slice, until it extends beyond the size of the
// original slice, or if you prepend something then it will create a new slice
func SlicesOnSlices2() {
	fmt.Println("SlicesOnSlices2==========================================================")

	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]

	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

	fmt.Println("Setting y[1] to 50")
	y[1] = 50

	fmt.Println("")
	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

	fmt.Println("Appending 5435 to y")
	y = append(y, 5435)

	fmt.Println("")
	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

	fmt.Println("Prepending 5435 to y")
	y = slices.Concat([]int{5435}, y)

	fmt.Println("")
	fmt.Println("This is slice x-")
	printSlice(x)

	fmt.Println("This is slice y-")
	printSlice(y)

}

func CloneSlice() {
	firstSlice := []uint32{7, 8, 9, 0}
	secondSlice := slices.Clone(firstSlice)

	secondSlice = append(secondSlice, 20)
}

func SlicesPanic() {
	fmt.Println("SlicesPanic->This program could die!")
	x := []int{1, 2, 3, 4, 5}

	var index int
	fmt.Print("Enter your index: ")
	fmt.Scanf("%d", &index)

	// To prevent this crash, just compare index to the cap of x
	// if cap(x) < index {
	// 	x[index] = 0xd3ad
	// }

	x[index] = 0xd3ad
	fmt.Printf("Changed %d to %d\n", index, x[index])
}

func printSlice(x []int) {
	for i, v := range x {
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("")
}
