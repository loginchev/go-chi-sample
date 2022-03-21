package main

import (
	//"context"
	"fmt"
	"log"
	"net/http"
	"os"

	/*"os/signal"
	"syscall"
	"time"*/

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	host, _ := os.LookupEnv("HOST")
	port, _ := os.LookupEnv("PORT")
	fmt.Printf("%v:%v", host, port)
	server := &http.Server{Addr: fmt.Sprintf("%v:%v", host, port), Handler: service()}
	server.ListenAndServe()
	// Server run context
	/*	serverCtx, serverStopCtx := context.WithCancel(context.Background())

		// Listen for syscall signals for process to interrupt/quit
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-sig

			// Shutdown signal with grace period of 30 seconds
			shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

			go func() {
				<-shutdownCtx.Done()
				if shutdownCtx.Err() == context.DeadlineExceeded {
					log.Fatal("graceful shutdown timed out.. forcing exit.")
				}
			}()

			// Trigger graceful shutdown
			err := server.Shutdown(shutdownCtx)
			if err != nil {
				log.Fatal(err)
			}
			serverStopCtx()
		}()

		// Run the server
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

		// Wait for server context to be stopped
		<-serverCtx.Done()*/

}
