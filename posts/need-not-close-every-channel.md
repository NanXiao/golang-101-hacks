# Need not close every channel
----
You don't need to close channel after using it, and it can be recycled automatically by the garbage collector. The following quote is from [The Go Programming Language](http://www.gopl.io/):  

> You needn't close every channel when you've finished with it. It's only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent. A channel that the garbage collector determines to be unreachable will have its resources reclaimed whether or not it is closed. (Don't confuse this with the close operation for open files. It is important to call the Close method on every file when you've finished with it.)

References:  
[Is it OK to leave a channel open?](http://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open);  
[The Go Programming Language](http://www.gopl.io/).