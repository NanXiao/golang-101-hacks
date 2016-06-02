init function
---
There is a `init()` function, as the name suggests, it will do some initialization work, such as initializing variables which may not be expressed easily, or calibrating program state. A file can contain one or more `init()` functions, as shown here:  

	package main

	import "fmt"
	
	var global int = 0
	
	func init() {
		global++
		fmt.Println("In first Init(), global is: ", global)
	}
	
	func init() {
		global++
		fmt.Println("In Second Init(), global is: ", global)
	}
	
	func main() {
		fmt.Println("In main(), global is: ", global)
	}

The execution result is like this:  

	In first Init(), global is:  1
	In Second Init(), global is:  2
	In main(), global is:  2
Since one package can contain multiple files, there may be many `init()` functions. You **should not** presume which file's `init()` functions are executed first. The only thing which is guaranteed is that the variables declared in package will be evaluated before all `init()` functions are executed in this package.  

See another example. The `$GOROOT/src` directory is like this:  

	# tree
	.
	├── foo
	│   └── foo.go
	└── play
	    └── main.go
There are `2` simple packages: `foo` and `play`. The `foo/foo.go` is here:  

	package foo
	
	import "fmt"
	
	var Global int
	
	func init() {
	        Global++
	        fmt.Println("foo init() is called, Global is: ", Global)
	}
While the `play/main.go` is:  

	package main

	import "foo"
	
	
	func main() {
	}
Build `play` command:  

	# go install play
	# play
	src/play/main.go:3: imported and not used: "foo"
The cause of this error is that `main.go` doesn't use any functions or variables exported by `foo` package. So if you just want an imported package's `init()` function is executed, and don't want to use package's other stuff, you should modify "`import "foo"`" to "`import _ "foo"`":  

	 package main

	import _ "foo"
	
	
	func main() {
	}
Now the build process will success, and the output of `play` command is like this:  

	# play
	foo init() is called, Global is:  1

References:  
[Effective Go](https://golang.org/doc/effective_go.html#init);  
[When is the init() function in go (golang) run?](http://stackoverflow.com/questions/24790175/when-is-the-init-function-in-go-golang-run).  