Let's tidy up the `$GOPATH` directory and only keep `Golang` source code files left over:  

	# tree
	.
	├── bin
	├── pkg
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        └── hello.go
	
	5 directories, 2 files
The `greet.go` is `greet` package which just provides one function:  

	# cat src/greet/greet.go
	package greet
	
	import "fmt"
	
	func Greet() {
	        fmt.Println("Hello 中国!")
	}

While `hello.go` is a `main` package which takes advantage of `greet` and can be built into an executable binary:  

	# cat src/hello/hello.go
	package main
	
	import "greet"
	
	func main() {
	        greet.Greet()
	}

(1) Enter the `src/hello` directory, and execute `go build` command:  

	# pwd
	/root/gowork/src/hello
	# go build
	# ls
	hello  hello.go
	# ./hello
	Hello 中国!

We can see a fresh `hello` command is created in the current directory.  

Check the `$GOPATH` directory:  

	# tree
	.
	├── bin
	├── pkg
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        ├── hello
	        └── hello.go
	
	5 directories, 3 files

Compared before executing `go build`, there is only a final executable command more.  

(2) Clear the `$GOPATH` directory again:  

	# tree
	.
	├── bin
	├── pkg
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        └── hello.go
	
	5 directories, 2 files

Running `go install` in `hello` directory:  

	# pwd
	/root/gowork/src/hello
	# go install
	#

Check the `$GOPATH` directory now: 

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

Not only the `hello` command is generated and put into `bin` directory, but also the `greet.a` is in the `pkg/linux_amd64`. While the `src` folder keeps clean with only source code files in it and unchanged.  

(3) There is `-i` flag in `go build` command which will install the packages that are dependencies of the target, but won't install the target. Let's check it:  

	# tree
	.
	├── bin
	├── pkg
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        └── hello.go
	
	5 directories, 2 files

Run `go build -i` under `hello` directory:  

	# pwd
	#/root/gowork/src/hello
	# go build -i
  
Check `$GOPATH` now:  

	# tree
	.
	├── bin
	├── pkg
	│   └── linux_amd64
	│       └── greet.a
	└── src
	    ├── greet
	    │   └── greet.go
	    └── hello
	        ├── hello
	        └── hello.go
Except a `hello` command in `src/hello` directory, a `greet.a` library is also generated in `pkg/linux_amd64` too.  

(4) By default, the `go build` uses the directory's name as the compiled binary's name, to modify it, you can use `-o` flag:  

	# pwd
	/root/gowork/src/hello
	# go build -o he
	# ls
	he  hello.go

Now, the command is `he`, not `hello`.

