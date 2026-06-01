package structtest

import (
	"errors"
	"fmt"
)

func ErrorTest() {
	err := getSpecialError()
	if specialErr, ok := errors.AsType[specialError](err); ok {
		fmt.Println(specialErr.developerMsg)
	}
}
