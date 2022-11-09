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
	count := 0
	for {
		_, flag := <-ch
		if flag {
			count++
			time.Sleep(time.Duration(i) * time.Millisecond * 500)
			fmt.Printf("顾客%d吃到了%d个寿司\n", i, count)
			n--
		} else {
			c.wg.Done()
			flag1 = false
			return
		}

		if n <= 0 {
			fmt.Printf("顾客%d吃饱了,一共吃了%d个寿司\n", i, count)
			c.wg.Done()
			return
		}
	}
}
