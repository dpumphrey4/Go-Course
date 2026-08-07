package structs

import (
	"fmt"
	"math/rand"
)

type Elf struct {
	Name         string
	Age          uint64
	PresentCount int64
	Retired      bool
}

// This takes in a copy of the elf, so the original copy is not updated.
func (elf Elf) SetRetiredv1(retired bool) {
	elf.Retired = retired
	fmt.Printf("SetRetiredv1: %t\n", elf.Retired)
}

// This takes a pointer to an elf so the original copy is updated.
func (elf *Elf) SetRetiredv2(retired bool) {
	elf.Retired = retired
	fmt.Printf("SetRetiredv2: %t\n", elf.Retired)
}

// Since no change happens here, a copy doesn't affect anything, but it would
// be more efficient to use a pointer.
func (elf Elf) GetRetired() bool {
	return elf.Retired
}

// Prints out the elf's current status.
func (elf *Elf) Print() {
	fmt.Printf("Elf: %s\nAge: %d\nPresents: %d\nRetired: %t\n\n", elf.Name, elf.Age, elf.PresentCount, elf.Retired)
}

// Changes it's local copy of the elf.
func changeElfv1(elf Elf) {
	elf.Age = rand.Uint64()
	elf.Name = "Elfv1"
	elf.PresentCount = 9000
	elf.Print()
}

// Changes the elf by using the pointer.
func changeElfv2(elf *Elf) {
	elf.Age = rand.Uint64()
	elf.Name = "Elfv2"
	elf.PresentCount = 24000
	elf.Print()
}

func StructTest() {
	elf := Elf{
		Name: "OgElf", Age: 20, PresentCount: 200, Retired: false,
	}

	fmt.Println("Elf original state:")
	elf.Print()

	elf.SetRetiredv1(true)
	fmt.Println("Elf after SetRetiredv1")
	elf.Print()

	elf.SetRetiredv2(true)
	fmt.Println("Elf after SetRetiredv2")
	elf.Print()

	changeElfv1(elf)
	fmt.Println("Elf after changeElfv1")
	elf.Print()

	changeElfv2(&elf)
	fmt.Println("Elf after changeElfv2")
	elf.Print()
}
