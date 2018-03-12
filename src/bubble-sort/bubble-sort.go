package main

import "log"

func bubbleSort(list []int) {
	unsortedUntilIndex := len(list) - 1
	sorted := false

	for sorted == false {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			if list[i] > list[i+1] {
				sorted = false
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
		unsortedUntilIndex--
	}
}

func main() {
	list := []int{65, 55, 45, 35, 25, 15, 10}
	log.Println(list)
	bubbleSort(list)
	log.Println(list)
}
