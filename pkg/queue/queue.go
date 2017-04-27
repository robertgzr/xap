package queue

import (
	"time"
)

type Track string

type Queue struct {
	L   List
	Pos int
}

func NewQeue() *Queue {
	return &Queue{
		L:   *NewList(),
		Pos: 0,
	}
}

func (q *Queue) Listen() {
	for {
		println("listening...")
		println(q.L.String())
		time.Sleep(5 * time.Second)
	}
}
