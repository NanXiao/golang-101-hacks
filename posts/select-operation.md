# Select operation
----
`Go`'s `select` operation looks similar to `switch`, but it's dedicated to poll send and receive operations channels. Check the following example:  

	package main
	
	import (
	        "fmt"
	        "time"
	)
	
	func main() {
	        ch1 := make(chan int)
	        ch2 := make(chan int)
	
	        go func(ch chan int) { <-ch }(ch1)
	        go func(ch chan int) { ch <- 2 }(ch2)
	
	        time.Sleep(time.Second)
	
	        for {
	                select {
	                case ch1 <- 1:
	                        fmt.Println("Send operation on ch1 works!")
	                case <-ch2:
	                        fmt.Println("Receive operation on ch2 works!")
	                default:
	                        fmt.Println("Exit now!")
	                        return
	                }
	        }
	}
The running result is like this:  

	Send operation on ch1 works!
	Receive operation on ch2 works!
	Exit now!

The `select` operation will check which `case` branch can be run, that means the send or receive action can be executed successfully. If more than one `case` are ready now, the `select` will randomly choose one to execute. If no `case` is ready, but there is a `default ` branch, then the `default ` block will be executed, else the `select` operation will block. In the above example, if the `main` goroutine doesn't sleep (`time.Sleep(time.Second)`), the other `2 func` goroutines won't obtain the opportunity to run, so only `default` block in `select` statement will be executed.  

The `select` statement won't process `nil` channel, so if a channel used for receive operation is closed, you should mark its value as `nil`, then it will be kicked out of the selection list. So a common pattern of selection on multiple receive channels looks like this:  

	for ch1 != nil && ch2 != nil {
	    select {
	    case x, ok := <-ch1:
	        if !ok {
	            ch1 = nil
				break
	        }
			......
	    case x, ok := <-ch2:
	        if !ok {
	            ch2 = nil
				break
	        }
			......
	    }
	}  

References:  
[The Go Programming Language Specification](https://golang.org/ref/spec);  
[breaking out of a select statement when all channels are closed](http://stackoverflow.com/questions/13666253/breaking-out-of-a-select-statement-when-all-channels-are-closed);  
[Curious Channels](http://dave.cheney.net/2013/04/30/curious-channels).

 
