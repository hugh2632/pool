package pool

import "log"

type ConcurrencyPool struct {
	ch   chan error
	idle chan struct{}
}

func (this *ConcurrencyPool) Initial(capacity int) *ConcurrencyPool {
	if capacity < 1 {
		log.Fatal("并发池的容量小于1")
	}
	this.ch = make(chan error)
	this.idle = make(chan struct{}, capacity)
	for i := 0; i < capacity; i++ {
		this.idle <- struct{}{}
	}
	go func() {
		for {
			select {
			case err := <-this.ch:
				if err != nil {
					log.Println(err.Error())
				}
				this.idle <- struct{}{}
			}
		}
	}()
	return this
}

func (this *ConcurrencyPool) Run(f func() error) {
	<-this.idle
	this.ch <- f()
}

