package interfaces

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	fmt.Println("Basic Interface============================================")
	err := getSpecialError()

	fmt.Println(err.Error())

	if specialErr, ok := errors.AsType[specialError](err); ok {
		fmt.Printf("%s\n\n", specialErr.developerMsg)
	}
}
