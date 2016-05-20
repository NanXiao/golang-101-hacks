# Prepend
----
`Go` has a built-in [append](https://golang.org/pkg/builtin/#append) function which add elements in the slice:  
	
	func append(slice []Type, elems ...Type) []Type

But how if we want to the "prepend" effect? Maybe we should use `copy` function. E.g.:  

	package main

	import "fmt"
	
	func main()  {
		var s []int = []int{1, 2}
		fmt.Println(s)
	
		s1 := make([]int, len(s) + 1)
		s1[0] = 0
		copy(s1[1:], s)
		s = s1
		fmt.Println(s)
	
	} 

The result is like this:  

	[1 2]
	[0 1 2]

But the above code looks ugly and cumbersome, so an elegant implementation maybe here:  

	s = append([]int{0}, s...)

BTW, I also have tried to write a "general-purpose" prepend:  

	func Prepend(v interface{}, slice []interface{}) []interface{}{
		return append([]interface{}{v}, slice...)
	}
But since `[]T` can't convert to an `[]interface{}` directly (please refer [https://golang.org/doc/faq#convert_slice_of_interface](https://golang.org/doc/faq#convert_slice_of_interface), it is just a toy, not useful.     
 
Reference:  
[Go – append/prepend item into slice](https://codingair.wordpress.com/2014/07/18/go-appendprepend-item-into-slice/).