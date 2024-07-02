package service

import (
	"fmt"
	"os"
	"strconv"
)

func FirstFilter() {
	s := os.Args[2]
	n, err := strconv.Atoi(os.Args[3])
	if err != nil || n <= 1 || n%2 == 0 {
		fmt.Println("Error: the number must be an odd integer greater than 1")
		return
	}
	runes := []rune(s)
	length := len(runes)
	for i := n; i >= 1; i -= 2 {
		spaces := (n - i) / 2
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {
			fmt.Print(string(runes[j%length]))
		}
		fmt.Println()
	}
}
