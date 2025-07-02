package main

import (
	"context"    // For managing request-scoped values, cancellation, and deadlines
	"fmt"        // For formatted I/O
	"log"        // For logging messages
	"net/http"   // For HTTP server functionality
	"os"         // For interacting with the operating system
	"os/signal"  // For handling OS signals
	"syscall"    // For specific system calls (like SIGINT, SIGTERM)
	"time"       // For time-related operations (e.g., timeouts)
)

// Database struct simulates a database connection.
// In a real application, this would hold an actual database client (e.g., *sql.DB).
type Database struct {
	connected bool // Simple flag to indicate connection status
}

// Connect simulates establishing a database connection.
func (db *Database) Connect() {
	log.Println("Database: Connecting...")
	time.Sleep(500 * time.Millisecond) // Simulate connection time
	db.connected = true
	log.Println("Database: Connected.")
}

// Close simulates closing the database connection.
func (db *Database) Close() {
	if db.connected {
		log.Println("Database: Closing connection...")
		time.Sleep(500 * time.Millisecond) // Simulate closing time
		db.connected = false
		log.Println("Database: Connection closed.")
	} else {
		log.Println("Database: Not connected, nothing to close.")
	}
}

func main() {
	// 1. Initialize and connect to the simulated database.
	db := &Database{}
	db.Connect()

	// 2. Set up the HTTP server and its handler.
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Server: Received request from %s for %s", r.RemoteAddr, r.URL.Path)
		// Simulate some work being done by the handler.
		// This simulates an "in-flight" request that needs time to complete.
		time.Sleep(2 * time.Second)
		fmt.Fprintf(w, "Hello from Go Server! Request processed.\n")
	})

	server := &http.Server{
		Addr:    ":8080", // Server will listen on port 8080
		Handler: mux,     // Assign the request multiplexer
	}

	// 3. Start the HTTP server in a separate goroutine.
	// This allows the main goroutine to continue executing and listen for signals.
	go func() {
		log.Println("Server: Starting on port 8080...")
		// ListenAndServe blocks until the server stops or an error occurs.
		// http.ErrServerClosed is expected during a graceful shutdown.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server: ListenAndServe error: %v", err)
		}
		log.Println("Server: Goroutine stopped.")
	}()

	// 4. Set up a channel to listen for OS termination signals.
	quit := make(chan os.Signal, 1) // Buffered channel of size 1
	// signal.Notify registers the given channel to receive notifications of the specified signals.
	// syscall.SIGINT is typically sent by Ctrl+C.
	// syscall.SIGTERM is a standard termination signal sent by `kill` command or container orchestrators.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Main: Waiting for termination signal...")
	// 5. Block the main goroutine until a termination signal is received.
	sig := <-quit

	log.Printf("Main: Received signal: %v. Initiating graceful shutdown...", sig)

	// 6. Create a context with a timeout for the server shutdown.
	// This ensures that the server doesn't hang indefinitely if requests take too long to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure the context's resources are released when main exits.

	// 7. Attempt to gracefully shut down the HTTP server.
	// server.Shutdown stops the server from accepting new connections and waits for
	// active connections to finish within the context's deadline.
	if err := server.Shutdown(ctx); err != nil {
		// If an error occurs (e.g., timeout), log a forced shutdown message.
		log.Printf("Server: Forced shutdown due to error: %v", err)
	} else {
		// If no error, the server shut down gracefully.
		log.Println("Server: HTTP server shut down gracefully.")
	}

	// 8. Close and clean up database resources.
	// This is done after the HTTP server has stopped accepting requests and
	// in-flight requests are handled, ensuring data integrity.
	db.Close()

	log.Println("Main: Application exited gracefully.")
}