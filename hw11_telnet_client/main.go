package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

var timeout time.Duration

const (
	DefaultTimeout = 10
	minArgsCnt     = 2
)

type Flags struct {
	Timeout time.Duration
}

func main() {
	Flags := new(Flags)
	flag.DurationVar(&Flags.Timeout, "timeout", DefaultTimeout, "timeout for client in seconds")
	flag.Parse()
	args := flag.Args()

	if len(args) != minArgsCnt {
		logger.Fatalf("not enough args passed. Expected %d, got %d", minArgsCnt, len(args))
	}

	c := NewTelnetClient(net.JoinHostPort(args[0], args[1]), Flags.Timeout, os.Stdin, os.Stdout)
	if err := c.Connect(); err != nil {
		logger.Fatalf("error connecting to host: %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := c.Send(); err != nil {
			logger.Errorf("error sending to host: %s", err)
		} else {
			logger.Infof("EOF")
		}
		cancel()
	}()

	go func() {
		if err := c.Receive(); err != nil {
			logger.Errorf("error receiving from host: %s", err)
		} else {
			logger.Infof("connection closed")
		}
		cancel()
	}()

	<-ctx.Done()
}
