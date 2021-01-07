package main

import ("fmt"
	"strings")

func main() {
	var s string
	fmt.Printf("Please enter a string\n")
	fmt.Scan(&s)
	//for pos := range s {
	if strings.ContainsRune(s, 'a') || strings.ContainsRune(s, 'i') || strings.ContainsRune(s, 'n') {
			fmt.Printf("Found!\n")
	//		return
	//	}
	} else{
	fmt.Printf("Not Fount!\n")
	}
}
