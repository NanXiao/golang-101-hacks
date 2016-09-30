# Use sync.WaitGroup to synchronize goroutines
----
(This post is a modification edition of [Use sync.WaitGroup in Golang](http://nanxiao.me/en/use-sync-waitgroup-in-golang/)).  

[sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup) provides a goroutine synchronization mechanism, and used for waiting for a collection of goroutines to finish. In the internal of `sync.WaitGroup` struct, there is a `counter` which records how many goroutines need to be waited are living now.  

`sync.WaitGroup` provides `3` methods: `Add`, `Done` and `Wait`. `Add` method is used to identify how many goroutines need to be waited, and it will add `counter` value. When a goroutine exits, it must call `Done`, and it will decrease `counter` value by `1`. The `main` goroutine blocks on `Wait`, once the `counter` becomes `0`, the `Wait` will return, and main goroutine can continue to run.

Let’s see an example:

	package main
	
	import (
	    "sync"
	    "time"
	    "fmt"
	)
	
	func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
	    defer wg.Done()
	    time.Sleep(sec * time.Second)
	    fmt.Println("goroutine exit")
	}
	
	func main() {
	    var wg sync.WaitGroup
	
	    wg.Add(2)
	    go sleepFun(1, &wg)
	    go sleepFun(3, &wg)
	    wg.Wait()
	    fmt.Println("Main goroutine exit")
	
	}
Because the `main` goroutine need to wait `2` goroutines, so the argument for `wg.Add` is `2`. The execution result is like this:

	goroutine exit
	goroutine exit
	Main goroutine exit
 