package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adamsafr/golang-basic/config"
	"github.com/adamsafr/golang-basic/pkg/http/middleware"
	"github.com/adamsafr/golang-basic/pkg/http/router"
	logService "github.com/adamsafr/golang-basic/pkg/service/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found.")
	}

	logService.LoadConfig()
}

var listenAddr string

func main() {
	config.Load()

	logger := logService.Inst()
	w := logger.Writer()
	defer w.Close()

	logger.Info("Application has been started.")

	flag.StringVar(&listenAddr, "listen-addr", ":8081", "server listen address")
	flag.Parse()

	done, quit := make(chan bool, 1), make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := newWebServer(log.New(w, "", 0))
	go gracefulShutdown(server, logger, quit, done)

	logger.Println("Server is ready to handle requests at", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func gracefulShutdown(server *http.Server, logger *logrus.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	close(done)
}

func newWebServer(logger *log.Logger) *http.Server {
	serveMux := http.NewServeMux()

	setUpRoutes(serveMux)

	return &http.Server{
		Addr:         listenAddr,
		Handler:      serveMux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

func setUpRoutes(serveMux *http.ServeMux) {
	for _, rt := range router.GetRouteList() {
		path, method, handler := rt.Path, rt.Method, rt.Handler

		serveMux.Handle(
			path,
			middleware.TimeoutMiddleware(
				middleware.RecoverWrapMiddleware(
					middleware.RequestMethodMiddleware(method, http.HandlerFunc(handler)),
				),
				config.Get().Response.Timeout,
			),
		)
	}
}
