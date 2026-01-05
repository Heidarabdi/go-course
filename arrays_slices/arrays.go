package main
import "fmt"

func main(){
	// Declaring Arrays and Initialization
	var arr = [5]int{1, 2, 3, 4, 5} // declaring and initializing an array
	fmt.Println("Array:", arr)
	colors := [...]string{"red", "green", "blue"} // array with inferred size
	fmt.Println("Colors:", colors)

	// Default values for uninitialized arrays
	var nums [3]int // nums is [0, 0, 0] by default
	fmt.Println("Default nums:", nums)
	nums2 := [3]bool{} // nums is [0, 0, 0] by default
	fmt.Println("Default nums2:", nums2)

	// Partially initialized arrays
	var nums3 = [3]int{5} // nums is [5, 0, 0] its partially initialized
	fmt.Println("Partially initialized nums3:", nums3)
	nums4 := [3]float32{5} // nums is [5, 0, 0] its partially initialized
	fmt.Println("Partially initialized nums4:", nums4)

	// Fully initialized arrays
	var nums5 = [3]int{1, 2, 3} // nums is [1, 2, 3] its fully initialized
	fmt.Println("Fully initialized nums5:", nums5)
	nums6 := [3]int{1, 2, 3} // nums is [1, 2, 3] its fully initialized
	fmt.Println("Fully initialized nums6:", nums6)

	// Arrays initialized with specific values at certain indices
	var nums7 = [3]int{0: 1, 2: 3} // nums is [1, 0, 3] initialized with specific indices
	fmt.Println("Specific indices nums7:", nums7)
	nums8 := [3]int{2:12} // nums is [1, 0, 3] initialized with specific indices
	fmt.Println("Specific indices nums8:", nums8)

	// Accessing Array Elements
	firstElement := arr[0] // accessing the first element
	fmt.Println("First element of arr:", firstElement)
	arr[1] = 10 // modifying the second element
	fmt.Println("Modified arr:", arr)

	// Array Length
	var length = [5]int{1, 2, 3, 4, 5}
	lengthOfArray := len(length) // lengthOfArray is 5
	fmt.Println("Length of length array:", lengthOfArray)

	colors2 := [...]string{"red", "green", "blue"}
	lengthOfColors := len(colors2) // lengthOfColors is 3
	fmt.Println("Length of colors2 array:", lengthOfColors)
}