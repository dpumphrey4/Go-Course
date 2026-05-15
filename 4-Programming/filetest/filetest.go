package filetest

import (
	"fmt"
	"os"
)

func CreateFile(filename string) {
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		// This will close the program, but we do not want to do that here
		// log.Fatal(err)
	}

	// Will always close the file
	defer f.Close()
}
