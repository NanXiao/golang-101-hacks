Conversion between array and slice
----
In `Go`, array is a fixed length of continuous memory with specified type, while slice is just a reference which points to an underlying array. Since they are different types, they can't assign value each other directly. See the following example:  

    package main
    
    import "fmt"
    
    func main() {
    	s := []int{1, 2, 3}
    	var a [3]int
    
    	fmt.Println(copy(a, s))
    }
Because `copy` only accepts slice argument, we can use the `[:]` to create a slice from array. Check next code:  

	package main
	
	import "fmt"
	
	func main() {
		s := []int{1, 2, 3}
		var a [3]int
	
		fmt.Println(copy(a[:2], s))
		fmt.Println(a)
	}

The running output is:  

	2
	[1 2 0]

The above example is copying value from slice to array, and the opposite operation is similar:  

	package main

	import "fmt"

    func main() {
    	a := [...]int{1, 2, 3}
    	s := make([]int, 3)
    
    	fmt.Println(copy(s, a[:2]))
    	fmt.Println(s)
    }

The execution result is:  

	2
	[1 2 0]

References:  
[In golang how do you convert a slice into an array](http://stackoverflow.com/questions/19073769/in-golang-how-do-you-convert-a-slice-into-an-array);  
[Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices).