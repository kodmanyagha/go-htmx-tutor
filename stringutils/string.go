package stringutils

import (
	"fmt"

	"go-htmx-tutor/mathutils"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func GreetAndAdd(name string, a int, b int) {
	greeting := Greet(name)
	fmt.Println(greeting)
	mathutils.PrintSum(a, b) // Call PrintSum from mathutils
}
