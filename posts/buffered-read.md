Buffered read
----
[bufio](https://golang.org/pkg/bufio/) package provides buffered read functions. Let's see an example:  

(1) Create a `test.txt` file first:  

	# cat test.txt
	abcd
	efg
	hijk
	lmn
You can see `test.txt` contains `4` lines.  

(2) See the following program:  

	package main
	
	import (
	        "bufio"
	        "bytes"
	        "fmt"
	        "io"
	        "io/ioutil"
	        "log"
	)
	
	func main() {
	        p, err := ioutil.ReadFile("test.txt")
	        if err != nil {
	                log.Fatal(err)
	        }
	
	        r := bufio.NewReader(bytes.NewReader(p))
	        for {
	                if s, err := r.ReadSlice('\n'); err == nil || err == io.EOF {
	                        fmt.Printf("%s", s)
	                        if err == io.EOF {
	                                break
	                        }
	                } else {
	                        log.Fatal(err)
	                }
	
	        }
	}

(a)  

	p, err := ioutil.ReadFile("test.txt")
`ioutil.ReadFile("test.txt")` reads the whole content of `test.txt` into a slice: `p`.  

(b)  

	r := bufio.NewReader(bytes.NewReader(p))

`bufio.NewReader(bytes.NewReader(p))` creates a [bufio.Reader](https://golang.org/pkg/bufio/#Reader) struct which implements buffered read function.  

(c)  

	for {
		if s, err := r.ReadSlice('\n'); err == nil || err == io.EOF {
			fmt.Printf("%s", s)
			if err == io.EOF {
				break
			}
		} else {
			log.Fatal(err)
		}

	}
		
Read and print each line.  

The running result is here:  

	abcd
	efg
	hijk
	lmn
We can also use [bufio.Scanner](https://golang.org/pkg/bufio/#Scanner) to implement above "print each line" function:  

	package main

	import (
	        "bufio"
	        "bytes"
	        "fmt"
	        "io/ioutil"
	        "log"
	)
	
	func main() {
	        p, err := ioutil.ReadFile("test.txt")
	        if err != nil {
	                log.Fatal(err)
	        }
	
	        s := bufio.NewScanner(bytes.NewReader(p))
	
	        for s.Scan() {
	                fmt.Println(s.Text())
	        }
	}


(a)  

	s := bufio.NewScanner(bytes.NewReader(p))
`bufio.NewScanner(bytes.NewReader(p))` creates a new [bufio.Scanner](https://golang.org/pkg/bufio/#Scanner) struct which splits the content by line by default.  

(b)  

	for s.Scan() {
		fmt.Println(s.Text())
	}
`s.Scan()` advances the `bufio.Scanner` to the next token (in this case, it is one optional carriage return followed by one mandatory newline), and we can use `s.Text()` function to get the content.  

We can also customize [SplitFunc](https://golang.org/pkg/bufio/#SplitFunc) function which doesn't separate content by line. Check the following code:  

	package main
	
	import (
	        "bufio"
	        "bytes"
	        "fmt"
	        "io/ioutil"
	        "log"
	)
	
	func main() {
	        p, err := ioutil.ReadFile("test.txt")
	        if err != nil {
	                log.Fatal(err)
	        }
	
	        s := bufio.NewScanner(bytes.NewReader(p))
	        split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	                for i := 0; i < len(data); i++ {
	                        if data[i] == 'h' {
	                                return i + 1, data[:i], nil
	                        }
	                }
	
	                return 0, data, bufio.ErrFinalToken
	        }
	        s.Split(split)
	        for s.Scan() {
	                fmt.Println(s.Text())
	        }
	}
The `split` function separates the content by "`h`", and the running result is:  

	abcd
	efg
	
	ijk
	lmn




