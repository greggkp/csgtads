package main

import "log"

func insertionSort(list []int) {

	for i := 1; i < len(list); i++ {
		pos := i
		tempVal := list[i]

		for ; pos > 0 && tempVal < list[pos-1]; pos-- {
			list[pos] = list[pos-1]
		}

		list[pos] = tempVal
	}

}

func main() {
	list := []int{65, 55, 45, 35, 25, 15, 10}
	log.Println(list)
	insertionSort(list)
	log.Println(list)
}
