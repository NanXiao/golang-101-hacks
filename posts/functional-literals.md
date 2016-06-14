Functional literals
----
A functional literal just represents an anonymous function. You can assign functional literal to a variable:  

    package main

	import (
		"fmt"
	)
	
	func main() {
		f := func() { fmt.Println("Hello, 中国！") }
		f()
	}
Or invoke functional literal directly (Please notice the `()` at the end of functional literal):  

	package main

	import (
		"fmt"
	)
	
	func main() {
		func() { fmt.Println("Hello, 中国！") }()
	}
The above `2` programs both output "`Hello, 中国！`".  

Functional literal is also a closure, so it can access the variables of its surrounding function. Check the following example which your real intention is `1` and `2` are printed:  

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		for i := 1; i <= 2; i++ {
			go func() {fmt.Println(i)}()
		}
		time.Sleep(time.Second)
	}
But the output is:  

	3
	3
The cause is the `func` goroutines don't get the opportunity to run until the `main` goroutine sleeps, and at that time, the variable `i` has been changed to `3`. Modify the above program as follows:

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		for i := 1; i <= 2; i++ {
			go func() {fmt.Println(i)}()
			time.Sleep(time.Second)
		}
	} 
The `func` goroutine can run before `i` is changed, so the running result is what you expect:  
  
	1
	2
But the idiom method should be passing `i` as an argument of the functional literal:  

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		for i := 1; i <= 2; i++ {
			go func(i int) {fmt.Println(i)}(i)
		}
		time.Sleep(time.Second)
	}

In above program, When "`go func(i int) {fmt.Println(i)}(i)`" is executed (Note: not goroutine is executed.), `i` defined in `main()` is assigned to `func`'s local parameter `i`. And the result is:  

	1
	2

P.S. You should notice, If you pass an argument while not use it, the `Go` compiler doesn't complain, but the closure will use the variable inherited from the parent function. That means the following statement:  

	go func(int) {fmt.Println(i)}(i) 
equals to:  

	go func() {fmt.Println(i)}()

References:  
[The Go Programming Language Specification](https://golang.org/ref/spec#Function_literals);  
[A question about passing arguments to closure](https://groups.google.com/forum/#!topic/golang-nuts/JXTEYyoPLio);  
[Why add “()” after closure body in Golang?](http://stackoverflow.com/questions/16008604/why-add-after-closure-body-in-golang).