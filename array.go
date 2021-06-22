package main

import "fmt"

func limitArray() {
	names := [5]string{"one", "two", "three"}
	names[3] = "four"
	names[4] = "five"
	fmt.Println(names)
}

func slice() {
	names := []string{"one", "two", "three"}
	names = append(names, "four", "five", "six")
	fmt.Println(names)
}

func main() {
	limitArray()
	slice()
}
