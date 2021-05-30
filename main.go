package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webparser/parserapp/parserhandler"
)

func main() {

	l := log.New(os.Stdout, "parser:", log.LstdFlags)
	parseHandler := parserhandler.GetNewLogger(l)
	serverMux := mux.NewRouter()

	//Register the handlers to the server mux
	postRouter := serverMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/url", parseHandler.GetURLResp)

	prodServer := &http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	go func() {
		l.Println("Startinhg server at port 8080")
		if err := prodServer.ListenAndServe(); err != nil {
			l.Printf("Error starting server: %v", err)
			os.Exit(1)
		}

	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("Stopping server as per user interrupt", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	prodServer.Shutdown(tc)
}
