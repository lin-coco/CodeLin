package lin

import "errors"

type Queue struct {
	slice []int
	cap   int
}

func NewQueue(size int) *Queue {
	return &Queue{
		slice: make([]int, 0, size),
		cap:   size,
	}
}

// 出队列
func (q *Queue) poll() (int, error) {
	if len(q.slice) == 0 {
		return 0, errors.New("没有任何元素了")
	}
	res := q.slice[0]
	q.slice = q.slice[1:]
	return res, nil
}

// 入队列
func (q *Queue) offer(val int) error {
	if len(q.slice) >= q.cap {
		return errors.New("容量不够进不去啦")
	}
	q.slice = append([]int{val}, q.slice...)
	return nil
}

// 队尾元素
func (q *Queue) peek() (int, error) {
	if len(q.slice) == 0 {
		return 0, errors.New("没有任何元素了")
	}
	res := q.slice[len(q.slice)-1]
	return res, nil
}
