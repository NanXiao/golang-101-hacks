Decorate types to implement io.Reader interface
----
The [io package](https://golang.org/pkg/io/) has provided a bunch of handy read functions and methods, but unfortunately, they all require the arguments satisfy [io.Reader](https://golang.org/pkg/io/#Reader) interface. See the following example:  

	package main
	
	import (
		"fmt"
		"io"
	)
	
	func main() {
		s := "Hello, world!"
		p := make([]byte, len(s))
		if _, err := io.ReadFull(s, p); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", p)
		}
	} 
Compile above program and an error is generated:  

	read.go:11: cannot use s (type string) as type io.Reader in argument to io.ReadFull:
        string does not implement io.Reader (missing Read method)
The [io.ReadFull](https://golang.org/pkg/io/#ReadFull) function requires the argument should be compliance with `io.Reader`, but `string` type doesn't provide `Read()` method, so we need to do some tricks on `s` variable. Modify `io.ReadFull(s, p)` into `io.ReadFull(strings.NewReader(s), p)`:  

	package main

	import (
		"fmt"
		"io"
		"strings"
	)
	
	func main() {
		s := "Hello, world!"
		p := make([]byte, len(s))
		if _, err := io.ReadFull(strings.NewReader(s), p); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", p)
		}
	}
This time, the compilation is OK, and the running result is:  

	Hello, world!
[strings.NewReader](https://golang.org/pkg/strings/#NewReader) function converts a `string` into a [strings.Reader](https://golang.org/pkg/bytes/#Reader) struct which supplies a [read](https://golang.org/pkg/bytes/#Reader.Read) method:  

	func (r *Reader) Read(b []byte) (n int, err error)  
Besides `string`, another common operation is to use [bytes.NewReader](https://golang.org/pkg/bytes/#NewReader) to convert a byte slice into a [bytes.Reader](https://golang.org/pkg/bytes/#Reader) struct which satisfies `io.Reader` interface. Do some modifications on the above example:  

	 package main

	import (
		"bytes"
		"fmt"
		"io"
		"strings"
	)
	
	func main() {
		s := "Hello, world!"
		p := make([]byte, len(s))
		if _, err := io.ReadFull(strings.NewReader(s), p); err != nil {
			fmt.Println(err)
		}
	
		r := bytes.NewReader(p)
		if b, err := r.ReadByte(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%c\n", b)
		}
	}
`bytes.NewReader` converts the `p` slice into a `bytes.Reader` struct. The output is like this:  

	H