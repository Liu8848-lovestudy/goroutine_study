package goroutine

import "sync"

type ReFactory struct {
}

func (r ReFactory) Create(types string, wg1 *sync.WaitGroup) restaurant {
	if types == "cooker" {
		return Cooker{}
	} else if types == "customer" {
		return Customer{
			wg: wg1,
		}
	} else {
		return nil
	}
}
