# We learned 
- about arrays in Go, which are fixed-size collections of elements of the same type.
## Declaring Arrays and Initialization
- Arrays can be declared and initialized using the `var` keyword or the short variable declaration syntax. The size of an array can be inferred using the `[...]` syntax. Arrays are useful for storing and manipulating collections of data in a structured way.
- `var arr = [5]int{1, 2, 3, 4, 5}` // declaring and initializing an array
- `colors := [...]string{"red", "green", "blue"}` // array with inferred size
-  If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type like 0 for integers, 0.0 for floats, "" for strings, and false for booleans.
- `var nums [3]int` or `nums := [3]int{}` // nums is [0, 0, 0] by default
- `var nums = [3]int{5}` or `nums := [3]int{5}` // nums is [5, 0, 0] its partially initialized
- `var nums = [3]int{1, 2, 3}` or `nums := [3]int{1, 2, 3}` // nums is [1, 2, 3] its fully initialized
- Arrays can be initialized with specific values at certain indices using index-value pairs.
- `var nums = [3]int{0: 1, 2: 3}` or `nums := [3]int{0: 1, 2: 3}` // nums is [1, 0, 3] initialized with specific indices

## Accessing Array Elements
- Array elements can be accessed using their index, starting from 0. You can read or modify elements by specifying the index in square brackets.
- `firstElement := arr[0]` // accessing the first element
- `arr[1] = 10` // modifying the second element

## Array Length
- The length of an array can be determined using the built-in `len` function.
- `var lenght = [5]int{1, 2, 3, 4, 5}`
- `lengthOfArray := len(lenght)` // lengthOfArray is 5
- `colors := [...]string{"red", "green", "blue"}`
- `lengthOfColors := len(colors)` // lengthOfColors is 3

## Slices
- Slices are dynamic, flexible views into arrays. They can grow and shrink in size and provide more functionality than arrays. Slices are created using three ways: using the []datatype{values} format, using the make function, or by slicing an existing array or slice.
- `slice1 := []int{1, 2, 3}` // creating a slice using the []datatype{values} format
- `slice2 := make([]int, 5)` // creating a slice using the make function with length 5
- `arr := [5]int{1, 2, 3, 4, 5}`
- `slice3 := arr[1:4]` // creating a slice by slicing an existing array
- slices can be accessed and modified using indices, similar to arrays.
- `firstElement := slice1[0]` // accessing the first element of the slice
- `slice1[1] = 10` // modifying the second element of the slice
- Appending elements to a slice can be done using the built-in `append` function, which returns a new slice with the added elements.
- `slice1 = append(slice1, 4, 5)` // appending elements to the slice it becomes [1, 10, 3, 4, 5]
- appeding one slice to another it must use ... after the slice name
- `slice2 := []int{6, 7, 8}`
- `slice1 = append(slice1, slice2...)` // appending slice2 to slice1 it becomes [1, 10, 3, 4, 5, 6, 7, 8]
- The length and capacity of a slice can be determined using the `len` and `cap` functions, respectively.
 - len will give the number of elements in the slice
 - cap will give the total space allocated to the slice because slices are dynamic in nature they can grow and shrink so the capacity is the total space allocated to the slice
- `slice := []int{1, 2, 3}`
- `lengthOfSlice := len(slice)` // lengthOfSlice is 3
- `capacityOfSlice := cap(slice)` // capacityOfSlice is 3
- Slices can be sliced further to create new slices.
- `slice1 := []int{1, 2, 3, 4, 5}`
- `slice2 := slice1[1:4]` // creating a new slice from slice1 it becomes [2, 3, 4]
- changing the length of the slice when slicing
- `slice3 := slice1[2:]` // creating a new slice from slice1 it becomes [3, 4, 5]
- `slice4 := slice1[:3]` // creating a new slice from slice1 it becomes [1, 2, 3]
-
