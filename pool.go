package pool

import "log"

type ConcurrencyPool struct {
	capacity int
	worker chan struct{}
}

func (this *ConcurrencyPool) Initial(capacity int) *ConcurrencyPool {
	if capacity < 1 {
		log.Fatal("并发池的容量小于1")
	}
	this.capacity = capacity
	this.worker = make(chan struct{}, capacity)
	for i := 0; i < capacity; i++ {
		this.worker <- struct{}{}
	}
	return this
}

func (this *ConcurrencyPool) GetIdleCount() int{
	return len(this.worker)
}

func (this *ConcurrencyPool) Wait(){
	<- this.worker
}

func (this *ConcurrencyPool) Done(){
	this.worker <- struct{}{}
}

