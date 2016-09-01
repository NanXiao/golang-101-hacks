# defer
----
The `defer` statement is used to postpone a function call executed immediately before the surrounding function returns. The common uses of `defer` include releasing resources (i.e., unlock the mutex, close file handle.), do some tracing(i.e., record the running time of function), etc. E.g., an ordinary accessing global variable exclusively is like this:  

	var mu sync.Mutex
	var m = make(map[string]int)
	
	func lookup(key string) int {
		mu.Lock()
		v := m[key]
		mu.Unlock()
		return v
	}
An equivalent but concise format using `defer` is as follow:  
	
	var mu sync.Mutex
	var m = make(map[string]int)
	
	func lookup(key string) int {
		mu.Lock()
		defer mu.Unlock()
		return m[key]
	}
You can see this style is more simpler and easier to comprehend.  

The `defer` statements are executed in Last-In-First-Out sequence, which means the functions in latter `defer` statements run before their previous buddies. Check the following example:  

	package main
	
	import "fmt"
	
	func main()  {
		defer fmt.Println("First")
		defer fmt.Println("Last")
	}

The running result is here:  

	Last
	First

Although the function in `defer` statement runs very late, the parameters of the function are evaluated when the `defer` statement is executed.  

	package main
	
	import "fmt"
	
	func main()  {
		i := 10
		defer fmt.Println(i)
		i = 100
	}
The running result is here: 

	10
Besides, if the deferred function is a function literal, it can also modify the return value:  

	package main
	
	import "fmt"
	
	func modify() (result int) {
		defer func(){result = 1000}()
		return 100
	}
	
	func main()  {
		fmt.Println(modify())
	}  

The value printed is `1000`, not `100`.  

References:  
[The Go Programming Language](http://www.gopl.io/);  
[Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover).
  