package goroutine

import (
	"fmt"
	"time"
)

var flag1 = true

type Cooker struct {
}

func (c Cooker) Action(i int, ch chan int) {

	for {

		if Num > 0 && flag1 {
			ch <- 1
			Num--
			time.Sleep(time.Duration(i) * time.Millisecond * 800)
			fmt.Printf("厨师%d号，做了一个寿司\n", i)
		}
		if Num <= 0 && flag1 {
			//once := sync.Once{}
			//once.Do(func() {
			//	close(ch)
			//	fmt.Println("今日寿司已出售完毕！")
			//})
			close(ch)
			flag1 = false
			return
		}
		if !flag1 {
			return
		}

	}

}
