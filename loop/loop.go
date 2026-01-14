package main
import "fmt"

func main() {
	
	// for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// while loop (using for)
	j := 1
	for j <= 5 {
		fmt.Println("While Iteration:", j)
		j++
	}

	// infinite loop (commented out to prevent actual infinite loop)
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// break and continue
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue // skip even numbers
		}
		if k > 7 {
			break // exit loop if k > 7
		}
		fmt.Println("Odd Number:", k)
	}

	// nested loops
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 2; y++ {
			fmt.Printf("x: %d, y: %d\n", x, y)
		}
	}


	// Range loop
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}