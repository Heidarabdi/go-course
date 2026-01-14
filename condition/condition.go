package main
import "fmt"

func main() {
	// if statement
	num := 7
	if num > 0 {
		fmt.Println("Number is greater than 9")
	}

	// if-else statement
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Else-if ladder
	if num < 0 {
		fmt.Println("Number is negative")
	} else if num == 0 {
		fmt.Println("Number is zero")
	} else {
		fmt.Println("Number is positive")
	}

	// Nested if statements
	if num >= 0 {
		if num%2 == 0 {
			fmt.Println("Number is non-negative and even")
		} else {
			fmt.Println("Number is non-negative and odd")
		}
	}

	// switch statement (single case)
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// switch statement (multiple cases)
	month := 5
	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Autumn")
	default:
		fmt.Println("Invalid month")
	}

}
