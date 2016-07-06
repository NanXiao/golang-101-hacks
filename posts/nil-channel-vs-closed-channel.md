# nil channel VS closed channel
----
The zero value of channel type is `nil`, and the send and receive operations on a `nil` channel will always block. Check the following example: 

	package main

	import "fmt"
	
	func main() {
	        var ch chan int
	
	        go func(c chan int) {
	                for v := range c {
	                        fmt.Println(v)
	                }
	        }(ch)
	
	        ch <- 1
	}

The running result is like this:  

	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send (nil chan)]:
	main.main()
	        /root/nil_channel.go:14 +0x64
	
	goroutine 5 [chan receive (nil chan)]:
	main.main.func1(0x0)
	        /root/nil_channel.go:9 +0x53
	created by main.main
	        /root/nil_channel.go:12 +0x37
We can see the `main` and `func` goroutines are both blocked.  

The `Go`'s built-in `close` function can be used to close the channel which must not be receive-only, and it should always be executed by sender, not receiver. Closing a `nil` channel will cause panic. See the following example:  

	package main
	
	func main() {
	        var ch chan int
	        close(ch)
	}
The running result is like this:  

	panic: close of nil channel

	goroutine 1 [running]:
	panic(0x456500, 0xc82000a170)
	 /usr/local/go/src/runtime/panic.go:481 +0x3e6
	main.main()
	 /root/nil_channel.go:5 +0x1e

Furthermore, there are also some subtleties of operating an already-closed channel:  

(1) Close an already channel also cause panic:  

	package main
	
	func main() {
	        ch := make(chan int)
	        close(ch)
	        close(ch)
	}
The running result is like this:  

	panic: close of closed channel
	
	goroutine 1 [running]:
	panic(0x456500, 0xc82000a170)
	 /usr/local/go/src/runtime/panic.go:481 +0x3e6
	main.main()
	 /root/nil_channel.go:6 +0x4d

(2) Send on a closed channel will also introduce panic:  

	package main

	func main() {
	        ch := make(chan int)
	        close(ch)
	        ch <- 1
	}

The running result is like this:  

	panic: send on closed channel
	
	goroutine 1 [running]:
	panic(0x456500, 0xc82000a170)
	 /usr/local/go/src/runtime/panic.go:481 +0x3e6
	main.main()
	 /root/nil_channel.go:6 +0x6c

(3) Receive on a closed channel will return the zero value for the channel's type without blocking:  

	package main

	import "fmt"
	
	func main() {
	        ch := make(chan int)
	        close(ch)
	        fmt.Println(<-ch)
	}
The executing result is like this:  

	0

The following is a summary of "`nil` channel VS closed channel":  

<table>
    <tr>
      <th>Operation type</th>
	  <th>Nil channel</th>
	  <th>Closed channel</th>
    </tr>
    <tr>
      <td>Send</td>
	  <td>Block</td>
	  <td>Panic</td>
    </tr>
    <tr>
      <td>Receive</td>
	  <td>Block</td>
	  <td>Not block, return zero value of channel's type</td>
    </tr>
	<tr>
      <td>Close</td>
	  <td>Panic</td>
	  <td>Panic</td>
    </tr>
</table>

References:  
[Package builtin](https://golang.org/pkg/builtin/#close);  
[Is it OK to leave a channel open?](http://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open);  
[The Go Programming Language Specification](https://golang.org/ref/spec). 

