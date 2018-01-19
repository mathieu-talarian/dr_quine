package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func quine() string {
	return "package main%[1]cimport (%[1]c%[2]c%[3]cfmt%[3]c%[2]c%[1]c%[2]c%[3]cos%[3]c%[2]c%[1]c%[2]c%[3]cos/exec%[3]c%[2]c%[1]c%[2]c%[3]cstrings%[3]c%[2]c%[1]c)%[1]cfunc quine() string {%[1]creturn %[3]c%[4]s%[3]c%[1]c}%[1]cfunc bins(x int) bool {%[1]ccompil := %[3]cgo%[3]c%[1]cargs := []string{%[3]cbuild%2[3]c, %[3]c-o%[3]c, fmt.Sprintf(%[3]cSully_%%d%[3]c, x-1), fmt.Sprintf(%[3]cSully_%%d.go%[3]c, x-1)}%[1]cerr1 := exec.Command(compil, args...).Run()%[1]cif err1 != nil {%[1]creturn false%[1]c}%[1]cerr2 := exec.Command(fmt.Sprintf(%[3]c./Sully_%%d%[3]c, x-1)).Run()%[1]cif err2 != nil {%[1]creturn false%[1]c}%[1]creturn true%[1]c}%[1]cfunc main () {%[1]cvar x = %[5]d%[1]cenv := strings.Split(os.Getenv(%[3]c_%[3]c), %[3]c/%[3]c)%[1]cxx := env[len(env)-1]%[1]cif !strings.Contains(xx, %[3]c_%[3]c) {%[1]cx++%[1]c}%[1]cfileName := fmt.Sprintf(%[3]cSully_%%d.go%[3]c, x-1)%[1]cfile, err := os.Create(fileName)%[1]cif err != nil {%[1]creturn%[1]c}%[1]cfmt.Fprintf(file, quine(), 10, 9, 34, quine(), x-1)%[1]cfile.Close()%[1]cif !bins(x) {%[1]creturn%[1]c}%[1]creturn}"
}

func bins(x int) bool {
	compil := "go"
	args := []string{"run", fmt.Sprintf("./Sully_%d.go", x-1)}
	fmt.Println(args)
	err1 := exec.Command(compil, args...).Start()
	fmt.Println(err1)
	if err1 != nil {
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
	fmt.Fprintf(file, quine(), 10, 9, 34, quine(), x-1)
	file.Close()
	bins(x)
	return
}
