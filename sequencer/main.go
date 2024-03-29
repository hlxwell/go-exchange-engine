package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	disruptor "github.com/smartystreets/go-disruptor"
)

// Producer and Consumer MSG framework
const (
	RingBufferSize   = 1024 * 64
	RingBufferMask   = RingBufferSize - 1
	ReserveOne       = 1
	ReserveMany      = 16
	ReserveManyDelta = ReserveMany - 1
	DisruptorCleanup = time.Millisecond * 10
)

var ringBuffer = [RingBufferSize]int64{}

func main() {
	NumPublishers := 3 //runtime.NumCPU()
	totalIterations := int64(1000 * 1000 * 20)
	iterations := totalIterations / int64(NumPublishers)
	totalIterations = iterations * int64(NumPublishers)
	fmt.Printf("Total: %d,  Iterations: %d, Publisher: %d, Consumer: 1\n", totalIterations, iterations, NumPublishers)

	runtime.GOMAXPROCS(NumPublishers)
	var consumer = &countConsumer{TotalIterations: totalIterations, Count: 0}
	consumer.WG.Add(1)

	controller := disruptor.Configure(RingBufferSize).WithConsumerGroup(consumer).BuildShared()
	controller.Start()
	defer controller.Stop()

	var wg sync.WaitGroup
	wg.Add(NumPublishers + 1)

	var sendWG sync.WaitGroup
	sendWG.Add(NumPublishers)

	for i := 0; i < NumPublishers; i++ {
		go func() {
			writer := controller.Writer()
			wg.Done()
			wg.Wait()
			current := disruptor.InitialSequenceValue
			for current < totalIterations {
				current = writer.Reserve(ReserveMany)
				for j := current - ReserveMany + 1; j <= current; j++ {
					ringBuffer[j&RingBufferMask] = j
				}
				writer.Commit(current-ReserveMany, current)
			}
			sendWG.Done()
		}()
	}
	wg.Done()

	t := time.Now().UnixNano()
	wg.Wait() //waiting for ready as a barrier
	fmt.Println("start to publish")

	sendWG.Wait()
	fmt.Println("Finished to publish")

	consumer.WG.Wait()
	fmt.Println("Finished to consume") //waiting for consumer

	t = (time.Now().UnixNano() - t) / 1000000 //ms
	fmt.Printf("opsPerSecond: %d\n", totalIterations*1000/t)
}

type countConsumer struct {
	Count           int64
	TotalIterations int64
	WG              sync.WaitGroup
}

func (cc *countConsumer) Consume(lower, upper int64) {
	for lower <= upper {
		message := ringBuffer[lower&RingBufferMask]
		if message != lower {
			warning := fmt.Sprintf("\nRace condition--Sequence: %d, Message: %d\n", lower, message)
			fmt.Printf(warning)
			panic(warning)
		}
		lower++
		cc.Count++
		//fmt.Printf("count: %d, message: %d\n", cc.Count-1, message)
		if cc.Count == cc.TotalIterations {
			cc.WG.Done()
			return
		}
	}
}
