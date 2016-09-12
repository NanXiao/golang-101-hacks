io.Writer interface
----
The inverse of [io.Reader](https://golang.org/pkg/io/#Reader) is [io.Writer](https://golang.org/pkg/io/#Writer) interface:  

	type Writer interface {
	        Write(p []byte) (n int, err error)
	}

Compared to `io.Reader`, since you no need to consider `io.EOF` error, the process of `Write` method is simple:  

(1) `err == nil`: All the data in `p` is written successfully;  
(2) `err != nil`: The data in `p` is partially or not written at all.  

Let's see an example:  

	package main

	import (
	        "log"
	        "os"
	)
	
	func main() {
	        f, err := os.Create("test.txt")
	        if err != nil {
	                log.Fatal(err)
	        }
	        defer f.Close()
	
	        if _, err = f.Write([]byte{'H', 'E', 'L', 'L', 'O'}); err != nil {
	                log.Fatal(err)
	        }
	}

After executing the program, the `test.txt` is created:  

	# cat test.txt
	HELLO