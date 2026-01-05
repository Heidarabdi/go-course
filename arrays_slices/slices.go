package main
import "fmt"

func main(){
	// Creating slices using different methods
	slice1 := []int{1, 2, 3} // creating a slice using the []datatype{values} format
	fmt.Println("Slice1:", slice1)

	// Creating a slice using the make function
	slice2 := make([]int, 5) // creating a slice using the make function with length 5
	fmt.Println("Slice2 (made with make):", slice2)

	// Creating a slice by slicing an existing array
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[1:4] // creating a slice by slicing an existing array
	fmt.Println("Slice3 (from array slice):", slice3)

	// Accessing and modifying slice elements
	firstElement := slice1[0] // accessing the first element of the slice
	fmt.Println("First element of slice1:", firstElement)
	slice1[1] = 10 // modifying the second element of the slice
	fmt.Println("Modified slice1:", slice1)

	// Appending elements to a slice
	slice1 = append(slice1, 4, 5) // appending elements to the slice it becomes [1, 10, 3, 4, 5]
	fmt.Println("Slice1 after append:", slice1)

	// Appending one slice to another
	slice2_new := []int{6, 7, 8}
	slice1 = append(slice1, slice2_new...) // appending slice2_new to slice1 it becomes [1, 10, 3, 4, 5, 6, 7, 8]
	fmt.Println("Slice1 after appending slice2_new:", slice1)

	// Length and capacity of a slice
	slice := []int{1, 2, 3}
	lengthOfSlice := len(slice) // lengthOfSlice is 3
	capacityOfSlice := cap(slice) // capacityOfSlice is 3
	fmt.Println("Slice:", slice)
	fmt.Println("Length of slice:", lengthOfSlice)
	fmt.Println("Capacity of slice:", capacityOfSlice)

	// Slicing further to create new slices
	slice1_example := []int{1, 2, 3, 4, 5}
	slice2_example := slice1_example[1:4] // creating a new slice from slice1_example it becomes [2, 3, 4]
	fmt.Println("Slice2_example (from slice1_example[1:4]):", slice2_example)

	// Changing the length of the slice when slicing
	slice3_example := slice1_example[2:] // creating a new slice from slice1_example it becomes [3, 4, 5]
	fmt.Println("Slice3_example (from slice1_example[2:]):", slice3_example)

	slice4_example := slice1_example[:3] // creating a new slice from slice1_example it becomes [1, 2, 3]
	fmt.Println("Slice4_example (from slice1_example[:3]):", slice4_example)
}