package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	nm := os.Args[0]
	fmt.Println(nm)
	a := strings.Contains(nm, "_")
	fmt.Println(a)
}
