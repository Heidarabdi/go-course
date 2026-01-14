package main
import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulus:", a%b)

	a += 5
	fmt.Println("Addition Assignment:", a)

	fmt.Println("Equal to:", a == b)
	fmt.Println("Not equal to:", a != b)
	fmt.Println("Greater than:", a > b)
	fmt.Println("Less than:", a < b)
	fmt.Println("Greater than or equal to:", a >= b)
	fmt.Println("Less than or equal to:", a <= b)

	fmt.Println("Logical AND:", a > 5 && b < 5)
	fmt.Println("Logical OR:", a > 5 || b > 5)
	fmt.Println("Logical NOT:", !(a > 5))

	fmt.Println("Bitwise AND:", a & b)
	fmt.Println("Bitwise OR:", a | b)
	fmt.Println("Bitwise XOR:", a ^ b)
	fmt.Println("Bitwise AND NOT:", a &^ b)
	fmt.Println("Left Shift:", a << 1)
	fmt.Println("Right Shift:", a >> 1)

	ptr := &a
	fmt.Println("Address of a:", ptr)
	fmt.Println("Value at address:", *ptr)
}