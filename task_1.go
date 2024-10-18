package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	var shutdown chan os.Signal
	signal.Notify(shutdown, os.Interrupt)

	for i := 0; i < 3; i++ {
		go func() {}()
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
