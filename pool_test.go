package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrencyPool_Run(t *testing.T) {
	var p ConcurrencyPool
	p.Initial(5)
	for i := 0; i < 100; i++ {
		go func(j int) {
			p.Run(func() error {
				fmt.Println("数字是:", j, "时间:", time.Now())
				time.Sleep(time.Second)
				return nil
			})
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
					p.Run(func() error {
						fmt.Println("数字是:", j, "时间:", time.Now())
						time.Sleep(time.Second)
						return nil
					})
				}(i)
			}
		}
	})
}
