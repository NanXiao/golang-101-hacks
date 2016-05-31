error vs errors
----
Handling errors is a crucial part of writing robust programs. When scanning the `Go` packages, it is not rare to see APIs which have multiple return values with an error among them. For example:   

>func Open(name string) (*File, error)  
>
>Open opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.

And the idiomatic method of using `os.Open` function is like this:  

	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	} 
	defer file.Close()
So to implement resilient `Go` programs, how to generate and deal with errors is a required course.  

`Go` provides both `error` and `errors`, and you shouldn't mix up them. `error` is a built-in interface type:  

	type error interface {
        Error() string
	}
So for any type, as long as it implements `Error() string` method, it will satisfy `error` interface automatically. `errors` is one of my favorite packages since it is very simple (The life will definitely be easier if every package is similar to `errors`!). Removing the comments, the amount of core code lines is very small:    

    package errors
    
    func New(text string) error {
    	return &errorString{text}
    }
    
    type errorString struct {
    	s string
    }
    
    func (e *errorString) Error() string {
    	return e.s
    }
The `New` function in `errors` package returns an `errorString` struct which complies with `error` interface. Check the following example:  

	package main
	
	import (
		"errors"
		"fmt"
	)
	
	func maxElem(s []int) (int, error) {
		if len(s) == 0 {
			return 0, errors.New("The slice must be non-empty!")
		}
	
		max := s[0]
		for _, v := range s[1:] {
			if v > max {
				max = v
			}
		}
		return max, nil
	}
	
	func main() {
		s := []int{}
		_, err := maxElem(s)
		if err != nil {
			fmt.Println(err)
		}
	} 

The execution result is here:  

	The slice must be non-empty!
In real life, you may prefer to use `Errorf` function defined in `fmt` package to create `error` interface, rather than use `errors.New()` directly:  

>func Errorf(format string, a ...interface{}) error  
>
>Errorf formats according to a format specifier and returns the string as a value that satisfies error.

So the above code can be refactored as follows:  

	func maxElem(s []int) (int, error) {
		......
		if len(s) == 0 {
			return 0, fmt.Errorf("The slice must be non-empty!")
		}
		......
	}

References:  
[The Go Programming Language](http://www.gopl.io/).