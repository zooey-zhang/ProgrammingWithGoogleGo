package main

import ("fmt"
	"encoding/json")

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

func main() {
	//fmt.Println(Person{"Alice", "MN"})
	s := Person{Name : "Bob", Address : "LA"}
	//fmt.Println(s)
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
	fmt.Println(string(b))

}


