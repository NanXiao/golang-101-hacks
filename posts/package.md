# Package
----
In `Go`, the packages can be divided into `2` categories:  

(1) `main` package: is used to generate the executable binary, and the `main` function is the entry point of the program. Take `hello.go` as an example:  

	package main
	
	import "greet"
	
	func main() {
		greet.Greet()
	}
  

(2) This category can also include `2` types:  
 
a) Library package: is used to generate the object files that can be reused by others. Take `greet.go` as an example:  

	package greet
	
	import "fmt"
	
	func Greet() {
		fmt.Println("Hello 中国!")
	}

b) Some other packages for special purposes, such as testing.

Nearly every program needs `Go` standard (`$GOROOT`) or third-pary (`$GOPATH`) packages. To use them, you should use `import` statement:  

	import "fmt"
	import "github.com/NanXiao/stack" 
Or:  

	import (
		"fmt"
		"github.com/NanXiao/stack"
	)
In the above examples, the "`fmt`" and "`github.com/NanXiao/stack`" are called `import path`, which is used to find the relevant package.  

You may also see the following cases:  

	import m "lib/math" // use m as the math package name
	import . "lib/math" // Omit package name when using math package

If the `go install` command can't find the specified package, it will complain the error messages like this:  

	... : cannot find package "xxxx" in any of:
	        /usr/local/go/src/xxxx (from $GOROOT)
	        /root/gowork/src/xxxx (from $GOPATH)

To avoid library conflicts, you'd better make your own packages' path the only one in the world: E.g., your `github` repository destination:

	 github.com/NanXiao/...
**Conventionally**, your package name should be same with the last item in `import path`; it is a good coding habit though not a must.  

Reference:  
[The Go Programming Language](http://www.gopl.io/).



 
