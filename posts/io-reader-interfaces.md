io.Reader interface
----
`io.Reader` interface is a basic and very frequently-used interface:  

	type Reader interface {
	        Read(p []byte) (n int, err error)
	}
For every type who satisfies the `io.Reader` interface, you can imagine it's a pipe. Someone writes contents into one end of the pipe, and you can use `Read()` method which the type has provided to read content from the other end of the pipe. No matter it is a common file, a network socket, and so on. Only if it is compatible with `io.Reader` interface, I can read content of it.  

Let's see an example:  

	package main
	
	import (
		"fmt"
		"io"
		"log"
		"os"
	)
	
	func main() {
		file, err := os.Open("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		p := make([]byte, 4)
		for {
			if n, err := file.Read(p); n > 0 {
				fmt.Printf("Read %s\n", p[:n])
			} else if err != nil {
				if err == io.EOF {
					fmt.Println("Read all of the content.")
					break
				} else {
					log.Fatal(err)
				}
			} else /* n == 0 && err == nil */ {
				/* do nothing */
			}
		}
	}

You can see after issuing a `read()` call, there are `3` scenarios need to be considered:  

(1) `n > 0`: read valid contents; process it;  
(2) `n == 0 && err != nil`: if `err` is `io.EOF`, it means all the content have been read, and there is nothing left; else something unexpected happened, need to do special operations;  
(3) `n == 0 && err == nil`: according to [io package document](https://golang.org/pkg/io/#Reader), it means nothing happened, so no need to do anything.  

Create a `test.txt` file which only contains `5` bytes:  

	# cat test.txt
	abcde
Executing the program, and the result is like this:  

	Read abcd
	Read e
	Read all of the content.

Reference:  
[io package document](https://golang.org/pkg/io/#Reader).