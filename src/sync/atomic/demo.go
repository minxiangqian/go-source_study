package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	_ "unsafe"
)

// atomicAddVal atomic包实现原子操作
func atomicAddVal() {
	var counter int32 // 定义一个 int32 类型的计数器
	group := sync.WaitGroup{}
	x := 10
	group.Add(x)
	// 启动 10 个 goroutine 并发地对计数器进行增加和减少操作
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				// 使用原子操作函数增加计数器
				atomic.AddInt32(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	//如果counter一直有变化 需要用atomic.LoadInt32获取值
	//fmt.Println("counter:", atomic.LoadInt32(&counter))
	fmt.Println("counter:", counter)

}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

//
//	mutexAddVal 锁实现原子操作

func mutexAddVal() {
	var counter Counter // 定义一个 Counter 类型的计数器
	group := sync.WaitGroup{}
	x := 10
	group.Add(x)
	// 启动 10 个 goroutine 并发地对计数器进行增加和减少操作
	for i := 0; i < x; i++ {

		go func() {
			for j := 0; j < 10000; j++ {
				counter.Increment() // 增加计数器
			}
			group.Done()
		}()
	}

	group.Wait()
	//如果counter一直有变化 需要用atomic.LoadInt32获取值
	//fmt.Println("counter:", atomic.LoadInt32(&counter))
	fmt.Println("counter:", counter.value)
}

func addVal() {
	var counter int32 // 定义一个 int32 类型的计数器
	group := sync.WaitGroup{}

	x := 10
	group.Add(x)
	// 启动 10 个 goroutine 并发地对计数器进行增加和减少操作
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				counter++
			}
			group.Done()

		}()
	}

	group.Wait()
	//如果counter一直有变化 需要用atomic.LoadInt32获取值
	//fmt.Println("counter:", atomic.LoadInt32(&counter))
	fmt.Println("counter:", counter)
}
