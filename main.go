package main

import (
	"log"
	"time"

	Q "github.com/penghap/delay/queue"
)

var (
	SIZE  = 10
	queue = Q.NewCycleQueue(SIZE)
)

type Task struct {
}

func (task *Task) Run() {
	log.Printf("time now: %s, slot: %d, cycleNum: %d\n", time.Now().Format(time.RFC3339), slot, cycleNum)
}

type DelayQueueItem struct {
	CycleNum int
	Tasks    []*Task
}

type DelayQueue struct {
	Size int
	Data []map[int]*DelayQueueItem
}

func (dq *DelayQueue) Insert(task *Task, slot int, cycleNum int) bool {
	//根据消费时间计算消息应该放入的位置
	//var second = (int)(t.Unix() - time.Now().Unix())
	//var cycleNum = int(second / SIZE)
	//var slot =  second % SIZE

	//加入到延时队列中
	if item := dq.Data[slot]; item == nil {
		dq.Data[slot] = make(map[int]*DelayQueueItem)
	}

	if dq.Data[slot][cycleNum] == nil {
		dq.Data[slot][cycleNum] = &DelayQueueItem{
			CycleNum: cycleNum,
			Tasks:    make([]*Task, 0),
		}
	}

	queueItem := dq.Data[slot][cycleNum]
	queueItem.Tasks = append(queueItem.Tasks, task)

	return true
}

func (dq *DelayQueue) Consume(slot int, cycleNum int) bool {
	// 根据消费时间计算消息应该放入的位置
	// var second = (int)(t.Unix() - time.Now().Unix())
	// var cycleNum = int(second / SIZE)

	//var slot =  second % SIZE

	item := dq.Data[slot]
	if item == nil {
		return true
	}
	queueItem, exist := item[cycleNum]
	if !exist {
		return true
	}

	for _, task := range queueItem.Tasks {
		task.Run()
	}
	queueItem.Tasks = queueItem.Tasks[:0]
	return true
}

var dq = DelayQueue{
	Size: SIZE,
	Data: make([]map[int]*DelayQueueItem, SIZE),
}

var (
	slot     = 0
	cycleNum = 0
)

func Loop() {
	for range time.Tick(time.Second) {
		log.Printf("slot: %d, cycleNum: %d\n", slot, cycleNum)
		// 改为异步，不能占用时间
		go func() {
			dq.Consume(slot, cycleNum)
		}()
		slot++
		cycleNum = cycleNum + int(slot)/SIZE
		slot = slot % SIZE
	}
}

func main() {
	task0 := &Task{}

	dq.Insert(task0, 0, 0)

	dq.Consume(0, 0)
	//
	task := &Task{}
	dq.Insert(task, 1, 1)

	task1 := &Task{}
	dq.Insert(task1, 1, 2)

	task2 := &Task{}
	dq.Insert(task2, 2, 2)

	Loop()
}
