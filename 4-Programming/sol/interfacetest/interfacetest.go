package interfacetest

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	fmt.Println("Basic Interface============================================")
	err := getSpecialError()

	fmt.Println(err.Error())

	if specialErr, ok := errors.AsType[specialError](err); ok {
		fmt.Println(specialErr.developerMsg)
	}
}
