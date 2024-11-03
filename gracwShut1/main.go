package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func tickerPractice(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	log.Println("начало работы")
	for {
		select {
		case <-ctx.Done():
			log.Println("завершение работы")
			return
		case <-ticker.C:
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				time.Sleep(1 * time.Second)
				continue
			}

		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	wg.Add(1)
	go tickerPractice(ctx, wg)

	<-stop
	log.Println("after stop")

	cancel()
	log.Println("after cancel")
	wg.Wait()
	log.Println("after Wait")
	fmt.Println("is ok")

}
