package main

import (
	"fmt"
	"bufio"
	"log"
	"os"

)

func main() {
	file, err := os.Open("./names.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, each_ln := range text {
		fmt.Println(each_ln)

	}

}




