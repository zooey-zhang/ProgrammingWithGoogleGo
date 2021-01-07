package main

import ("fmt"
)

func main() {
	var cnt int;
	fmt.Println("Please enter the number of the integer")
	fmt.Scan(&cnt)
	var arr []int
	arr = make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&arr[i])
	}
	BubbleSort(arr)
	for i := 0; i < cnt; i++ {
		fmt.Println(arr[i])
	}
}

func BubbleSort(slice []int) {
	for true {
		swap := true;
		for i := 1; i < len(slice); i++ {
			if (slice[i] < slice[i-1]) {
				Swap(i, slice)
				swap = false
			}

		}
		if swap {
			break
		}

	}

}

func Swap(index int, slice []int) {
	temp := slice [index - 1]
	slice[index - 1] = slice[index]
	slice[index] = temp

}
