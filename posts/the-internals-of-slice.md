# The internals of slice
----
There are `3` components of slice:  
a) `Pointer`: Points to the start position of slice in the underlying array;  
b) `length` (type is `int`): the number of the valid elements of the slice;  
b) `capacity` (type is `int`): the total number of slots of the slice.

Check the following code:  

	package main
	
	import (
		"fmt"
		"unsafe"
	)
	
	func main() {
		var s1 []int
		fmt.Println(unsafe.Sizeof(s1))
	}
The result is `24` on my `64-bit` system (The `pointer` and `int` both occupy `8` bytes).  

In the next example, I will use `gdb` to poke the internals of slice. The code is like this:  

	package main
	
	import "fmt"
	
	func main() {
	        s1 := make([]int, 3, 5)
	        copy(s1, []int{1, 2, 3})
	        fmt.Println(len(s1), cap(s1), &s1[0])
	
	        s1 = append(s1, 4)
	        fmt.Println(len(s1), cap(s1), &s1[0])
	
	        s2 := s1[1:]
	        fmt.Println(len(s2), cap(s2), &s2[0])
	}
	  
Use `gdb` to step into the code:  

	5       func main() {
	(gdb) n
	6               s1 := make([]int, 3, 5)
	(gdb)
	7               copy(s1, []int{1, 2, 3})
	(gdb)
	8               fmt.Println(len(s1), cap(s1), &s1[0])
	(gdb)
	3 5 0xc820010240
	
Before executing "`s1 = append(s1, 4)`", `fmt.Println` outputs the length(`3`), capacity(`5`) and the starting element address(`0xc820010240`) of the slice, let's check the memory layout of `s1`:  

	10              s1 = append(s1, 4)
	(gdb) p &s1
	$1 = (struct []int *) 0xc82003fe40
	(gdb) x/24xb 0xc82003fe40
	0xc82003fe40:   0x40    0x02    0x01    0x20    0xc8    0x00    0x00    0x00
	0xc82003fe48:   0x03    0x00    0x00    0x00    0x00    0x00    0x00    0x00
	0xc82003fe50:   0x05    0x00    0x00    0x00    0x00    0x00    0x00    0x00
	(gdb)
Through examining the memory content of `s1`(the start memory address is `0xc82003fe40`), we can see its content matches the output of `fmt.Println`.  

Continue executing, and check the result before "`s2 := s1[1:]`":  

	(gdb) n
	11              fmt.Println(len(s1), cap(s1), &s1[0])
	(gdb)
	4 5 0xc820010240
	13              s2 := s1[1:]
	(gdb) x/24xb 0xc82003fe40
	0xc82003fe40:   0x40    0x02    0x01    0x20    0xc8    0x00    0x00    0x00
	0xc82003fe48:   0x04    0x00    0x00    0x00    0x00    0x00    0x00    0x00
	0xc82003fe50:   0x05    0x00    0x00    0x00    0x00    0x00    0x00    0x00
We can see after appending a new element(`s1 = append(s1, 4)`), the length of `s1` is changed to `4`, but the capacity remains the original value.  

Let's check the internals of `s2`:  

	(gdb) n
	14              fmt.Println(len(s2), cap(s2), &s2[0])
	(gdb)
	3 4 0xc820010248
	15      }
	(gdb) p &s2
	$3 = (struct []int *) 0xc82003fe28
	(gdb) x/24hb 0xc82003fe28
	0xc82003fe28:   0x48    0x02    0x01    0x20    0xc8    0x00    0x00    0x00
	0xc82003fe30:   0x03    0x00    0x00    0x00    0x00    0x00    0x00    0x00
	0xc82003fe38:   0x04    0x00    0x00    0x00    0x00    0x00    0x00    0x00
The element start address of `s2` is `0xc820010248`, actually the second element of `s1`(`0xc82003fe40`), and the length(`3`) and capacity(`4`) are both one less than the counterparts of `s1`(`4` and `5` respectively).