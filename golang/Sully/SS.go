package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Create("file.go")
	if err != nil {
		return
	}
	c := "package main%[1]cimport %[2]cos%[2]c%[1]cfunc main() {os.Create(%[2]ctest.txt%[2]c)}"
	fmt.Fprintf(file, c, 10, 34)
	file.Close()
	args := []string{"build", "file.go"}
	err1 := exec.Command("go", args...).Run()
	if err1 != nil {
		log.Fatal("err1 ", err1)
		return
	}
	err2 := exec.Command("./file").Start()
	if err2 != nil {
		log.Fatal("err2 ", err2)
		return
	}
}
