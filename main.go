// main.go
package main

import (
	"fmt"

	"github.com/shaileshyadav7771/mathutil"
)

func main() {
	a, b := 10, 5

	sum := mathutil.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	diff := mathutil.Subtract(a, b)
	fmt.Printf("%d - %d = %d\n", a, b, diff)

	prod := mathutil.Multiply(a, b)
	fmt.Printf("%d * %d = %d\n", a, b, prod)
}
