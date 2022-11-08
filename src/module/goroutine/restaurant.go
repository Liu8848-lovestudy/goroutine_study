package goroutine

var Num = 100

type restaurant interface {
	Action(i int, ch1 chan int)
}
