package pool

import "log"

type ConcurrencyPool struct {
	capacity int
	ch   chan error
	worker chan struct{}
}

func (this *ConcurrencyPool) Initial(capacity int) *ConcurrencyPool {
	if capacity < 1 {
		log.Fatal("并发池的容量小于1")
	}
	this.capacity = capacity
	this.ch = make(chan error)
	this.worker = make(chan struct{}, capacity)
	for i := 0; i < capacity; i++ {
		this.worker <- struct{}{}
	}
	go func() {
		for {
			select {
			case err := <-this.ch:
				if err != nil {
					log.Println(err.Error())
				}
				this.worker <- struct{}{}
			}
		}
	}()
	return this
}

func (this *ConcurrencyPool) GetIdleCount() int{
	return len(this.worker)
}

func (this *ConcurrencyPool) Run(f func() error) {
	<-this.worker
	this.ch <- f()
}

