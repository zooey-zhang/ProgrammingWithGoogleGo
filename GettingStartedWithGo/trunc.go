package main

import ("fmt")

func main() {
	var x float32 //= 12.34
	fmt.Printf("Enter a float value \n")
	fmt.Scanf("%f", &x)
	fmt.Printf("%f\n", x)
	var y int = int(x)
	fmt.Printf("%d", y)
}



