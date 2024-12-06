package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	err := ensureSingleInstance()
	if err != nil {
		log.Log("error ensuring single instance: %s", err.Error())

		os.Exit(1)
	}

	go read()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			respondJSON(w, map[string]interface{}{
				"error": "invalid url",
			})

			return
		}

		response, err := request(url)
		if err != nil {
			respondJSON(w, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		response.Forward(w)
	})

	server := &http.Server{
		Addr:    ":4269",
		Handler: mux,
	}

	listener, err := net.Listen("tcp", ":4269")
	if err != nil {
		log.Log("error listening: %s", err.Error())

		os.Exit(1)
	}

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Log("error serving: %s", err.Error())

			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	server.Close()
	listener.Close()

	os.Exit(0)
}

func respondJSON(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsn, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write(jsn)
}
