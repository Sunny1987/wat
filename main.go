package main

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webparser/parserapp/parserhandler"
)

func main() {
	//heroku related updates for port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	l := log.New(os.Stdout, "WAT:", log.LstdFlags)
	parseHandler := parserhandler.GetNewLogger(l)
	serverMux := mux.NewRouter()

	//Register the handlers to the server mux
	postRouter := serverMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/scan", parseHandler.GetURLResp)
	postRouter.HandleFunc("/login", parseHandler.Login)
	postRouter.Use(parseHandler.MiddlewareValidation)

	getRouter := serverMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/getscans", parseHandler.GetUserBasedScans)
	getRouter.Use(parseHandler.MiddlewareValidation)

	//CORS
	ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))

	//heroku related updates
	port = ":" + port

	//local test
	//port := ":8080"

	prodServer := &http.Server{
		Addr:         port,
		Handler:      ch(serverMux),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	go func() {
		myFigure := figure.NewFigure("WAT", "", true)
		myFigure.Print()
		l.Println("version: 1.0.0")
		l.Println("Author: Sabyasachi Roy")
		l.Println("Starting server...")
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
