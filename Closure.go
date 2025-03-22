package main

import "fmt"

func main() {
	// Closure untuk menyimpan counter
	counter := func() func() int {
		count := 1
		return func() int{
			count++
			return count
		}
	}()

	fmt.Print(counter(),",")
	fmt.Print(counter(),",")
	fmt.Print(counter(),",")
}