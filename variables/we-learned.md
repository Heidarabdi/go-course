# We Learned
- Go variables have types like `int`, `string`, `bool`, `float64`, etc.
- variables are declared like this: `var varName type = value, or if you dont want to specify the type, you can use `varName := value` so it can infer the type automatically.
## Variables Declaration variables can be declared different way: 
- `var stdName string = "John Doe" or `stdName := "John Doe"` or `var stdName = "John Doe"` its called declaration with values.
- variables can be declared without initial values, like this: `var varName type`, and if we print it will show the zero value of the type like string for "", bool for false, int for 0, float64 for 0.0, etc.
- value assigment after declaration like `var stdName string` then `stdName = "John Doe"`
- note isnt possible to declare variable using `:=` without initial values.
- also remeber `:=` can't be used outside a function body or it will cause an error.
- Multiple declaration in one line like `var a, b, c int = 1, 2, 3` or `a, b, c := 1, 2, 3`
- also we can use `type` keyword to create multiple variables but it will only use with one tupe like `var a, b, c string = "A", "B", "C"` , if we want to use different types in multiple declaration we dont need to use `type` keyword like `var a, b, c = "A", 2, true` or `a, b, c := "A", 2, true`
- we can do also bracket multiple declaration like: `var ( name string, age int = 30 isStudent bool = false )`.
- `const` keyword is used to declare constant variables that cannot be changed after declaration like `const pi = 3.14` or `const greeting string = "Hello, World!"`
- `const` is the same `var` but its value cant be changed after declaration.
## Variable Naming
- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords
## GO output
- To print there is three functions: `fmt.Print()`, `fmt.Println()`, and `fmt.Printf()`
- `fmt.Print()` is used to print the value without a new line.
- `fmt.Println()` is used to print the value with a new line.
- `fmt.Printf()` is used to print formatted strings using verbs like `%d` for integers, `%s` for strings, `%f` for floats, etc.