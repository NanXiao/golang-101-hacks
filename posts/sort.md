Sort
----
`sort` package defines an [interface](https://golang.org/pkg/sort/#Interface) whose name is `Interface`:  

	type Interface interface {  
	        // Len is the number of elements in the collection.  
	        Len() int  
	        // Less reports whether the element with  
	        // index i should sort before the element with index j.  
	        Less(i, j int) bool  
	        // Swap swaps the elements with indexes i and j.  
	        Swap(i, j int)  
	}

For slice, or any other collection types, provided that it implements the `Len()`, `Less` and `Swap` functions, you can use `sort.Sort()` function to arrange the elements in the order.  

Let's see the following example:  

	package main
	
	import (
		"fmt"
		"sort"
	)
	
	type command struct  {
		name string
	}
	
	type byName []command
	
	func (a byName) Len() int           { return len(a) }
	func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a byName) Less(i, j int) bool { return a[i].name < a[j].name }
	
	
	func main() {
		c := []command{
			{"breakpoint"},
			{"help"},
			{"args"},
			{"continue"},
		}
		fmt.Println("Before sorting: ", c)
		sort.Sort(byName(c))
		fmt.Println("After sorting: ", c)
	}

To avoid losing focus of demonstrating how to use `sort.Interface`, the `command` struct is simplified to only contain one `string` member: `name`. The comparison method (`Less`) is just contrasting the `name` in alphabetic order.  

Check the running result of the program:  

	Before sorting:  [{breakpoint} {help} {args} {continue}]
	After sorting:  [{args} {breakpoint} {continue} {help}]

We can see after sorting, the items in `c` are rearranged.