# Create Go workspace
----
Once the `Go` build environment is ready, the next step is to create workspace for development:  

(1) Set up a new empty directory:  

	# mkdir gowork

(2) Use a new environment variable `$GOPATH` to point it:  
	
	# cat /etc/profile
	......
	GOPATH=/root/gowork
	export GOPATH
	...... 

The workspace should contain `3` subdirectories:  

>src: contains the Go source code.  
>pkg: contains the package objects. You could think them as libraries which are used in linkage stage to generate the final executable files.  
>bin: contains the executable files. 

Let's see an example:

(1) Create a `src` directory in `$GOPATH`, which is `/root/gowork` in my system:  
 
	# mkdir src
	# tree
	.
	└── src
	
	1 directory, 0 files

(2) Since `Go` organizes source code using "`package`" concept , and every "`package`" should occupy a distinct directory, I create a `greet` directory in `src`:  

	# mkdir src/greet

Then create a new `Go` source code file (`greet.go`) in `src/greet`:  

	# cat src/greet/greet.go
	package greet
	
	import "fmt"
	
	func Greet() {
	        fmt.Println("Hello 中国!")
	}

You can consider this `greet` directory provides a `greet` package which can be used by other programs.  
 
(3) Create another package `hello` which utilizes the `greet` package:  

	# mkdir src/hello
	# cat src/hello/hello.go
	package main
	
	import "greet"
	
	func main() {
	        greet.Greet()
	}

You can see in `hello.go`, the `main` function calls `Greet` function offered by `greet` package.  

(4) Now our `$GOPATH` layout is like this:  

	# tree
	.
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        └── hello.go
	
	3 directories, 2 files

Let's compile and install `hello` package:  

	# go install hello

Check the `$GOPATH` layout again:  

	# tree
	.
	├── bin
	│   └── hello
	├── pkg
	│   └── linux_amd64
	│       └── greet.a
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        └── hello.go
	
	6 directories, 4 files

You can see the executable command `hello` is generated in `bin` folder. Because `hello` needs `greet` package's help, a `greet.a` object is also produced in `pkg` directory, but in system related subdirectory: `linux_amd64`.  

Run `hello` command:  

	# ./bin/hello
	Hello 中国!

Working as expected!

(5) You should add `$GOPATH/bin` to `$PATH` environment variable for facility:  

	PATH=$PATH:$GOPATH/bin
	export PATH

Then you can run `hello` directly:  

	# hello
	Hello 中国!

Reference:  
[How to Write Go Code](https://golang.org/doc/code.html).
