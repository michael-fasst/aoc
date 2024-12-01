package main

import (
	"fmt"
	"sort"
)

func main() {

	if len(left_list) != len(right_list) {
		fmt.Println("Error: the two lists are not the same size")
		return
	}

	//////PART ONE/////
	sort.Ints(left_list)
	sort.Ints(right_list)

	var total int

	for i := 0; i < len(left_list); i++ {
		value_left := left_list[i]
		value_right := right_list[i]
		difference := value_left - value_right

		if difference < 0 {
			difference = -difference
		}
		total += difference
	}

	fmt.Println("Result part 1 =", total)
	////////////////////

	//////PART TWO/////
	var total_part_two int

	for i := 0; i < len(left_list); i++ {
		number_of_repeat := 0
		value_left := left_list[i]
		for j := 0; j < len(right_list); j++ {
			value_right := right_list[j]
			if value_left == value_right {
				number_of_repeat++
			}
		}
		total_part_two += number_of_repeat * value_left
	}

	fmt.Println("Result part 2 =", total_part_two)
	////////////////////
}
