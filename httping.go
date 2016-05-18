package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"strings"
)

func main() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"status":"ok"}`))
	}

	listener, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		signals := make(chan os.Signal)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		listener.Close()
	}()

	if err := http.Serve(listener, handler); err != nil && !strings.Contains(err.Error(), "closed network connection") {
		fmt.Println(err)
		os.Exit(1)
	}
}
