package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}

// go run .
// go help 
// go mod tidy - скачивает все зависимости