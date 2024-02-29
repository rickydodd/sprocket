package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string

	flag.StringVar(&port, "port", "8080", "the port over which to serve")
	flag.Parse()

	// Specify the engine and set certain qualities of the server
	engine := gin.Default()
	srv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      engine,
	}

	// Routing
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Listen and serve on a goroutine
	go func() {
		err := srv.ListenAndServe()
		// If the error is the server closed, then ignore.
		// This case is handled later in the program, during the shutdown
		// sequence.
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("server (error):", err)
		}
	}()

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal // Wait for program to receive interrupt or terminate signal.
	log.Println("server: shutdown sequence initiated")

	// Give the server five seconds before issuing the shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	<-ctx.Done()

	// Issue graceful shutdown.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("server (error):", err)
	}

	log.Println("server: shutdown sequence successful.")

	log.Println("application: closing...")
	os.Exit(0)
}
