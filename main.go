package main

import (
	"fmt"
	"strings"
)

/* naked return */
func lenAndUpper(name string) (length int, uppercase string) {
	/* func finish => defer  */
	defer fmt.Println("Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(lenAndUpper("name"))
}
