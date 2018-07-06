package main

import "fmt"
import (
	"goProject/core"
)
func main() {
	fmt.Println("Hello, World!")
	var sum int
	sum = core.Add(2, 4)
	fmt.Print(sum)

}
