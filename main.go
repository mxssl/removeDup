package main

import "fmt"

func main() {
	arr := []int{10, 5, 1, 16, 10, 1, 16}
	normalizedArr := removeDuplicate(arr)
	fmt.Println(normalizedArr)
}

func removeDuplicate(arr []int) []int {
	tempMap := make(map[int]bool)
	result := []int{}

	for _, value := range arr {
		if !tempMap[value] {
			tempMap[value] = true
			result = append(result, value)
		}
	}

	return result
}
