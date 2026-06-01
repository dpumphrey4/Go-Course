package main

import (
	"fmt"
	"programming/filetest"
	"programming/gopherstest"
	"programming/maptest"
	"programming/slicetest"
	"programming/structtest"
)

func basicSlices() {
	fmt.Println("Basic Slices===========================================")
	slicetest.SlicesOnSlices1()
	slicetest.SlicesOnSlices2()
	slicetest.CloneSlice()
	slicetest.SlicesPanic()
}

func basicMaps() {
	fmt.Println("Basic Maps=============================================")
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
	fmt.Println("Basic Files============================================")
	filetest.CreateFile("data.txt")
	filetest.WriteFile("data2.txt", []byte("This is my message, I love gophers."))
	filetest.Writing("data3.txt", []byte("This is my message, I love gophers.\n"), "So do I.")
	filetest.ReadFile("data3.txt")
	filetest.RemoveFiles([]string{"data.txt", "data2.txt", "data3.txt"})
}

func main() {
	basicSlices()
	basicMaps()
	writingFiles()
	structtest.ErrorTest()
	gopherstest.MainGopher()
}
