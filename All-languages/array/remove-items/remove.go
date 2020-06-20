package main

import "fmt"

func main() {
	arrays := []int{1,232,545,12,56,12,10}
	input := 10

	for i,num := range arrays {
		if num == input {
			fmt.Println("value : ",num)
			fmt.Println("Index",i)
			arrays = RemoveIndex(arrays, i)

		}
	}
	fmt.Println(arrays)
}


func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
