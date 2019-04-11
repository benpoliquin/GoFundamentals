package main

import (
	"fmt"
	"os"
)

func main() {
	var output, sep string
	for index := 1; index < len(os.Args); index++ {
		output += sep + os.Args[index]
		sep = " "
	}
	fmt.Println(output)
}
