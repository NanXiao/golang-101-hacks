In `Golang`, the length is also a part of array type. So the following code declares an array:  

	var array [3]int
while "`var slice []int`" defines a slice. Because of this characteristic, arrays with the same array element type but different length can't assign values each other. I.E.:

	package main
	
	import "fmt"
	
	func main() {
		var a1 [2]int
		var a2 [3]int
		a2 = a1
		fmt.Println(a2)
	}
The compiler will complain:  

	cannot use a1 (type [2]int) as type [3]int in assignment

Changing "`var a1 [2]int`" to "`var a1 [3]int`" will make it work.  

Another caveat you should pay attention to is the following code declares an array, not a slice:  

	array := [...]int {1, 2, 3} 
You can verify it by the following code:  

	package main
	
	import (
		"fmt"
		"reflect"
	)
	
	func main() {
		array := [...]int {1, 2, 3}
		slice := []int{1, 2, 3}
		fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))
	}

The output is:  

	[3]int []int

Additionally, since in `Golang`, the function argument is passed by "value", so if you use an array as a function argument, the function just does the operations on the copy of the original copy. Check the following code: 

	package main

	import (
		"fmt"
	)
	
	func changeArray(array [3]int) {
		for i, _ := range array {
			array[i] = 1
		}
		fmt.Printf("In changeArray function, array address is %p, value is %v\n", &array, array)
	}
	
	func main() {
		var array [3]int
	
		fmt.Printf("Original array address is %p, value is %v\n", &array, array)
		changeArray(array)
		fmt.Printf("Changed array address is %p, value is %v\n", &array, array)
	}   

The output is:  

	Original array address is 0xc082008680, value is [0 0 0]
	In changeArray function, array address is 0xc082008700, value is [1 1 1]
	Changed array address is 0xc082008680, value is [0 0 0]

From the log, you can see the array's address in `changeArray` function is not the same with array's address in `main` function, so the content of original array will definitely not be modified. Furthermore, if the array is very large, copying them when passing argument to function may generate more overhead than you want, you should know about it.   