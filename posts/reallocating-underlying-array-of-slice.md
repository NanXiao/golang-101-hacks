# Reallocating underlying array of slice
----
When appending data into slice, if the underlying array of the slice doesn't have enough space, a new array will be allocated. Then the elements in old array will be copied into this new memory, accompanied with adding new data behind. So when using `Go` built-in `append` function, you must always keep the idea that "the array may have been changed" in mind, and be very careful about it, otherwise, it may bite you!

Let me explain it through a contrived example:  

	package main
	
	import (
		"fmt"
	)
	
	func addTail(s []int)  {
		var ns [][]int
		for _, v := range []int{1, 2} {
			ns = append(ns, append(s, v))
		}
		fmt.Println(ns)
	}
	
	func main() {
		s1 := []int{0, 0}
		s2 := append(s1, 0)
	
		for _, v := range [][]int{s1, s2} {
			addTail(v)
		}
	}   
The `s1` is `[0, 0]`, and the `s2` is `[0, 0, 0]`; in `addTail` function, I want to add `1` and `2` behind the slice. So the wanted output is like this:  

	[[0 0 1] [0 0 2]]
	[[0 0 0 1] [0 0 0 2]]

But the actual result is:  

	[[0 0 1] [0 0 2]]
	[[0 0 0 2] [0 0 0 2]]

The operations on `s1` are successful, while `s2` not.  

Let's use `delve` to debug this issue and check the internal mechanism of slice: Add breakpoint on `addTail` function, and it is first hit when processing `s1`: 

	(dlv) n
	> main.addTail() ./slice.go:8 (PC: 0x401022)
	     3: import (
	     4:         "fmt"
	     5: )
	     6:
	     7: func addTail(s []int)  {
	=>   8:         var ns [][]int
	     9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	    12:         fmt.Println(ns)
	    13: }
	(dlv) p s
	[]int len: 2, cap: 2, [0,0]
	(dlv) p &s[0]
	(*int)(0xc82000a2a0)

 
The length and capacity of `s1` are both `2`, and the underlying array address is `0xc82000a2a0`, so what happened when executing the following statement:  

	ns = append(ns, append(s, v))
Since the length and capacity of `s1` are both `2`, there is no room for new buddy. To append a new value, a new array must be allocated, and it contains both `[0, 0]` from `s1` and the new value(`1` or `2`). You can consider `append(s, v)` generated an anonymous new slice, and it is appended in `ns`. We can check it after running "`ns = append(ns, append(s, v))`":  

	(dlv) n
	> main.addTail() ./slice.go:9 (PC: 0x401217)
	     4:         "fmt"
	     5: )
	     6:
	     7: func addTail(s []int)  {
	     8:         var ns [][]int
	=>   9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	    12:         fmt.Println(ns)
	    13: }
	    14:
	(dlv) p ns
	[][]int len: 1, cap: 1, [
	        [0,0,1],
	]
	(dlv) p ns[0]
	[]int len: 3, cap: 4, [0,0,1]
	(dlv) p &ns[0][0]
	(*int)(0xc82000e240)
	(dlv) p s
	[]int len: 2, cap: 2, [0,0]
	(dlv) p &s[0]
	(*int)(0xc82000a2a0)

We can see the length of anonymous slice is `3`, capacity is `4`,  and the underlying array address is `0xc82000e240`, different from `s1`'s (`0xc82000a2a0`). Continue executing until exit loop:  

	(dlv) n
	> main.addTail() ./slice.go:12 (PC: 0x40124c)
	     7: func addTail(s []int)  {
	     8:         var ns [][]int
	     9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	=>  12:         fmt.Println(ns)
	    13: }
	    14:
	    15: func main() {
	    16:         s1 := []int{0, 0}
	    17:         s2 := append(s1, 0)
	(dlv) p ns
	[][]int len: 2, cap: 2, [
	        [0,0,1],
	        [0,0,2],
	]
	(dlv) p &ns[0][0]
	(*int)(0xc82000e240)
	(dlv) p &ns[1][0]
	(*int)(0xc82000e280)
	(dlv) p &s[0]
	(*int)(0xc82000a2a0)
We can see `s1`, `ns[0]` and `ns[1]` have `3` independent array.  

Now, let's follow the same steps to check what happened on `s2`:  

	(dlv) n
	> main.addTail() ./slice.go:8 (PC: 0x401022)
	     3: import (
	     4:         "fmt"
	     5: )
	     6:
	     7: func addTail(s []int)  {
	=>   8:         var ns [][]int
	     9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	    12:         fmt.Println(ns)
	    13: }
	(dlv) p s
	[]int len: 3, cap: 4, [0,0,0]
	(dlv) p &s[0]
	(*int)(0xc82000e220)
	
The length of `s2` is `3`, and capacity is `4`, so there is one slot for adding new element. Check the `s2` and `ns`' values after executing "`ns = append(ns, append(s, v))`" the first time:  

	(dlv)
	> main.addTail() ./slice.go:9 (PC: 0x401217)
	     4:         "fmt"
	     5: )
	     6:
	     7: func addTail(s []int)  {
	     8:         var ns [][]int
	=>   9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	    12:         fmt.Println(ns)
	    13: }
	    14:
	(dlv) p ns
	[][]int len: 1, cap: 1, [
	        [0,0,0,1],
	]
	(dlv) p &ns[0][0]
	(*int)(0xc82000e220)
	(dlv) p s
	[]int len: 3, cap: 4, [0,0,0]
	(dlv) p &s[0]
	(*int)(0xc82000e220)
We can see the new anonymous slice's array address is also `0xc82000e220`, that's because the `s2` has enough space to hold new value, no new array is allocated. Check the `s2` and `ns` again after adding `2`:  

	(dlv)
	> main.addTail() ./slice.go:12 (PC: 0x40124c)
	     7: func addTail(s []int)  {
	     8:         var ns [][]int
	     9:         for _, v := range []int{1, 2} {
	    10:                 ns = append(ns, append(s, v))
	    11:         }
	=>  12:         fmt.Println(ns)
	    13: }
	    14:
	    15: func main() {
	    16:         s1 := []int{0, 0}
	    17:         s2 := append(s1, 0)
	(dlv) p ns
	[][]int len: 2, cap: 2, [
	        [0,0,0,2],
	        [0,0,0,2],
	]
	(dlv) p &ns[0][0]
	(*int)(0xc82000e220)
	(dlv) p &ns[1][0]
	(*int)(0xc82000e220)
	(dlv) p s
	[]int len: 3, cap: 4, [0,0,0]
	(dlv) p &s[0]
	(*int)(0xc82000e220)
All `3` slices point to the same array, so the later value(`2`) will override previous item(`1`).  

So in a conclusion, `append` is very tricky since it can modify the underlying array without noticing you. You must know the memory layout behind every slice clearly, else the slice can give you a big, unwanted surprise! 