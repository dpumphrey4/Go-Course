package main

import (
	"fmt"
	"programming/maptest"
	"programming/slicetest"
)

func basicSlices() {
	fmt.Printf("Basic Slices===========================================")
	slicetest.SlicesOnSlices1()
	slicetest.SlicesOnSlices2()
	slicetest.CloneSlice()
	slicetest.SlicesPanic()
}

func basicMaps() {
	fmt.Printf("Basic Maps=============================================")
	myMap := maptest.MakeMap()
	maptest.AddToMap("slices", 42, myMap)
	myValue, err := maptest.GetFromMap("slices", myMap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("This is the value I received: %d\n", myValue)
	}

}

func writingFiles() {

}

func main() {
	basicSlices()
	basicMaps()
	writingFiles()

}
