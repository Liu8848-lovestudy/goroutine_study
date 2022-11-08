package goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	wg *sync.WaitGroup
}

func (c Customer) Action(i int, ch chan int) {
	n := i
	for {
		_, flag := <-ch
		if flag {
			time.Sleep(time.Duration(i) * time.Millisecond * 500)
			fmt.Printf("顾客%d吃到了一个寿司\n", i)
			n--
		} else {
			c.wg.Done()
			flag1 = false
			return
		}

		if n <= 0 {
			fmt.Printf("顾客%d吃饱了\n", i)
			c.wg.Done()
			return
		}
	}
}
