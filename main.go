package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	// Capture argument to form hash
	str := os.Args[1]

	// Instansiate hash and add argument
	hash := sha1.New()
	io.WriteString(hash, str)

	// Create and iterate over slice and test even/odd
	nums := hash.Sum(nil)
	even := nums[:0]
	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, 1)
		} else {
			even = append(even, 0)
		}
	}

	fmt.Printf("%d \n", even)

}
