package interfaces

import "fmt"

type specialError struct {
	StatusCode   int
	Message      string
	developerMsg string
}

func (r specialError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Message)
}

func getSpecialError() error {
	err := specialError{StatusCode: 100, Message: "You failed", developerMsg: "You failed because the answer was 24."}
	return err
}
