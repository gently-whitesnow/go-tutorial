package greeting

import (
    "errors"
    "fmt"
    "math/rand"
)
func Hello(text string) (string, error) {
	if text == "" {
        return "", errors.New("empty name")
    }
	// var message string
	message := fmt.Sprintf(randomFormat(), text)
    return message, nil
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}