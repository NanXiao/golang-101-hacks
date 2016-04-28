In `Golang`, the packages can be divided into `2` categories:  

(1) `main` package: is used to generate the executable binary, and the `main` function is the entry point of the program. Take `hello.go` as an example:  

	package main
	
	import "greet"
	
	func main() {
		greet.Greet()
	}
  

(2) This category can also include `2` types:  
 
a) Library package: is used to generate the object files that can be reused by others. Take `greet.go` as an example:  

	package greet
	
	import "fmt"
	
	func Greet() {
		fmt.Println("Hello 中国!")
	}

b) Some other packages for special purposes, such as testing, etc.

