package lin

import "errors"

// 双向链表
type LinkedList struct {
	Start *Node
	End   *Node
	Size  int
}

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

func NewLinkList() *LinkedList {
	return &LinkedList{
		Start: nil,
		End:   nil,
		Size:  0,
	}
}

// 索引查找
func (l *LinkedList) Get(i int) (int, error) {
	if l.Size == 0 || i >= l.Size {
		return 0, errors.New("超出索引范围")
	}
	node := l.Start
	for i > 0 {
		node = node.Next
		i--
	}
	return node.Val, nil
}

// 值查找
func (l *LinkedList) Index(val int) int {
	node := l.Start
	i := 0
	for node != nil {
		if node.Val == val {
			return i
		}
		node = node.Next
		i++
	}
	return -1
}

// 值查找
func (l *LinkedList) LastIndex(val int) int {
	node := l.End
	i := l.Size - 1
	for node != nil {
		if node.Val == val {
			return i
		}
		node = node.Prev
		i--
	}
	return -1
}

// 插入
func (l *LinkedList) Add(val int) {
	newNode := &Node{
		Prev: nil,
		Next: nil,
		Val:  val,
	}
	if l.Size == 0 {
		l.Start = newNode
		l.End = l.Start
	} else {
		l.End.Next = newNode
		newNode.Prev = l.End
		l.End = l.End.Next
	}
	l.Size++
}

func (l *LinkedList) AddStart(val int) {
	newNode := &Node{
		Prev: nil,
		Next: nil,
		Val:  val,
	}
	if l.Size == 0 {
		l.Start = newNode
		l.End = l.Start
	} else {
		newNode.Next = l.Start
		l.Start.Prev = newNode
		l.Start = l.Start.Prev
	}
	l.Size++
}

func (l *LinkedList) AddIndex(i int, val int) error {
	if i > l.Size || i < 0 {
		return errors.New("不在索引范围内")
	}
	if i == l.Size {
		l.Add(val)
		return nil
	}
	if i == 0 {
		l.AddStart(val)
		return nil
	}
	newNode := &Node{
		Prev: nil,
		Next: nil,
		Val:  val,
	}
	j := 0
	tmpNode := l.Start
	for j < i-1 {
		tmpNode = tmpNode.Next
	}
	newNode.Next = tmpNode.Next
	tmpNode.Next.Prev = newNode
	newNode.Prev = tmpNode
	tmpNode.Next = newNode
	l.Size++
	return nil
}

// 删除
func (l *LinkedList) Delete(i int) error {
	if l.Size == 0 || i >= l.Size || i < 0 {
		return errors.New("不在索引范围内")
	}
	if i == 0 {
		l.Start = l.Start.Next
		return nil
	}
	if i == l.Size-1 {
		l.End = l.End.Prev
		return nil
	}
	j := 0
	tmpNode := l.Start
	for j < i-1 {
		tmpNode = tmpNode.Next
	}
	tmpNode.Next = tmpNode.Next.Next
	tmpNode.Next.Prev = nil
	if tmpNode.Next.Next != nil {
		tmpNode.Next.Next.Prev = tmpNode
	}
	return nil
}

// 替换
func (l *LinkedList) Replace(i int, val int) error {
	if l.Size == 0 || i >= l.Size || i < 0 {
		return errors.New("不在索引范围内")
	}
	j := 0
	node := l.Start
	for j < i {
		node = node.Next
	}
	node.Val = val
	return nil
}
