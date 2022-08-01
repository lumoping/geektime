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

	g.Go(startServer(&app, ":8080"))

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

func startServer(app *http.Server, addr string, routers ...Router) func() error {
	return func() error {
		mux := http.NewServeMux()
		for _, router := range routers {
			mux.HandleFunc(router.pattern, router.HandlerFunc)
		}
		app.Handler = mux
		app.Addr = addr
		return app.ListenAndServe()
	}
}

type Router struct {
	pattern string
	http.HandlerFunc
}
