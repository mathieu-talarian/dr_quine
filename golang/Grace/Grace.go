package main
import (
"os"
"fmt"
)
/*
	Grace_go
*/
const a = 42
const b = "Hello"
var c = func() {
f, err := os.Create("Grace_kid.go")
if err != nil {
return
}
c := "package main%[1]cimport (%[1]c%[2]cos%[2]c%[1]c%[2]cfmt%[2]c%[1]c)%[1]c/*%[1]c%[3]cGrace_go%[1]c*/%[1]cconst a = 42%[1]cconst b = %[2]cHello%[2]c%[1]cvar c = func() {%[1]cf, err := os.Create(%[2]cGrace_kid.go%[2]c)%[1]cif err != nil {%[1]creturn%[1]c}%[1]cc := %[2]c%[4]s%[2]c%[1]cfmt.Fprintf(f, c, 10, 34, 9, c)%[1]c}%[1]cfunc main() {%[1]cc()%[1]c}%[1]c"
fmt.Fprintf(f, c, 10, 34, 9, c)
}
func main() {
c()
}
