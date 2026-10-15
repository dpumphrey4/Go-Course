package fileoperations

import (
	"fmt"
	"os"
)

// Demos creating a file object
func CreateFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		// This will close the program, but we do not want to do that here
		// log.Fatal(err)
		return
	}

	// Will always close the file
	defer f.Close()

	readFileObj(f)
}

// Writes to a file with one call
func WriteFile(filename string, data []byte) {
	fmt.Printf("Writing to a file: %s\n\n", filename)

	// Single line to write to a file, just provide permissions
	err := os.WriteFile(filename, data, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Should print our file content
	ReadFile(filename)
}

func Writing(filename string, dataBytes []byte, dataStr string) {
	// Creating the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fmt.Printf("Writing to an open file: %s\n\n", filename)

	// Writing bytes and a string to the file
	f.Write(dataBytes)
	f.WriteString(dataStr)

	// This might print out EOF and no content because writing to the file was buffered
	// Think of it like C's printf, sometimes you have to fflush stdout for the print to run
	// immediately
	readFileObj(f)
}

// Reads an entire file at once.
func ReadFile(filename string) {
	fmt.Printf("Reading an entire file at once: %s\n", filename)

	// This will open, read and close the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert bytes to string
	fmt.Println(string(data))

	fmt.Println()
}

// Removing files in the list
func RemoveFiles(files []string) {
	for _, v := range files {
		os.Remove(v)
	}
}

// Reading a file from a currently open file
func readFileObj(file *os.File) {
	fmt.Printf("Reading from open file: %s\n", file.Name())

	buf := make([]byte, 1000)

	_, err := file.Read(buf)
	if err != nil {
		fmt.Printf("%s\n\n", err)
		return
	}

	fmt.Println(buf)
	fmt.Println("")

}
