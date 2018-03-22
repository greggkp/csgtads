package main

import "log"

func selectionSort(list []int) {

	for i := 0; i < len(list); i++ {
		lowestValIndex := i;
		for j := i+1; j < len(list); j++ {
			if list[j] < list[lowestValIndex] {
				lowestValIndex = j
			}
		}
		if lowestValIndex != i {
			temp := list[i]
			list[i] = list[lowestValIndex]
			list[lowestValIndex] = temp
		}
	}
}

func main() {
	list := []int{65, 55, 45, 35, 25, 15, 10}
	log.Println(list)
	selectionSort(list)
	log.Println(list)
}
