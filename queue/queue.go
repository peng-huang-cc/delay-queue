package queue

import (
	"errors"
	"fmt"
)

/**
一、队列
  1.队列是一个有序列表，可以用数组或链表来实现
  2.遵循先入先出的原则：先存入队列的数据，要先取出，后存入的要后取出

二、使用数组环形队列
  //对首=对首+1%对列最大值——————取模
  //队尾=队尾+1%对列最大值——————取模
  //队列长度=(队尾 +对列最大值-对首)%对列最大值——————取模
*/
type CycleQueue struct {
	Capacity int //队列容量
	length   int
	Front    int //队头
	Rear     int //队尾 D
	Data     []interface{}
}

// 返回q的元素个数， 也是队列的当前长度
func (q *CycleQueue) Length() int {
	return q.length
}

// 判断是否为空
func (q *CycleQueue) IsEmpty() bool {
	return q.Capacity == 0
}

//判断是否满
func (q *CycleQueue) IsFull() bool {
	return q.Capacity == q.length
}

// append element
func (q *CycleQueue) Push(elem interface{}) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}
	q.Data[q.Rear] = elem
	q.Rear = (q.Rear + 1) % q.length
	q.Capacity++
	return nil
}

// pop element
func (q *CycleQueue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	var elem = q.Data[q.Front]
	q.Data[q.Front] = nil
	q.Front = (q.Front + 1) % q.length // front指针向后移一位，若到最后则转到数组头部
	q.Capacity--
	return elem, nil
}

//遍历
func (q *CycleQueue) Traverse() {
	fmt.Println("--Traverse--")
	for i := q.Front; i != q.Rear; i = (i + 1) % q.length {
		fmt.Println(q.Data[i])
	}
	return
}

//清空
func (q *CycleQueue) Empty() {
	q.Front = 0
	q.Rear = 0
	q.Capacity = 0
	q.Data = make([]interface{}, q.length)
}

func (q *CycleQueue) Display() {
	fmt.Println("--Display--")
	fmt.Println("Front----", q.Front)
	fmt.Println("Rear----", q.Rear)
	fmt.Println("队列的长度：", q.Length())
	fmt.Println("队列是否为空：", q.IsEmpty())
	fmt.Println("队列是否已满：", q.IsFull())
}


//初始化
func NewCycleQueue(length int) *CycleQueue {
	return &CycleQueue{
		Data:     make([]interface{}, length),
		Front:    0,
		Rear:     0,
		Capacity: 0,
		length:   length,
	}
}

