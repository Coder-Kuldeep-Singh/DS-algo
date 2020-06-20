package main
import "fmt"

func main() {

	arrays := []int{1,2,232,454,2323,354,5656,434354}
	maximum := arrays[0]

	for _,num := range arrays {
		if num > maximum {
			maximum = num
		}
	}
	fmt.Println(maximum)
}
