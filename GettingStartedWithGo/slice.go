package main
import ("fmt"
	"sort")

func main() {
	var cnt int
	fmt.Printf("Please enter the number of the items you would like to enter")
	fmt.Scan(&cnt)
	var arr[]int
	arr = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	for i := range arr {
		fmt.Printf("%d", arr[i])

	}
	

}


