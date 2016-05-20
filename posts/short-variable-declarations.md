# Short variable declaration
----
Short variable declaration is a very convenient manner of "declaring variable" in `Golang`:  

	i := 10

It is shorthand of following (Please notice there is no type):  

	var i = 10

The `Golang` compiler will infer the type according to the value of variable. It is a very handy feature, but on the other side of coin, it also brings some pitfalls which you should pay attention to:  

(1) This format can only be used in functions:  

	package main

	i := 10

	func main() {
		fmt.Println(i)
	}

The compiler will complain the following words:  

	syntax error: non-declaration statement outside function body

(2) You must declare **at least 1 new variable**: 

	package main
	
	import "fmt"
	
	func main() {
	    var i = 1
	
	    i, err := 2, true
	
	    fmt.Println(i, err)
	}

In `i, err := 2, false` statement, only `err` is a new declared variable, `var` is  actually declared in `var i = 1`.  

(3) The short variable declaration can shadow the global variable declaration, and it may not be what you want, and gives you a big surprise:  

	package main
	
	import "fmt"
	
	var i = 1
	
	func main() {
	
	    i, err := 2, true
	
	    fmt.Println(i, err)
	}

`i, err := 2, true` actually declares a **new local i** which makes the **global i** inaccessible in `main` function. To use the global variable but not introducing a new local one, one solution maybe like this:  

	package main
	
	import "fmt"
	
	var i int
	
	func main() {
	
	    var err bool
	
	    i, err = 2, true
	
	    fmt.Println(i, err)
	}

参考资料：  
[Short variable declarations](https://golang.org/ref/spec#Short_variable_declarations).
