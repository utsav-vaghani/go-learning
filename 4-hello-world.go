/**
	- Every single files of GO code contains this statement as in GO.
	- packaging can be consider as directory structure which helps in writing modularized code.
	- Each package contains some specific kind of sentimental functionalities, for instance `controllers`.
	- A package can contains multiple sub-packages.
	- Accessability of functions, types or variable within cross-packages depends on whether its exported or unexported.

	- Naming convention: package name should not contains underscores(_), hiphens(-) or camelCase.
	- Driver code should be placed under package `main`
**/
package main

/**
	import section, as we need to import a package `fmt` we have single line import section.
	in upcoming chapters, we will use multi-line import statements.
	import section contains built-in packages,  third party libraries, internal libraries etc.
 **/
import "fmt"

/**
	As our `main` function contains `m` (lowercase) as first character, its unexported so it can not
	be accessible out side of main package.
**/
func main() {
	fmt.Println("Hello world")
}

/**
	To run this program use command:
		`go run 4-hello-world.go`
	when we hit this command GO will try to find out `main` function inside `4-hello-world.go` and run it.
 **/
