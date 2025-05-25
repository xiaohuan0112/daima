package main

import "fmt"

func main() {
	a := []int{9, 9}
	fmt.Println(num_jia_1(a))

}

func num_jia_1(aaa []int) []int {
	for i := len(aaa) - 1; i >= 0; i-- {
		if aaa[i] < 9 {
			aaa[i]++
			return aaa
		}
		aaa[i] = 0
	}
	return append([]int{1}, aaa[0:]...)
}
