# delay-queue

## 模拟延时队列的调用
   
   环 + 定时器
    
## 组成部分

### 环（循环队列）

```golang
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
```
 
### 定时器

```golang
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
```

### 消费

```golang
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
```