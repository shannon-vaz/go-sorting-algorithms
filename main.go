package main

import (
	"fmt"
	"math"
)

func mergeSort(array []int, n int, prefix string) {
	fmt.Printf(prefix+"mergeSort(%v, %v)\n", array, n)
	if n > 1 {
		m := int(math.Floor(float64(n / 2)))
		subPrefix := prefix[:len(prefix)-2] + "   |--"
		mergeSort(array[:m], m, subPrefix)
		mergeSort(array[m:], n-m, subPrefix)
		merge(array, n, m, prefix)
	}
}

func merge(array []int, n, m int, prefix string) {
	fmt.Printf(prefix+"merge(%v, %v)\n", array, n)
	temp := make([]int, n)
	i, j := 0, m
	for k := 0; k < n; k++ {
		if j >= n {
			temp[k] = array[i]
			i++
		} else if i >= m {
			temp[k] = array[j]
			j++
		} else if array[i] < array[j] {
			temp[k] = array[i]
			i++
		} else {
			temp[k] = array[j]
			j++
		}
	}
	for index, value := range temp {
		array[index] = value
	}
}

func quickSort(array []int, low, high int, prefix string) {
	fmt.Printf(prefix+"quickSort(%v, %v, %v)\n", array[low:high+1], low, high)
	if low < high {
		subPrefix := prefix[:len(prefix)-2] + "   |--"
		pivotIndex := partition(array, low, high, prefix)
		quickSort(array, low, pivotIndex-1, subPrefix)
		quickSort(array, pivotIndex+1, high, subPrefix)
	}
}

func partition(array []int, low, high int, prefix string) int {
	fmt.Printf(prefix+"partition(%v, %v, %v)\n", array[low:high+1], low, high)
	pivot := array[low]
	i := low + 1
	j := high
	prefix = prefix + "-- "
	for {
		fmt.Printf(prefix+"pivot = %d, array = %v, i = %d, j = %d\n", pivot, array[low:high+1], i, j)
		for ; i <= high && array[i] < pivot; i++ {
		}
		for ; j >= low && array[j] > pivot; j-- {
		}
		if i >= j {
			fmt.Printf(prefix+"break: i = %d, j = %d\n", i, j)
			break
		}
		fmt.Printf(prefix+"swap array[%d] <-> array[%d] - %d <-> %d\n", i, j, array[i], array[j])
		array[i], array[j] = array[j], array[i]
		i, j = i+1, j-1
	}
	fmt.Printf(prefix+"Swap pivot with array[%d] i.e. %d\n", j, array[j])
	array[low], array[j] = array[j], array[low]
	return j
}

func selectionSort(array []int, n int) {
	for i := 0; i < n; i++ {
		j := i
		fmt.Printf("i = %d, array = %v\n", i, array)
		for k := i + 1; k < n; k++ {
			if array[k] < array[j] {
				j = k
			}
		}
		array[i], array[j] = array[j], array[i]
	}
}

func insertionSort(array []int, n int) {
	for i := 1; i < n; i++ {
		fmt.Printf("i = %d, array = %v\n", i, array)
		temp := array[i]
		j := i - 1
		for ; j >= 0 && temp < array[j]; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func readArrayOfInt(array []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("> ")
		fmt.Scanf("%d\n", &array[i])
	}
}

const (
	mergeSortOP = iota
	quickSortOP
	selectionSortOP
	insertionSortOP
)

func main() {
	var array []int
	var size int
	fmt.Println("\nWelcome to sorting fuction")
	fmt.Printf("Enter the size of the array: ")
	fmt.Scanf("%d\n", &size)
	array = make([]int, size)
	fmt.Println("Enter the values of the array:")
	readArrayOfInt(array, size)
	switch choice := getOperation(); choice {
	case mergeSortOP:
		fmt.Println("Tree of calls:")
		mergeSort(array, size, "|--")
	case quickSortOP:
		fmt.Println("Tree of calls:")
		quickSort(array, 0, size-1, "|--")
	case selectionSortOP:
		selectionSort(array, size)
	case insertionSortOP:
		insertionSort(array, size)
	default:
		fmt.Println("Invalid option")
		return
	}
	fmt.Printf("Sorted array: %v", array)
}

func getOperation() int {
	fmt.Println()
	var choice int
	fmt.Println("\nChoose sorting function:")
	fmt.Printf("%d. Merge Sort\n", mergeSortOP)
	fmt.Printf("%d. Quick Sort\n", quickSortOP)
	fmt.Printf("%d. Selection Sort\n", selectionSortOP)
	fmt.Printf("%d. Insertion Sort\n", insertionSortOP)
	fmt.Print("> ")
	fmt.Scanf("%d\n", &choice)
	return choice
}
