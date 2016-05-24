# Check data race
----
"Data race" is a common but notorious issue in concurrency programs. sometimes it is difficult to debug and reproduce, especially in some big system, so this will make people very frustrated. Thankfully, the `Go` toolchain provides a `race detector` which can help us quickly spot and fix this kind of issue, and this can save our time even lives!  

Take the following classic "data race" program as an example:  

	package main
	
	import (
	        "fmt"
	        "sync"
	)
	
	var global int
	var wg sync.WaitGroup
	
	func count() {
	        defer wg.Done()
	        for i := 0; i < 10000; i++{
	                global++
	        }
	}
	
	func main() {
	        wg.Add(2)
	        go count()
	        go count()
	        wg.Wait()
	        fmt.Println(global)
	}
	
Two tasks increase `global` variable simultaneously, so the final value of `global` is non-deterministic. Using `race detector` to check it:  

	# go run -race race.go
	==================
	WARNING: DATA RACE
	Read by goroutine 7:
	  main.count()
	      /root/gocode/src/race.go:14 +0x6d
	
	Previous write by goroutine 6:
	  main.count()
	      /root/gocode/src/race.go:14 +0x89
	
	Goroutine 7 (running) created at:
	  main.main()
	      /root/gocode/src/race.go:21 +0x6d
	
	Goroutine 6 (running) created at:
	  main.main()
	      /root/gocode/src/race.go:20 +0x55
	==================
	19444
	Found 1 data race(s)
	exit status 66
Cool! the `race detector` finds the issue precisely, and it also provides the detailed tips of how to modifying it. Adding the lock of writing the `global` variable:  

	package main
	
	import (
		"fmt"
		"sync"
	)
	
	var global int
	var wg sync.WaitGroup
	var w sync.Mutex
	
	func count() {
		defer wg.Done()
		for i := 0; i < 10000; i++{
			w.Lock()
			global++
			w.Unlock()
		}
	}
	
	func main() {
		wg.Add(2)
		go count()
		go count()
		wg.Wait()
		fmt.Println(global)
	}  

 This time, race detector is calm:  

	# go run -race non_race.go
	20000

Please be accustomed to use this powerful tool frequently, you will appreciate it, I promise!

Reference:  
[Introducing the Go Race Detector](https://blog.golang.org/race-detector).