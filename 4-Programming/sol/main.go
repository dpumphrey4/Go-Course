package main

import (
	"fmt"
	"programming/fileoperations"
	"programming/gophers"
	"programming/interfaces"
	"programming/maps"
	"programming/slices"
	"programming/structs"
)

func basicSlices() {
	fmt.Println("Basic Slices===========================================")
	slices.SlicesOnSlices1()
	slices.SlicesOnSlices2()
	slices.CloneSlice()
	slices.SlicesPanic()
	slices.SlicePointer()
}

func basicMaps() {
	fmt.Println("Basic Maps=============================================")
	myMap := maps.MakeMap()
	maps.AddToMap("slices", 42, myMap)
	myValue, err := maps.GetFromMap("slices", myMap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("This is the value I received: %s = %d\n\n", "slices", myValue)
	}

}

func writingFiles() {
	fmt.Println("Basic Files============================================")
	fileoperations.CreateFile("data.txt")
	fileoperations.WriteFile("data2.txt", []byte("This is my message, I love gophers."))
	fileoperations.Writing("data3.txt", []byte("This is my message, I love gophers.\n"), "So do I.")
	fileoperations.ReadFile("data3.txt")
	fileoperations.RemoveFiles([]string{"data.txt", "data2.txt", "data3.txt"})
}

func main() {
	basicSlices()
	basicMaps()
	writingFiles()
	structs.StructTest()
	interfaces.ErrorTest()
	gophers.MainGopher()
}
