package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"os"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)

		resp, _ := json.Marshal(JSONResponse{"Hello, world!"})
		w.Write(resp)
	})
	http.ListenAndServe(":"+port, nil)
}
