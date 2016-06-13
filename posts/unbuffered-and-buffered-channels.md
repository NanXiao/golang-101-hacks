Unbuffered and buffered channels
----
`Go`'s built-in `channel` type provides a handy method for communicating and synchronizing. The channel is divided into two categories: unbuffered and buffered.  

(1) Unbuffered channel  
For unbuffered channel, the sender will block on the channel until the receiver receives the data from the channel, whilst the receiver will also block on the channel until sender sends data into the channel. Check the following example:  

    package main
    
    import (
    	"fmt"
    	"time"
    )
    
    func main() {
    	ch := make(chan int)
    
    	go func(ch chan int) {
    		fmt.Println("Func goroutine begins sending data")
    		ch <- 1
    		fmt.Println("Func goroutine ends sending data")
     	}(ch)
    
    	fmt.Println("Main goroutine sleeps 2 seconds")
    	time.Sleep(time.Second * 2)
    	
    	fmt.Println("Main goroutine begins receiving data")
    	d := <- ch
    	fmt.Println("Main goroutine received data:", d)
    
    	time.Sleep(time.Second)
    }
The running result likes this:  

	Main goroutine sleeps 2 seconds
	Func goroutine begins sending data
	Main goroutine begins receiving data
	Main goroutine received data: 1
	Func goroutine ends sending data
After the `main` goroutine is launched, it will sleep immediately("`Main goroutine sleeps 2 seconds`" is printed), and this will cause `main` goroutine relinquishes the `CPU` to the `func` goroutine("`Func goroutine begins sending data`" is printed). But since the `main` goroutine is sleeping and can't receive data from the channel, so `ch <- 1` operation in `func` goroutine can't complete until `d := <- ch` in `main` goroutine is executed(The final `3` logs are printed).  

(2) Buffered channel  
Compared with unbuffered counterpart, the sender of buffered channel will block when there is **no** empty slot of the `channel`, while the receiver will block on the channel when it is empty. Modify the above example:  

	package main
	
	import (
		"fmt"
		"time"
	)
	
	func main() {
		ch := make(chan int, 2)
	
		go func(ch chan int) {
			for i := 1; i <= 5; i++ {
				ch <- i
				fmt.Println("Func goroutine sends data: ", i)
			}
			close(ch)
		}(ch)
	
		fmt.Println("Main goroutine sleeps 2 seconds")
		time.Sleep(time.Second * 2)
	
		fmt.Println("Main goroutine begins receiving data")
		for d := range ch {
			fmt.Println("Main goroutine received data:", d)
		}
	}
The executing result is as follows:  

	Main goroutine sleeps 2 seconds
	Func goroutine sends data:  1
	Func goroutine sends data:  2
	Main goroutine begins receiving data
	Main goroutine received data: 1
	Main goroutine received data: 2
	Main goroutine received data: 3
	Func goroutine sends data:  3
	Func goroutine sends data:  4
	Func goroutine sends data:  5
	Main goroutine received data: 4
	Main goroutine received data: 5
In this sample, since the channel has `2` slots, so the `func` goroutine will not block until it sends the third element.  

P.S., "`make(chan int, 0)`" is equal to "`make(chan int)`", and it will create an unbuffered `int` channel too.