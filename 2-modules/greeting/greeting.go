package greeting

import "fmt"

func Hello(text string) string {
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", text)
    return message
}