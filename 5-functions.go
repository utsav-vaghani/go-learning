package main


/**
	GO provides different way of defining functions.
	one can define a function using keyword `func`.
	
	Function definitions:
		1. Normal function (No parameters, No returning values)
			func <function-name>() {
				// function statements.
			}

			ex. 
			func main() {
				fmt.Println("hello-world")
			}
		
		2. Function with parameters (No returning values)
			func <function-name>(par1 type1, par2 type2, ....) {
				// function statements.
			}

			ex.
			func helloworld(name string) {
				fmt.Println("hello ", name)
			}

		3. Function with returning values
			- In GO a function can return multiple values.
			func <function-name>(parameters...) (type1, type2, ...)

			ex.
			func helloworld(name string) string {
				return "hello " + name
			}
		
		4. Function with named return values
			- In GO, one cane declare a variable for returning values which will be returned to caller
				while returning from the function.
			- It also called naked-return as to return values we dont need to give variables.

			func <function-name>(parameters...) (ret1 type1, ret2 type2, ...) {
				// function statements
			}

			ex.
			func helloworld(name string)(helloStr string) {
				helloStr = "hello " + name
				return 		// naked return as we are not returning `helloStr`
			}

			- Facts
				naked return should not be used in big function definition as its reduce the readability.
		
		5. Function with the receiver.
			- In GO, there is no such class but we can implement the OOP paradigm using defining methods on 
				types.
			- we can define methods on any type. and access that method using variable of that type.
			- Difference b/w method and function: method is a function on any type.

			To define method on any type we can use the receiver.
			func (t type) <function-name>(parameters...)(return...) {
				// function statements.
			}
			here t is receiver variable and `type` can be any type

			we will cover receiver concept in detail in upcoming chapters.
	
	- Functions can be Exported or unexported. Exported functions are those, which has first character of name
	  is in upper-case.
	  ex. func HelloWorld() {} >> it is exported as `H` in name `HelloWorld` is in upper-case.
	      func helloWorld() {} >> it is unexported as `h` in name `helloWorld` is in lower-case.
	- Functions those are exported are accessible out of package as well.
	- Un-exported functions are accessible on defining package only.

	Warning:
		- Be conscious while exporting functions.
**/