package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type JSONResponse struct {
	Response string `json:"response"`
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Empty port")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		resp, _ := json.Marshal(JSONResponse{"Hello, world!"})
		w.Write(resp)
	})

	logger.Info("Start on :" + port)

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go serv.ListenAndServe()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	serv.Shutdown(timeout)
}
