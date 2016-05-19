# “nil slice” vs “nil map”
----
Slice and map are all reference types in `Golang`, and their default values are `nil`:  

	package main
	
	import "fmt"
	
	func main() {
		var (
			s []int
			m map[int]bool
		)
		if s == nil {
			fmt.Println("The value of s is nil")
		}
		if m == nil {
			fmt.Println("The value of m is nil")
		}
	}  
The result is like this：  

	The value of s is nil
	The value of m is nil

When a slice's value is `nil`, you could also do operations on it, such as `append`:  

	package main
	
	import "fmt"
	
	func main() {
		var s []int
		fmt.Println("Is s a nil? ", s == nil)
		s = append(s, 1)
		fmt.Println("Is s a nil? ", s == nil)
		fmt.Println(s)
	}

The result is like this：  

	Is s a nil?  true
	Is s a nil?  false
	[1]

A caveat you should notice is the length of both `nil` and empty slice is `0`:  
  
	package main
	
	import "fmt"
	
	func main() {
		var s1 []int
		s2 := []int{}
		fmt.Println("Is s1 a nil? ", s1 == nil)
		fmt.Println("Length of s1 is: ", len(s1))
		fmt.Println("Is s2 a nil? ", s2 == nil)
		fmt.Println("Length of s2 is: ", len(s2))
	}
The result is like this：  

	Is s1 a nil?  true
	Length of s1 is:  0
	Is s2 a nil?  false
	Length of s2 is:  0
So you should compare the slice's value with `nil` to check whether it is a `nil`.  

Accessing a `nil` map is OK, but storing a `nil` map cause program panic:  

	package main
	
	import "fmt"
	
	func main() {
		var m map[int]bool
		fmt.Println("Is m a nil? ", m == nil)
		fmt.Println("m[1] is ", m[1])
		m[1] = true
	}

The result is like this:  

	Is m a nil?  true
	m[1] is  false
	panic: assignment to entry in nil map
	
	goroutine 1 [running]:
	panic(0x4cc0e0, 0xc082034210)
		C:/Go/src/runtime/panic.go:481 +0x3f4
	main.main()
		C:/Work/gocode/src/Hello.go:9 +0x2ee
	exit status 2
	
	Process finished with exit code 1

So the best practice is to initialize a `map` before using it, like this:  

	m := make(map[int]bool)

BTW, you should use the following pattern to check whether there is an element in map or not:  

	if v, ok := m[1]; !ok {
		.....
	}

Reference:  
[The Go Programming Language](http://www.gopl.io/).

