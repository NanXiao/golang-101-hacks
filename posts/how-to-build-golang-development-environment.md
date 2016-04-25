Build `Golang` development environment is always easy. Take `Linux` OS as an example (Because I work as a root user, so if you login as a non-root user, maybe you need `sudo` to execute some commands), what you should do is just download the binary package which matches your system from [here](https://golang.org/dl/), and uncompress it:
  
	# wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
	# tar -C /usr/local/ -xzf go1.6.2.linux-amd64.tar.gz
Now, there is an extra `go` directory under `/usr/local`. It's done! Too easy, right? Yes, but there are still some windup work to do:  

(1) To run `Golang` utilities (`go`, `gofmt`, etc) conveniently, you should add `/usr/local/go` into `$PATH` environment variable:  

	# cat /etc/profile  
	......
	PATH=$PATH:/usr/local/go/bin
	export PATH 
	......

(2) It is **strongly recommended** to install `Golang` in `/usr/local/go` under `*nix` and `c:\Go` under `Windows` since these default directories have already been embedded in `Golang` binary distributions. If you choose another directory, you **must** set `$GOROOT` environment variable:  

	# cat /etc/profile  
	......
	GOROOT=/path/to/go
	export GOROOT

So the `$GOROOT` is only need when you install `Golang` on a custom directory, not default.  

References:  
[Getting Started](https://golang.org/doc/install);  
[You don’t need to set GOROOT, really](http://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really).

  



   