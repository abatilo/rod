package rod

import "fmt"

const (
	// ErrExpectElement error code
	ErrExpectElement = "expect js to return an element"
	// ErrExpectElements error code
	ErrExpectElements = "expect js to return an array of elements"
)

// Error ...
type Error struct {
	Err     error
	Code    string
	Details interface{}
}

// Error ...
func (e *Error) Error() string {
	return fmt.Sprintf("[rod] %s\n%v", e.Code, e.Details)
}

// Unwrap ...
func (e *Error) Unwrap() error {
	return e.Err
}
