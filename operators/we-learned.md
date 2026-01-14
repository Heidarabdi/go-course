# We Learned
- Go has various operators to perform operations on variables and values.

## Types of Operators
- Arithmetic Operators: Used for basic mathematical operations like addition (`+`), subtraction (`-`), multiplication (`*`), division (`/`), and modulus (`%`).
- Assignment Operators: Used to assign values to variables, such as `=`, `+=`, `-=`, `*=`, `/=`, and `%=`. For example, `a += 5` is equivalent to `a = a + 5`.
- Comparison Operators: Used to compare two values, including `==` (equal to), `!=` (not equal to), `>` (greater than), `<` (less than), `>=` (greater than or equal to), and `<=` (less than or equal to).
- Logical Operators: Used to combine multiple boolean expressions, such as `&&` (logical AND), `||` (logical OR), and `!` (logical NOT).
- Bitwise Operators: Used to perform bit-level operations on integers, including `&` (AND), `|` (OR), `^` (XOR), `&^` (AND NOT), `<<` (left shift), and `>>` (right shift).
- Other Operators: Include the address operator (`&`) to get the memory address of a variable and the dereference operator (`*`) to access the value at a given memory address.
## Operator Precedence
- Operator precedence determines the order in which operators are evaluated in an expression. For example, multiplication and division have higher precedence than addition and subtraction.
- Parentheses `()` can be used to explicitly specify the order of operations in an expression.
## We can use operators like this:
```go
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