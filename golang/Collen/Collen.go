package main
import "fmt"
/*
	Collen_go
*/
func test() {}
func main() {
/*
	Collen_go2
*/
c := "package main%[1]cimport %[2]cfmt%[2]c%[1]c/*%[1]c%[3]cCollen_go%[1]c*/%[1]cfunc test() {}%[1]cfunc main() {%[1]c/*%[1]c%[3]cCollen_go2%[1]c*/%[1]cc := %[2]c%[4]s%[2]c%[1]cfmt.Printf(c, 10, 34, 9, c)%[1]c}%[1]c"
fmt.Printf(c, 10, 34, 9, c)
}
