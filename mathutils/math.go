package mathutils

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintSum(a int, b int) {
	sum := Add(a, b)
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, sum)
}
