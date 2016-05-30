Goroutine
----
A running `Go` program is composed of one or more goroutines, and each goroutine can be considered as an independent task. Goroutine and thread have many commonalities, such as: every goroutine(thread) has its private stack and registers; if the main goroutine(thread) exits, the program will exit, and so on. But on modern Operating System (E.g., `Linux`), the actual execution and scheduled unit is thread, so if a goroutine wants to become running, it must "attach" to a thread. Let's see an example:  

    package main
    
    import (
    	"time"
    )
    
    func main() {
    	time.Sleep(1000 * time.Second)
    }
What the program does is just sleeping for a while, not does anything useful. After launching it on `Linux`, use `Delve` to attach the running process and observe the details of it:  

    (dlv) threads
    * Thread 1040 at 0x451f73 /usr/local/go/src/runtime/sys_linux_amd64.s:307 runtime.futex
      Thread 1041 at 0x451f73 /usr/local/go/src/runtime/sys_linux_amd64.s:307 runtime.futex
      Thread 1042 at 0x451f73 /usr/local/go/src/runtime/sys_linux_amd64.s:307 runtime.futex
      Thread 1043 at 0x451f73 /usr/local/go/src/runtime/sys_linux_amd64.s:307 runtime.futex
      Thread 1044 at 0x451f73 /usr/local/go/src/runtime/sys_linux_amd64.s:307 runtime.futex
We can see there are `5` threads of this process, let's confirm it by checking `/proc/1040/task/` directory:  

	# cd /proc/1040/task/
	# ls
	1040  1041  1042  1043  1044
Yeah, the thread information of `Delve` is right! Check the particulars of goroutines:  

	(dlv) goroutines
	[4 goroutines]
	  Goroutine 1 - User: /usr/local/go/src/runtime/time.go:59 time.Sleep (0x43e236)
	  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x426f73)
	  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x426f73)
	* Goroutine 4 - User: /usr/local/go/src/runtime/lock_futex.go:206 runtime.notetsleepg (0x40b1ce)
	  
There is only one `main` goroutine, what the hell of the other `3` goroutines? Actually, the other `3` goroutines are system goroutines, and you can refer related info [here](https://github.com/derekparker/delve/issues/553). The number of `main` goroutine is `1`, and you can inspect it:  

	(dlv) goroutine 1
	Switched from 4 to 1 (thread 1040)
	(dlv) bt
	0  0x0000000000426f73 in runtime.gopark
	   at /usr/local/go/src/runtime/proc.go:263
	1  0x0000000000426ff3 in runtime.goparkunlock
	   at /usr/local/go/src/runtime/proc.go:268
	2  0x000000000043e236 in time.Sleep
	   at /usr/local/go/src/runtime/time.go:59
	3  0x0000000000401013 in main.main
	   at ./gocode/src/goroutine.go:8
	4  0x0000000000426b9b in runtime.main
	   at /usr/local/go/src/runtime/proc.go:188
	5  0x0000000000451000 in runtime.goexit
	   at /usr/local/go/src/runtime/asm_amd64.s:1998
 
Using `go` keyword can create and start a goroutine, see another case:  

    package main
    
    import (
    	"fmt"
    	"time"
    )
    
    func main() {
    	ch := make(chan int)
    
    	go func(chan int) {
    		var count int
    		for {
    			count++
    			ch <- count
    			time.Sleep(10 * time.Second)
    		}
    	}(ch)
    
    	for v := range ch {
    		fmt.Println(v)
    	}
    }
    
The `go func` statement spawns another goroutine which works as a producer; while the `main` goroutine behaves as a consumer. And the output should be:  

	1
	2
	......
Use `Delve` to check the goroutine aspects:  

	(dlv) goroutines
	[6 goroutines]
	  Goroutine 1 - User: ./gocode/src/goroutine.go:20 main.main (0x40106c)
	  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x429fc3)
	  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x429fc3)
	  Goroutine 4 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x429fc3)
	  Goroutine 5 - User: /usr/local/go/src/runtime/time.go:59 time.Sleep (0x442ab6)
	* Goroutine 6 - User: /usr/local/go/src/runtime/lock_futex.go:206 runtime.notetsleepg (0x40cf4e)
	(dlv) goroutine 1
	Switched from 6 to 1 (thread 1997)
	(dlv) bt
	0  0x0000000000429fc3 in runtime.gopark
	   at /usr/local/go/src/runtime/proc.go:263
	1  0x000000000042a043 in runtime.goparkunlock
	   at /usr/local/go/src/runtime/proc.go:268
	2  0x00000000004047eb in runtime.chanrecv
	   at /usr/local/go/src/runtime/chan.go:470
	3  0x0000000000404354 in runtime.chanrecv2
	   at /usr/local/go/src/runtime/chan.go:360
	4  0x000000000040106c in main.main
	   at ./gocode/src/goroutine.go:20
	5  0x0000000000429beb in runtime.main
	   at /usr/local/go/src/runtime/proc.go:188
	6  0x0000000000455de0 in runtime.goexit
	   at /usr/local/go/src/runtime/asm_amd64.s:1998
	(dlv) goroutine 5
	Switched from 1 to 5 (thread 1997)
	(dlv) bt
	0  0x0000000000429fc3 in runtime.gopark
	   at /usr/local/go/src/runtime/proc.go:263
	1  0x000000000042a043 in runtime.goparkunlock
	   at /usr/local/go/src/runtime/proc.go:268
	2  0x0000000000442ab6 in time.Sleep
	   at /usr/local/go/src/runtime/time.go:59
	3  0x00000000004011d6 in main.main.func1
	   at ./gocode/src/goroutine.go:16
	4  0x0000000000455de0 in runtime.goexit
	   at /usr/local/go/src/runtime/asm_amd64.s:1998
The number of `main` goroutine is `1`, whilst `func` is `5`.  

Another caveat you should pay attention to is the switch point among goroutines. It can be blocking system call, channel operations, etc.  

Reference:  
[Effective Go](https://golang.org/doc/effective_go.html#goroutines);  
[Performance without the event loop](http://dave.cheney.net/2015/08/08/performance-without-the-event-loop);  
[How Goroutines Work](http://blog.nindalf.com/how-goroutines-work/). 
  