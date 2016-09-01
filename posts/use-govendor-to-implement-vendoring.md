Use govendor to implement vendoring
----
The meaning of vendoring in `Go` is squeezing a project's all dependencies into its `vendor` directory. Since `Go 1.6`, if there is a `vendor` directory in current package  or its parent's directory, the dependency will be searched in `vendor` directory **first**. [Govendor](https://github.com/kardianos/govendor) is such a tool to help you make use of the `vendor` feature. In the following example, I will demonstrate how to use `govendor` step by step:

(1) To be more clear, I clean `$GOPATH` folder first:  

	# tree
	.
	
	0 directories, 0 files
(2) I still use [playstack](https://github.com/NanXiao/playstack) project to do a demo, download it:  

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
The `playstack` depends on another 3rd-party package: [stack](https://github.com/NanXiao/stack).  

(3) Install `govendor`:  

	# go get -u github.com/kardianos/govendor

(4) Change to `playstack` directory, and run "`govendor init`" command:  

	# cd src/github.com/NanXiao/playstack/
	# govendor init
	# tree
	.
	├── LICENSE
	├── play
	│   └── main.go
	└── vendor
	    └── vendor.json
	
	2 directories, 3 files
You can see there is an additional `vendor` folder which contains `vendor.json` file:  

	# cat vendor/vendor.json
	{
	        "comment": "",
	        "ignore": "test",
	        "package": [],
	        "rootPath": "github.com/NanXiao/playstack"
	}

(5) Execute "`govendor add +external`" command:  

	# govendor add +external
	# tree
	.
	├── LICENSE
	├── play
	│   └── main.go
	└── vendor
	    ├── github.com
	    │   └── NanXiao
	    │       └── stack
	    │           ├── LICENSE
	    │           ├── README.md
	    │           └── stack.go
	    └── vendor.json

Yeah, the `stack` project is copied to `vendor` directory now. Look at `vendor/vendor.json` file again:  

	# cat vendor/vendor.json
	{
	        "comment": "",
	        "ignore": "test",
	        "package": [
	                {
	                        "checksumSHA1": "3v5ClsvqF5lU/3E3c+1gf/zVeK0=",
	                        "path": "github.com/NanXiao/stack",
	                        "revision": "bfb214dbdb387d1c561b3b6f305ee0d8444c864b",
	                        "revisionTime": "2016-04-01T05:28:44Z"
	                }
	        ],
	        "rootPath": "github.com/NanXiao/playstack"
	}
The `stack` package info has been updated in `vendor/vendor.json` file.  

Notice: "`govendor add`" copies packages from `$GOPATH`, and you can use "`govendor fetch`" to download packages from network. You can verify it through removing `stack` package in `$GOPATH`, and execute "`govendor fetch github.com/NanXiao/stack`" command.  

(6) Update `playstack` in `github`:  

![image](https://raw.githubusercontent.com/NanXiao/golang-101-Hacks/master/images/govendor-playstack.JPG) 
  
This time, clean `$GOPATH` folder and run "`go get github.com/NanXiao/playstack/play`" again:  

	# go get github.com/NanXiao/playstack/play
	# tree
	.
	├── bin
	│   └── play
	├── pkg
	│   └── linux_amd64
	│       └── github.com
	│           └── NanXiao
	│               └── playstack
	│                   └── vendor
	│                       └── github.com
	│                           └── NanXiao
	│                               └── stack.a
	└── src
	    └── github.com
	        └── NanXiao
	            └── playstack
	                ├── LICENSE
	                ├── play
	                │   └── main.go
	                └── vendor
	                    ├── github.com
	                    │   └── NanXiao
	                    │       └── stack
	                    │           ├── LICENSE
	                    │           ├── README.md
	                    │           └── stack.go
	                    └── vendor.json
	
	18 directories, 8 files

Compared to previous case, it is no need to store `stack` in `$GOPATH/src/github.com/NanXiao` directory, since `playstack` has embedded it in its `vendor` folder.  

This is just a simple intro of `govendor`, for more commands' usages, you should visit its project [home page](https://github.com/kardianos/govendor). 

Reference:  
[What does the term “vendoring” or “to vendor” mean for Ruby on Rails?](http://stackoverflow.com/questions/11378921/what-does-the-term-vendoring-or-to-vendor-mean-for-ruby-on-rails);  
[Understanding and using the vendor folder](https://blog.gopheracademy.com/advent-2015/vendor-folder/);  
[Go Vendoring Beginner Tutorial](https://gocodecloud.com/blog/2016/03/29/go-vendoring-beginner-tutorial/).
