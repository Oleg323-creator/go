package main

import (
	"context"
	"geckoapi1/pkg/connectors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Runner struct {
	ConnectorType string
	PollingRate   int
	RateFrom      string
	RateTo        string
}

func NewRunner(conType string, pollRate int, from string, to string) *Runner {
	return &Runner{
		ConnectorType: conType,
		PollingRate:   pollRate,
		RateFrom:      from,
		RateTo:        to,
	}
}

func (r *Runner) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Duration(r.PollingRate) * time.Second)
	log.Println("Starting:")
	for {
		select {
		case <-ctx.Done():
			log.Println("Finishing:")
			return
		case <-ticker.C:
			conn, err := connectors.NewConnector(r.ConnectorType)
			if err != nil {
				return
			}
			rates, err := conn.GetRates(r.RateFrom, r.RateTo)
			if err != nil {
				return
			}

			log.Println(time.Now().Unix(), r.RateFrom, rates, r.ConnectorType)
			time.Sleep(1 * time.Second)
			continue
		}
	}
}

func main() {
	geckoRun := NewRunner(connectors.CoingeckoType, 3, "BTC", "USDT")
	cryptCompRun := NewRunner(connectors.CryptoCompType, 1, "ETH", "USDT")

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go geckoRun.Run(ctx, wg)
	go cryptCompRun.Run(ctx, wg)

	<-stop

	cancel()
	wg.Wait()
}
