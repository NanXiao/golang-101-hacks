"go get" command
----
"`go get`" command is the standard way of downloading and installing packages and related dependencies, and let's check the particulars of it through an example:  
(1) Create a [playstack](https://github.com/NanXiao/playstack) repository in github;  
(2) There is a `LICENSE` file and `play` directory in `playstack` folder;  
(3) The `play` directory includes one `main.go` file:  

	package main
	
	import (
		"fmt"
		"github.com/NanXiao/stack"
	)
	
	func main() {
		s := stack.New()
		s.Push(0)
		s.Push(1)
		s.Pop()
		fmt.Println(s)
	}
The `main` package has one dependency package: [stack](https://github.com/NanXiao/stack). Actually, the `main()` function doesn't play anything meaningful, and we just use this project as a sample. So the directory structure of `playstack` looks like this:  

	.
	├── LICENSE
	└── play
	    └── main.go
	
	1 directory, 2 files
 
Clean the `$GOPATH` directory, and use "`go get`" command to download `playstack`:  

	# tree
	.
	
	0 directories, 0 files
	# go get github.com/NanXiao/playstack
	package github.com/NanXiao/playstack: no buildable Go source files in /root/gocode/src/github.com/NanXiao/playstack

"`go get`" command complains "`no buildable Go source files in ...`", and it is because the objects which "`go get`" works are **packages**, not **repositories**.There is no `*.go` source files in `playstack`, so it is not a valid package.  

Tidy up `$GOPATH` folder, and execute "`go get github.com/NanXiao/playstack/play`" instead:  

	# tree
	.
	
	0 directories, 0 files
	# go get github.com/NanXiao/playstack/play
	# tree
	.
	├── bin
	│   └── play
	├── pkg
	│   └── linux_amd64
	│       └── github.com
	│           └── NanXiao
	│               └── stack.a
	└── src
	    └── github.com
	        └── NanXiao
	            ├── playstack
	            │   ├── LICENSE
	            │   └── play
	            │       └── main.go
	            └── stack
	                ├── LICENSE
	                ├── README.md
	                ├── stack.go
	                └── stack_test.go
	
	11 directories, 8 files
We can see not only `playstack` and its dependency (`stack`) are all downloaded, but also the command (`play`) and library (`stack`) are all installed in the right place.  

The mechanism behind "`go get`" command is it will fetch the repositories of packages and dependencies (E.g., use "`git clone`".), and you can check the detailed workflow by "`go get -x`":  

	# tree
	.
	
	0 directories, 0 files
	# go get -x github.com/NanXiao/playstack/play
	cd .
	git clone https://github.com/NanXiao/playstack /root/gocode/src/github.com/NanXiao/playstack
	cd /root/gocode/src/github.com/NanXiao/playstack
	git submodule update --init --recursive
	cd /root/gocode/src/github.com/NanXiao/playstack
	git show-ref
	cd /root/gocode/src/github.com/NanXiao/playstack
	git submodule update --init --recursive
	cd .
	git clone https://github.com/NanXiao/stack /root/gocode/src/github.com/NanXiao/stack
	cd /root/gocode/src/github.com/NanXiao/stack
	git submodule update --init --recursive
	cd /root/gocode/src/github.com/NanXiao/stack
	git show-ref
	cd /root/gocode/src/github.com/NanXiao/stack
	git submodule update --init --recursive
	WORK=/tmp/go-build054180753
	mkdir -p $WORK/github.com/NanXiao/stack/_obj/
	mkdir -p $WORK/github.com/NanXiao/
	cd /root/gocode/src/github.com/NanXiao/stack
	/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/NanXiao/stack.a -trimpath $WORK -p github.com/NanXiao/stack -complete -buildid de4d90fa03d8091e075c1be9952d691376f8f044 -D _/root/gocode/src/github.com/NanXiao/stack -I $WORK -pack ./stack.go
	mkdir -p /root/gocode/pkg/linux_amd64/github.com/NanXiao/
	mv $WORK/github.com/NanXiao/stack.a /root/gocode/pkg/linux_amd64/github.com/NanXiao/stack.a
	mkdir -p $WORK/github.com/NanXiao/playstack/play/_obj/
	mkdir -p $WORK/github.com/NanXiao/playstack/play/_obj/exe/
	cd /root/gocode/src/github.com/NanXiao/playstack/play
	/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/github.com/NanXiao/playstack/play.a -trimpath $WORK -p main -complete -buildid e9a3c02979f7c6e57ce4452278c52e3e0e6a48e8 -D _/root/gocode/src/github.com/NanXiao/playstack/play -I $WORK -I /root/gocode/pkg/linux_amd64 -pack ./main.go
	cd .
	/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/github.com/NanXiao/playstack/play/_obj/exe/a.out -L $WORK -L /root/gocode/pkg/linux_amd64 -extld=gcc -buildmode=exe -buildid=e9a3c02979f7c6e57ce4452278c52e3e0e6a48e8 $WORK/github.com/NanXiao/playstack/play.a
	mkdir -p /root/gocode/bin/
	mv $WORK/github.com/NanXiao/playstack/play/_obj/exe/a.out /root/gocode/bin/play
From the above output, we can see `playstack` repository is cloned first, then `stack`, At last the compilation and installation are executed.  

If you only want to download the source files, and not compile and install, using "`go get -d`" command:  
 
	# tree
	.
	
	0 directories, 0 files
	# go get -d github.com/NanXiao/playstack/play
	# tree
	.
	└── src
	    └── github.com
	        └── NanXiao
	            ├── playstack
	            │   ├── LICENSE
	            │   └── play
	            │       └── main.go
	            └── stack
	                ├── LICENSE
	                ├── README.md
	                ├── stack.go
	                └── stack_test.go
	
	6 directories, 6 files

You can also use "`go get -u`" to update packages and their dependencies.
 
Reference:  
[Command go](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies);  
[How does "go get" command know which files should be downloaded?](https://groups.google.com/forum/#!topic/golang-nuts/-V9QR8ncf4w).