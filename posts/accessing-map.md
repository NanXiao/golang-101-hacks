# Accessing map
----
Map is a reference type which points to a hash table, and you can use it to construct a "key-value" database which is very handy in practice programming. E.g., the following code will calculate the count of every element in a slice:  

	package main
	
	import (
		"fmt"
	)
	
	func main() {
		s := []int{1, 1, 2, 2, 3, 3, 3}
		m := make(map[int]int)
	
		for _, v := range s {
			m[v]++
		}
	
		for key, value := range m {
			fmt.Printf("%d occurs %d times\n", key, value)
		}
	} 

The output is like this:  

	3 occurs 3 times
	1 occurs 2 times
	2 occurs 2 times
Moreover, according to [Go spec](https://golang.org/ref/spec#Map_types): "A map is an **unordered** group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.". So if you run the above program another time, the output may be different:  

	2 occurs 2 times
	3 occurs 3 times
	1 occurs 2 times
You must not presume the element order of a map.  

The key type of the map must can be compared with "`==`" operator: the built-in types, such as int, string, etc, satisfy this requirement; while slice not. For struct type, if its members all can be compared by "`==`" operator, then this struct can also be used as key.  

When you access a non-exist key of the map, the map will return the `nil` value of the element. I.e.:  

	package main
	
	import (
		"fmt"
	)
	
	func main() {
		m := make(map[int]bool)
	
		m[0] = false
		m[1] = true
	
		fmt.Println(m[0], m[1], m[2])
	}
The output is:  

	false true false

the value of `m[0]` and `m[2]` are both `false`, so you can't discriminate whether the key is really in map or not. The solution is to use “comma ok” method:  

	value, ok := map[key]
if the key does exit, `ok` will be `true`; else `ok` will be `false`.  

Sometimes, you may not care the values of the map, and use map just as a set. In this case, you can declare the value type as an empty struct: `struct{}`. An example is like this:  

	package main
	
	import (
		"fmt"
	)
	
	func check(m map[int]struct{}, k int) {
		if _, ok := m[k]; ok {
			fmt.Printf("%d is a valid key\n", k)
		}
	}
	func main() {
		m := make(map[int]struct{})
		m[0] = struct{}{}
		m[1] = struct{}{}
	
		for  i := 0; i <=2; i++ {
			check(m, i)
		}
	}  

The output is:

	0 is a valid key
	1 is a valid key
Using built-in `delete` function, you can delete a entry in the map, even the key doesn't exist:  

	delete(map, key)

References:  
[Effective Go](https://golang.org/doc/effective_go.html);  
[The Go Programming Language Specification](https://golang.org/ref/spec);  
[The Go Programming Language](http://www.gopl.io/).