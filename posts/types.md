# Types
----
Types in `Go` are divided into `2` categories: named and unnamed. Besides predeclared types (such as `int`, `rune`, etc), you can also define named type yourself. E.g.:  

	type mySlice []int
Unnamed types are defined by type literal. I.e., `[]int` is an unnamed type.  

According to [Go spec](https://golang.org/ref/spec#Types), there is an underlying type of every type:  

>Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.

So, in above example, both named type `mySlice` and unnamed type `[]int` have the same underlying type: `[]int`.  

`Go` has strict rules of assigning values of variables. For example:  

	package main
	
	import "fmt"
	
	type mySlice1 []int
	type mySlice2 []int
	
	func main() {
		var s1 mySlice1
		var s2 mySlice2 = s1
	
		fmt.Println(s1, s2)
	}
The compilation will complain the following error:  

	cannot use s1 (type mySlice1) as type mySlice2 in assignment

Although the underlying type of `s1` and `s2` are same: `[]int`, but they belong to `2` different named types (`mySlice1` and `mySlice2`), so they can't assign values each other. But if you modify `s2`'s type to `[]int`, the compilation will be OK:  

	package main

	import "fmt"
	
	type mySlice1 []int
	
	func main() {
		var s1 mySlice1
		var s2 []int = s1
	
		fmt.Println(s1, s2)
	}

The magic behind it is one rule of [assignability](https://golang.org/ref/spec#Assignability):  

>x's type V and T have identical underlying types and at least one of V or T is not a named type.

References:  
[Go spec](https://golang.org/ref/spec#Types);  
[Learning Go - Types](http://www.laktek.com/2012/01/27/learning-go-types/);  
[Golang pop quiz](https://twitter.com/davecheney/status/734646224696016901).