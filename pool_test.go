package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrencyPool_Run(t *testing.T) {
	var p ConcurrencyPool
	p.Initial(5)
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(p.GetIdleCount())
		}

	}()
	for i := 0; i < 5; i++ {
		go func(j int) {
			p.Wait()
			fmt.Println("数字是:", j, "时间:", time.Now())
			time.Sleep(time.Second * time.Duration(j))
			fmt.Println("数字是:", j, "OK时间:", time.Now())
			p.Done()
		}(i)
	}

	select {}
}

func BenchmarkConcurrencyPool_Run(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var p ConcurrencyPool
			p.Initial(5)
			for i := 0; i < 100; i++ {
				go func(j int) {
					p.Wait()
					fmt.Println("数字是:", j, "时间:", time.Now())
					time.Sleep(time.Second)
					p.Done()
				}(i)
			}
		}
	})
}

func TestConcurrencyPool_Run2(t *testing.T) {
	var p ConcurrencyPool
	p.Initial(5)
	t.Log(p.GetIdleCount())
}
