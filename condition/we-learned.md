# We Learned
- Conditional statements in Go allow you to control the flow of your program based on certain conditions.
## Types of Conditional Statements
- If Statement: The most basic form of conditional statement. It executes a block of code if a specified condition is true.
- If-Else Statement: Extends the if statement to execute one block of code if the condition is true and another block if the condition is false.
- Else-If Ladder: Allows you to check multiple conditions in sequence. The first true condition's block is executed.
- Nested If Statements: You can place an if statement inside another if or else block to create more complex conditions.
- Switch Statement: A multi-way branch statement that allows you to execute one block of code among many based on the value of a variable or expression. it has single case and multiple case options.
## We can use conditional statements like this:

```go
package main
import "fmt"

func main() {
    a := 10
    b := 20

    // If Statement
    if a < b {
        fmt.Println("a is less than b")
    }

    // If-Else Statement
    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is not greater than b")
    }

    // Else-If Ladder
    if a == b {
        fmt.Println("a is equal to b")
    } else if a < b {
        fmt.Println("a is less than b")
    } else {
        fmt.Println("a is greater than b")
    }

    // Nested If Statements
    if a < 15 {
        if b > 15 {
            fmt.Println("a is less than 15 and b is greater than 15")
        }
    }

    // Switch Statement (single case)
    switch a {
    case 5:
        fmt.Println("a is 5")
    case 10:
        fmt.Println("a is 10")
    default:
        fmt.Println("a is neither 5 nor 10")
    }

    // Switch Statement (multiple cases)
    day := 3
    switch day {
    case 1, 2, 3:
        fmt.Println("It's the beginning of the week")
    case 4, 5:
        fmt.Println("It's midweek")
    case 6, 7:
        fmt.Println("It's the weekend")


}
