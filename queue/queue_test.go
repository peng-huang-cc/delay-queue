package queue

import (
	"testing"
)

const (
	size = 2
)

func TestNewCycleQueue(t *testing.T) {
	q := NewCycleQueue(size)
	if len(q.Data) != size {
		t.Errorf("Data length(%d) should be (%d)", len(q.Data), size)
	}
	t.Logf("Data length(%d) equal (%d)", len(q.Data), size)
}

func TestIsEmpty(t *testing.T)  {
	q := NewCycleQueue(size)

	if b := q.IsEmpty(); b == false {
		t.Errorf("初始化队列后, 队列应该为空 (%t)", b)
	}
}

func TestPush(t *testing.T) {
	q := NewCycleQueue(size)

	for i := 0; i < size; i++ {
		q.Push(i)
	}

	if b := q.IsFull(); !b {
		t.Errorf("Queue should be full %d", len(q.Data))
	}

	if b := q.Capacity == size; !b {
		t.Errorf("Queue Capacity(%d) should be %d", q.Capacity, size)
	}

	if b := q.Rear == q.Front; !b {
		t.Errorf("Queue Rear(%d) should be Front(%d)", q.Rear, q.Front)
	}

	if e := q.Push(12); e != nil {
		t.Logf("Can not push element, %v", e)
	}

}


func TestPop(t *testing.T) {
	q := NewCycleQueue(size)

	for i := 0; i < size; i++ {
		q.Push(i)
	}


	for i := 0; i < size; i++ {
		q.Pop()
	}

	if b := 0 == q.Front; !b {
		t.Errorf("Queue Front(%d) should be Front(%d)", q.Rear, 0)
	}


	if b := q.Rear == q.Front; !b {
		t.Errorf("Queue Rear(%d) should be Front(%d)", q.Rear, q.Front)
	}
}