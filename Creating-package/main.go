// main packag
package main

import (
    "fmt"

    "Creating-package/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
	sum := greetings.Add(12,8)
	fmt.Println(sum)
}