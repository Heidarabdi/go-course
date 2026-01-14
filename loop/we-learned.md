# We Learned
- Loops in Go allow you to execute a block of code multiple times based on certain conditions.
## Types of Loops
- For Loop: The only loop construct in Go, which can be used in various forms such as traditional for loops, while-like loops, and infinite loops.
- Continue Statement: Used to skip the current iteration of a loop and move to the next iteration.
- Break Statement: Used to exit a loop prematurely before its normal termination condition is met.
- Nested Loops: You can place one loop inside another loop to create more complex iterations.
- Range Statement: Used to iterate over elements in a data structure, such as arrays, slices, strings, maps, and channels.

## We can use loops like this:

```go
package main
import "fmt"

func main() {
    // Traditional For Loop
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i)
    }

    // While-like Loop
    j := 0
    for j < 5 {
        fmt.Println("While-like Iteration:", j)
        j++
    }

    // Infinite Loop with Break
    k := 0
    for {
        if k >= 5 {
            break
        }
        fmt.Println("Infinite Loop Iteration:", k)
        k++
    }

    // Using Continue Statement
    for l := 0; l < 5; l++ {
        if l%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println("Continue Statement Iteration (odd numbers):", l)
    }

    // Nested Loops
    for m := 1; m <= 3; m++ {
        for n := 1; n <= 2; n++ {
            fmt.Printf("Nested Loop Iteration: m=%d, n=%d\n", m, n)
        }
    }

    // Range Statement with Slice
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("Range Statement Iteration: index=%d, value=%d\n", index, value)
    }
}
```