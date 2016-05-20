# Pass slice as a function argument
----
In `Go`, the function parameters are passed by value. With respect to use slice as a function argument, that means the function will get the copies of the slice: a pointer which points to the starting address of the underlying array, accompanied by the length and capacity of the slice. Oh boy! Since you know the address of the memory which is used to store the data, you can tweak the slice now. Let's see the following example:  

	package main
	
	import (
		"fmt"
	)
	
	func modifyValue(s []int)  {
		s[1] = 3
		fmt.Printf("In modifyValue: s is %v\n", s)
	}
	func main() {
		s := []int{1, 2}
		fmt.Printf("In main, before modifyValue: s is %v\n", s)
		modifyValue(s)
		fmt.Printf("In main, after modifyValue: s is %v\n", s)
	}
The result is here:  

	In main, before modifyValue: s is [1 2]
	In modifyValue: s is [1 3]
	In main, after modifyValue: s is [1 3]
You can see, after running `modifyValue` function, the content of slice `s` is changed. Although the `modifyValue` function just gets a copy of the memory address of slice's underlying array, it is enough!  

See another example:  

	package main
	
	import (
		"fmt"
	)
	
	func addValue(s []int) {
		s = append(s, 3)
		fmt.Printf("In addValue: s is %v\n", s)
	}
	
	func main() {
		s := []int{1, 2}
		fmt.Printf("In main, before addValue: s is %v\n", s)
		addValue(s)
		fmt.Printf("In main, after addValue: s is %v\n", s)
	}

The result is like this:  

	In main, before addValue: s is [1 2]
	In addValue: s is [1 2 3]
	In main, after addValue: s is [1 2]

This time, the `addValue` function doesn't take effect on the `s` slice in `main` function. That's because it just manipulate the copy of the `s`, not the "real" `s`.   

So if you really want the function to change the content of a slice, you can pass the address of the slice:  

	package main
	
	import (
		"fmt"
	)
	
	func addValue(s *[]int) {
		*s = append(*s, 3)
		fmt.Printf("In addValue: s is %v\n", s)
	}
	
	func main() {
		s := []int{1, 2}
		fmt.Printf("In main, before addValue: s is %v\n", s)
		addValue(&s)
		fmt.Printf("In main, after addValue: s is %v\n", s)
	}	 

The result is like this:  

	In main, before addValue: s is [1 2]
	In addValue: s is &[1 2 3]
	In main, after addValue: s is [1 2 3]
