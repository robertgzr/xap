package main

import "rbg.re/robertgzr/xapper/pkg/queue"

func main() {
	q := queue.NewQeue()
	q.Listen()
}
