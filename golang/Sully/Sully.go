package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func quine() string {
	return strings.Join([]string{
		"package main",
		"\n",
		"import (",
		"%[2]c%[3]cfmt%[3]c%[2]c",
	}, "\n")
}

func bins(x int) bool {
	compil := "go"
	args := []string{"build", "-o", fmt.Sprintf("Sully_%d", x-1), fmt.Sprintf("Sully_%d.go", x-1)}
	err1 := exec.Command(compil, args...).Run()
	if err1 != nil {
		return false
	}
	err2 := exec.Command(fmt.Sprintf("./Sully_%d", x-1)).Run()
	if err2 != nil {
		return false
	}
	return true
}

func main() {
	var x = 5
	env := strings.Split(os.Getenv("_"), "/")
	xx := env[len(env)-1]
	if !strings.Contains(xx, "_") {
		x++
	}
	fileName := fmt.Sprintf("Sully_%d.go", x-1)
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	fmt.Fprintf(file, quine(), 10, 9, 34, quine())
	file.Close()
	if !bins(x) {
		return
	}
	return
}
