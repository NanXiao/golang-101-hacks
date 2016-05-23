switch
----
Compared to other programming languages (such as `C`), `Go`'s `switch-case` statement doesn't need explicit "`break`", and not have `fall-though` characteristic. Take the following code as an example:  

	package main

	import (
		"fmt"
	)
	
	func checkSwitch(val int) {
		switch val {
		case 0:
		case 1:
			fmt.Println("The value is: ", val)
		}
	}
	func main() {
		checkSwitch(0)
		checkSwitch(1)
	}

The output is:  

	The value is:  1

Your real intention is the "`fmt.Println("The value is: ", val)`" will be executed when `val` is `0` or `1`, but in fact, the statement only takes effect when `val` is `1`. To fulfill your request, there are `2` methods:  

(1) Use `fallthrough`:  

	
	func checkSwitch(val int) {
		switch val {
		case 0:
			fallthrough
		case 1:
			fmt.Println("The value is: ", val)
		}
	}
	
(2) Put `0` and `1` in the same `case`:  

	func checkSwitch(val int) {
		switch val {
		case 0, 1:
			fmt.Println("The value is: ", val)
		}
	}

`switch` can also be used as a better `if-else`, and you may find it may be more clearer and simpler than multiple `if-else` statements.E.g.:    

	package main
	
	import (
		"fmt"
	)
	
	func checkSwitch(val int) {
		switch {
		case val < 0:
			fmt.Println("The value is less than zero.")
		case val == 0:
			fmt.Println("The value is qual to zero.")
		case val > 0:
			fmt.Println("The value is more than zero.")
		}
	}
	func main() {
		checkSwitch(-1)
		checkSwitch(0)
		checkSwitch(1)
	}

The output is:  

	The value is less than zero.
	The value is qual to zero.
	The value is more than zero.