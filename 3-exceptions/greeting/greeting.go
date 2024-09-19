package greeting

import (
    "errors"
    "fmt"
)
func Hello(text string) (string, error) {
	if text == "" {
        return "", errors.New("empty name")
    }
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", text)
    return message, nil
}