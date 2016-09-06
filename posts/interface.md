Interface
----
Interface is a reference type which contains some method definitions. Any type which implements all the methods defined by a reference type will satisfy this interface type automatically. Through interface, you can approach object-oriented programming. Check the following example:  

	package main
	
	import "fmt"
	
	type Foo interface {
		foo()
	}
	
	type A struct {
	}
	
	func (a A) foo() {
		fmt.Println("A foo")
	}
	
	func (a A) bar() {
		fmt.Println("A bar")
	}
	
	func callFoo(f Foo) {
		f.foo()
	}
	
	func main() {
		var a A
		callFoo(a)
	}
The running result is:  

	A foo
Let's analyze the code detailedly:  

(1)  

	type Foo interface {
		foo()
	}
The above code defines a interface `Foo` which has only one method: `foo()`.  

(2)  

	type A struct {
	}
	
	func (a A) foo() {
		fmt.Println("A foo")
	}
	
	func (a A) bar() {
		fmt.Println("A bar")
	}
struct `A` has `2` methods: `foo()` and `bar()`. Since it already implements `foo()` method, it satisfies `Foo` interface.  

(3)  

	func callFoo(f Foo) {
		f.foo()
	}
	
	func main() {
		var a A
		callFoo(a)
	}
`callFoo` requires a variable whose type is `Foo` interface, and passing `A` is OK. The `callFoo` will use `A`'s `foo()` method, and "`A foo`" is printed.  

Let's change the `main()` function:  

	func main() {
		var a A
		callFoo(&a)
	}
This time, the argument of `callFoo()` is `&a`, whose type is `*A`. Compile and run the program, you may find it also outputs: "`A foo`". So `*A` type has all the methods which `A` has. But the reverse is not true:  

	package main
	
	import "fmt"
	
	type Foo interface {
		foo()
	}
	
	type A struct {
	}
	
	func (a *A) foo() {
		fmt.Println("A foo")
	}
	
	func (a *A) bar() {
		fmt.Println("A bar")
	}
	
	func callFoo(f Foo) {
		f.foo()
	}
	
	func main() {
		var a A
		callFoo(a)
	}
Compile the program:  

	example.go:26: cannot use a (type A) as type Foo in argument to callFoo:
	A does not implement Foo (foo method has pointer receiver)

You can see also `*A` type has implemented `foo()` and `bar()` methods, it doesn't mean `A` type has both methods by default.  

BTW, every type satisfies the empty interface: `interface{}`.  

The interface type is actually a tuple which contains `2` elements: `<type, value>`, `type` identifies the type of the variable which stores in the interface while `value` points to the actual value. The default value of an interface type is `nil`, which means both `type` and `value` are `nil`: `<nil, nil>`. When you check whether an interface is empty or not:  

	var err error
	if err != nil {
		...
	}  
You must remember only if both `type` and `value` are `nil` means the interface value is `nil`.  

Reference:  
[The Go Programming Language](http://www.gopl.io/).



 