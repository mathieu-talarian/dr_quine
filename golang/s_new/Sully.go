package main
import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)
func main() {
	var x = 5
	if x <= 0 {return}
	env := strings.Split(os.Getenv("_"), "/")
	xx := env[len(env)-1]
	if !strings.Contains(xx, "_") { x++ }
	file, err := os.Create(fmt.Sprintf("Sully_%d.go", x-1))
	if err != nil {return}
	c := "package main%[1]cimport (%[1]c%[2]cfmt%[2]c%[1]c%[2]cos%[2]c%[1]c%[2]cos/exec%[2]c%[1]c%[2]cstrings%[2]c%[1]c)%[1]cfunc main() {%[1]cvar x = %[4]d%[1]cif x <= 0 {return}%[1]cenv := strings.Split(os.Getenv(%[2]c_%[2]c), %[2]c/%[2]c)%[1]cxx := env[len(env)-1]%[1]cif !strings.Contains(xx, %[2]c_%[2]c) { x++ }%[1]cfile, err := os.Create(fmt.Sprintf(%[2]cSully_%%d.go%[2]c, x-1))%[1]cif err != nil {return}%[1]cc := %[2]c%[3]s%[2]c%[1]cfmt.Fprintf(file, c, 10, 34, c, x-1)%[1]cfile.Close()%[1]cexec.Command(%[2]cgo%[2]c, []string{%[2]cbuild%[2]c, fmt.Sprintf(%[2]cSully_%%d.go%[2]c, x-1)}...).Run()%[1]cexec.Command(fmt.Sprintf(%[2]c./Sully_%%d%[2]c, x-1)).Run()%[1]creturn%[1]c}%[1]c"
	fmt.Fprintf(file, c, 10, 34, c, x-1)
	file.Close()
	exec.Command("go", []string{"build", fmt.Sprintf("Sully_%d.go", x-1)}...).Run()
	exec.Command(fmt.Sprintf("./Sully_%d", x-1)).Run()
	return
}
