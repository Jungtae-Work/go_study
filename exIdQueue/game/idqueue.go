package game

import (
	"container/list"
	"fmt"
	"strings"
)

type IdQueue struct {
	v   *list.List
	max int
}

func NewIdQueue(start, end int) IdQueue {
	que := IdQueue{list.New(), end}
	for i := start; i <= end; i++ {
		que.v.PushBack(i)
	}
	return que
}

func (q *IdQueue) Push(id int) {
	q.v.PushBack(id)
}

func (q *IdQueue) Pop() int {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front).(int)
	} else {
		q.max++
	}
	return q.max
}

func (q *IdQueue) String() string {
	var tmp strings.Builder

	tmp.WriteString("v: [")
	for it := q.v.Front(); it != nil; it = it.Next() {
		tmp.WriteString(fmt.Sprintf(" %v", it.Value))
	}
	tmp.WriteString(fmt.Sprintf("](%d), max: %d", q.v.Len(), q.max))

	return tmp.String()
}
