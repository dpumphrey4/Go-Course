package maptest

import (
	"fmt"
)

// Insert/Update	m[key] = value	Adds or updates an entry.
// Get Value	val := m[key]	Returns the value or zero value if missing.
// Check Existence	val, ok := m[key]	ok is true if the key exists.
// Delete	delete(m, key)	Removes a key; safe if key doesn't exist.
// Length	len(m)	Returns the number of items in the map.

func MakeMap() map[string]uint32 {
	x := make(map[string]uint32)
	x["gophers"] = 0xd3adb33f
	x["blue"] = 1234
	x["channels"] = 3
	x["go-routines"] = 5

	printMap(x)
	return x
}

func AddToMap(key string, value uint32, myMap map[string]uint32) {
	myMap[key] = value
	printMap(myMap)
}

func GetFromMap(key string, myMap map[string]uint32) (uint32, error) {
	// If you don't check for the existance, it'll give a default value like 0
	val, ok := myMap[key]
	if ok == true {
		return val, nil
	}

	return val, fmt.Errorf("Failed to find %s in the map.", key)
}

func printMap(x map[string]uint32) {
	fmt.Printf("My map: ")
	for s, i := range x {
		fmt.Printf("%s: %d\n", s, i)
	}
}
