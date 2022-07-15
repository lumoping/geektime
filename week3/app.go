package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	app := http.Server{}

	exit := make(chan interface{})

	var helloHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "Hello GeekTime!")
	}

	var shutdownHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		exit <- 1
	}

	var longTimeHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)
		_, _ = fmt.Fprintln(w, "Long!")
	}

	g.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", helloHandlerFunc)
		mux.HandleFunc("/long", longTimeHandlerFunc)
		mux.HandleFunc("/shutdown", shutdownHandlerFunc)
		app.Handler = mux
		app.Addr = ":8080"
		return app.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-exit:
		case <-ctx.Done():
		}
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		return app.Shutdown(timeout)
	})

	g.Go(func() error {
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-s:
			return fmt.Errorf("os signal: %v", sig)
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	fmt.Printf("Exit : %v\n", g.Wait())
}
