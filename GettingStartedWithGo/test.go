
package main
import "fmt"

func main() {
	var x int
	var y *int
	z := 3
	y = &z
	x = &y
}
