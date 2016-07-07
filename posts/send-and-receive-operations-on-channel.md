# Send and receive operations on channel
----
`Go`'s built-in `channel` type provides a handy method for communicating and synchronizing: The producer pushes data into channel and the consumer pulls data from it.  
 
The send operation on channel is simple, as long as the filled-in stuff is a valid expression and matches the channel's type:  

	Channel <- Expression
Take the following code as an example:  

	package main

	func send() int {
		return 2
	}
	func main()  {
		ch := make(chan int, 2)
		ch <- 1
		ch <- send()
	}
Receive operation on channel pulls the value from the channel, and you can save it or discard it if you don't care what you have got. Check the following example:    

	package main
	
	import "fmt"
	
	func main()  {
		ch := make(chan int)
		go func(ch chan int) {
			ch <- 1
			ch <- 2
		}(ch)
		<-ch
		fmt.Println(<-ch)
	}
The running result is `2`, and that's because the first value (`1`) is left out in `<-ch` statement.

Compared to its send sibling, the receive operation is a little tricky: in assignment and initialization, there will be another return value which indicates whether this communication is successful or not. And the idiom of this variable's name is `ok`:  

	v, ok := <- ch 
The value of `ok` is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it is a zero value generated because the channel is closed and empty. That means although the channel is closed, as long as there is still data in the channel, the receive operation can of course get things from it. See the following code:  

	package main
	
	import "fmt"
	
	func main()  {
		ch := make(chan int)
	
		go func(ch chan int) {
			ch <- 1
			ch <- 2
			close(ch)
		}(ch)
	
		for i := 1; i <= 3; i++ {
			v, ok := <- ch
			fmt.Printf("value is %d, ok is %v\n", v, ok)
		}
	} 

The executing result is like this:  

	value is 1, ok is true
	value is 2, ok is true
	value is 0, ok is false

We can see after `func` goroutine executes closing channel operation, the value of `v` got from channel is the zero value of integer type: `0`, and `ok` is `false`.  

Reference:  
[The Go Programming Language Specification](https://golang.org/ref/spec).
 