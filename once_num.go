package main

import "fmt"

func main() {
	a := []int{2, 2, 3, 3, 4, 5}
	fmt.Println(once_num(a))
}

func once_num(ccc []int) int {
	mapp := make(map[int]int)

	for _, v := range ccc {
		mapp[v]++
	}
	for i, v := range mapp {
		if v == 1 {
			return i
		}
	}
	return -1

}
