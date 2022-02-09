/*
Write a program to sort an array of integers. The program should partition the
array into 4 parts, each of which is sorted by a different goroutine. Each
partition should be of approximately equal size. Then the main goroutine should
merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine
which sorts ¼ of the array should print the subarray that it will sort. When
sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	input := getInput()
	chunckedSlices := chunkSlice(input, 4)
	c := make(chan []int, 4)

	go sortSlice(0, chunckedSlices[0], c)
	go sortSlice(1, chunckedSlices[1], c)
	go sortSlice(2, chunckedSlices[2], c)
	go sortSlice(3, chunckedSlices[3], c)
	arr1 := make([]int, 0, len(input)/2)
	arr2 := make([]int, 0, len(input)/2)
	finalArr := make([]int, 0)
	merge(<-c, <-c, &arr1)
	merge(<-c, <-c, &arr2)
	merge(arr1, arr2, &finalArr)
	fmt.Println("Final sorted output", finalArr)
}

func merge(l []int, r []int, ans *[]int) {
	if len(l) == 0 && len(r) == 0 {
		fmt.Println("Sub-merge result", ans)
		return
	}
	if len(l) == 0 {
		*ans = append(*ans, r[0])
		merge(l, r[1:], ans)
		return
	}
	if len(r) == 0 {
		*ans = append(*ans, l[0])
		merge(l[1:], r, ans)
		return
	}
	if l[0] < r[0] {
		*ans = append(*ans, l[0])
		merge(l[1:], r, ans)
		return
	} else {
		*ans = append(*ans, r[0])
		merge(l, r[1:], ans)
		return
	}
}

func sortSlice(id int, ints []int, c chan []int) {
	fmt.Println("Unsorted goroutine ", id, ints)
	sort.Ints(ints)
	fmt.Println("Sorted goroutine", id, ints)
	c <- ints
}

func getInput() []int {
	res := make([]int, 0, 10)
	fmt.Print("type integers to sort (end with x): ")
	var i int
	for {
		_, err := fmt.Scan(&i)
		if err != nil {
			break
		}
		res = append(res, i)
	}
	fmt.Println("Array to be sorted is", res)
	return res
}

func chunkSlice(slice []int, partitionSize int) [][]int {
	var chunks [][]int
	for i := 0; i < partitionSize; i++ {
		start := i * len(slice) / partitionSize
		end := (i + 1) * len(slice) / partitionSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[start:end])
	}

	return chunks
}
