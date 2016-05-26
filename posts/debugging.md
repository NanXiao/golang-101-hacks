Debugging
----
No one can write bug-free code, so debugging is a requisite skill for every software engineer. Here are some tips about debugging `Go` programs:  

(1) Print  
Yes! Printing logs seems the easiest method, but it is indeed the most effective approach in most cases. `Go` has provided a big family of printing functions in [fmt](https://golang.org/pkg/fmt/) package, and using them neatly is an expertise you should grasp.  

(2) Debugger  
In some scenarios, maybe you need the specialized debugger tools to help you spot the root cause. You can use `gdb`, but since "`GDB does not understand Go programs well.`" (from [here](https://golang.org/doc/gdb)), I suggest taking [Delve](https://github.com/derekparker/delve), a dedicated debugger for `Go`, instead.  

No matter using `gdb` or `Delve`, if you want to debug the executable file, you must pass `-gcflags "-N -l"` during compiling binary to disable code optimization, else some weird things can happen during debugging, such as you can't print the value of an already declared variable.  

Except debugging the precompiled file, `Delve` can compile and debug the code on the fly, see the following example:  

	package main
	
	import "fmt"
	
	func main() {
		ch := make(chan int)
		go func(chan int) {
			for _, v := range []int{1, 2} {
				ch <- v
			}
			close(ch)
		}(ch)
	
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("The channel is closed.")
	}
 
Debugging flow is like this:  

	# dlv debug channel.go
	Type 'help' for list of commands.
	(dlv) b channel.go:14
	Breakpoint 1 set at 0x401079 for main.main() ./channel.go:14
	(dlv) c
	> main.main() ./channel.go:14 (hits goroutine(1):1 total:1) (PC: 0x401079)
	     9:                         ch <- v
	    10:                 }
	    11:                 close(ch)
	    12:         }(ch)
	    13:
	=>  14:         for v := range ch {
	    15:                 fmt.Println(v)
	    16:         }
	    17:         fmt.Println("The channel is closed.")
	    18: }
	(dlv) n
	> main.main() ./channel.go:15 (PC: 0x4010c9)
	    10:                 }
	    11:                 close(ch)
	    12:         }(ch)
	    13:
	    14:         for v := range ch {
	=>  15:                 fmt.Println(v)
	    16:         }
	    17:         fmt.Println("The channel is closed.")
	    18: }
	(dlv) p v
	1
	(dlv) goroutine
	Thread 12380 at ./channel.go:15
	Goroutine 1:
	        Runtime: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x42a283)
	        User: ./channel.go:15 main.main (0x4010c9)
	        Go: /usr/local/go/src/runtime/asm_amd64.s:145 runtime.rt0_go (0x453321)
	(dlv) goroutines
	[5 goroutines]
	* Goroutine 1 - User: ./channel.go:15 main.main (0x4010c9)
	  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x42a283)
	  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x42a283)
	  Goroutine 4 - User: /usr/local/go/src/runtime/proc.go:263 runtime.gopark (0x42a283)
	  Goroutine 5 - User: ./channel.go:9 main.main.func1 (0x4013a8)
	(dlv)

Compared with `gdb`, `Delve` doesn't provide `start` command, so you need to set breakpoint first, then run `continue` command. You can see, `Delve` provides fruitful commands, e.g., you can check every goroutine status, so I think you should practice it frequently, and you will love it!

Happy debugging!