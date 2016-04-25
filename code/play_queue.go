package main

import (
	"fmt"
	"github.com/NanXiao/queue"
)

func main() {
	q := queue.New()
	q.Enqueue(1)
	fmt.Println(int(q.Dequeue()))
}
