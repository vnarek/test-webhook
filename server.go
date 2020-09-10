package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	argsPorts := os.Args[1:]

	ports := make([]int, 0)

	for _, argsPort := range argsPorts {
		port, err := strconv.Atoi(argsPort)
		if err != nil {
			log.Fatal(err)
		}
		ports = append(ports, port)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	wg := sync.WaitGroup{}
	for _, port := range ports {
		wg.Add(1)
		go func(port int) {
			log.Printf("running on port: %d", port)
			if err := httpServerFactory(port).ListenAndServe(); err != nil {
				wg.Done()
				log.Fatal(err)
			}
		}(port)
	}
	wg.Wait()
}

func httpServerFactory(port int) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if _, err := writer.Write([]byte(fmt.Sprintf("port: %d", port))); err != nil {
				log.Fatal(err)
			}
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
