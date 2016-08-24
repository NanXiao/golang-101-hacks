# Type assertion and type switch
----
Sometimes, you may want to know the exact type of an interface variable. In this scenario, you can use `type assertion`:  

	x.(T)

`x` is the variable whose type must be **interface**, and `T` is the type which you want to check. For example:  

	package main
	
	import "fmt"
	
	func printValue(v interface{}) {
		fmt.Printf("The value of v is: %v", v.(int))
	}
	
	func main() {
		v := 10
		printValue(v)
	}
The running result is:  

	The value of v is: 10

In the above example, using `v.(int)` to assert the `v` is `int` variable.  

if the `type assertion` operation fails, a running panic will occur: change 

	fmt.Printf("The value of v is: %v", v.(int))  

into:  

	fmt.Printf("The value of v is: %v", v.(string))

Then executing the program will get following error:  

	panic: interface conversion: interface is int, not string

	goroutine 1 [running]:
	panic(0x4f0840, 0xc0820042c0)
	......

To avoid this, `type assertion` actually returns an additional `boolean` variable to tell whether this operations holds or not. So modify the program as follows:  

	package main

	import "fmt"
	
	func printValue(v interface{}) {
		if v, ok := v.(string); ok {
			fmt.Printf("The value of v is: %v", v)
		} else {
			fmt.Println("Oops, it is not a string!")
		}
	
	}
	
	func main() {
		v := 10
		printValue(v)
	}
This time, the output is:  

	Oops, it is not a string!

Furthermore, you can also use `type switch` which makes use of `type assertion` to determine the type of variable, and do the operations accordingly. Check the following example:  

	package main

	import "fmt"
	
	func printValue(v interface{}) {
		switch v := v.(type) {
		case string:
			fmt.Printf("%v is a string\n", v)
		case int:
			fmt.Printf("%v is an int\n", v)
		default:
			fmt.Printf("The type of v is unknown\n")
		}
	}
	
	func main() {
		v := 10
		printValue(v)
	}
The running result is here:  

	10 is an int
Compared to `type assertion`, `type switch` uses keyword `type` instead of the specified variable type (such as `int`) in the parentheses.  

References:  
[Effective Go](https://golang.org/doc/effective_go.html);
[Go – x.(T) Type Assertions](https://codingair.wordpress.com/2014/07/21/go-x-t-type-assertions/);  
[How to find a type of a object in Golang?](http://stackoverflow.com/questions/20170275/how-to-find-a-type-of-a-object-in-golang).