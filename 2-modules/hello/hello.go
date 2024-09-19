package main

import (
    "fmt"

    "lesson/greeting"
)

func main() {
    // Get a greeting message and print it.
    message := greeting.Hello("Sanya")
    fmt.Println(message)
}