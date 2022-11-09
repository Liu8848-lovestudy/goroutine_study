package main

import (
	"fmt"
	"goroutine_study/src/module/goroutine"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 100)
	rf := goroutine.ReFactory{}
	for i := 1; i < 6; i++ {
		go rf.Create("cooker", &wg).Action(i, ch)
	}
	for j := 1; j < 11; j++ {
		wg.Add(1)
		go rf.Create("customer", &wg).Action(j, ch)
	}
	//if goroutine.Num <= 0 {
	//	close(ch)
	//}
	wg.Wait()
	_, flag := <-ch
	if flag {
		close(ch)
	}

	fmt.Println("饭店打样！")

}
