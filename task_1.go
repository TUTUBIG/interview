package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	stops := make([]chan struct{}, 0)

	for i := 0; i < 3; i++ {
		go func(i int) {
			t := time.NewTicker(time.Second)
			stop := make(chan struct{})
			stops = append(stops, stop)
			for {
				select {
				case <-t.C:
					job(i)
				case <-ctx.Done():
					stop <- struct{}{}
				}
			}
		}(i)
	}

	<-shutdown
	cancel()
	for _, stop := range stops {
		<-stop
	}

	fmt.Println("done")

	return
}

func job(i int) {
	now := time.Now().Unix()
	fmt.Println("job start", i, now)
	time.Sleep(5 * time.Second)
	fmt.Println("job end", i, now)
}
