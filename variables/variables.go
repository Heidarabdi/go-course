package main
import "fmt"

// var and conts is needed when creating variables it can be used inside or outside
var stdName string
const PI = 3.14 // constant vars cannot be changed

func main(){
	var age int = 30
	height := 5.9 // short variable declaration
	var a, b, c int = 1, 2, 3 // multiple variable declaration
	var i,k,f = "Hello", 2.5, true // multiple variable declaration with different types
	l, m, n := "Go", 42, false // multiple short variable declaration
	var (
		x int = 10
		y int = 20
		z int = 30
	)
	fmt.Println("Age:", age, "Height:", height, "A:", a, "B:", b, "C:", c, "L:", l, "M:", m, "N:", n, "X:", x, "Y:", y, "Z:", z)
	fmt.Printf("PI: %.2f\n", PI)
	fmt.Printf("I: %s, K: %.1f, F: %t\n", i, k, f) // formatted printing
	stdName = "John Doe" // assigning value to variable
	fmt.Println("Student Name:", stdName)
	fmt.Print(PI) // Print doesn't add a new line after printing
}
