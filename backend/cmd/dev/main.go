package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	// Start reader service
	readerCmd := exec.Command("go", "run", "./cmd/reader/main.go", "-config", *configPath)
	readerCmd.Stdout = os.Stdout
	readerCmd.Stderr = os.Stderr
	if err := readerCmd.Start(); err != nil {
		log.Fatalf("Failed to start reader: %v", err)
	}

	// Start writer service
	writerCmd := exec.Command("go", "run", "./cmd/writer/main.go", "-config", *configPath)
	writerCmd.Stdout = os.Stdout
	writerCmd.Stderr = os.Stderr
	if err := writerCmd.Start(); err != nil {
		log.Fatalf("Failed to start writer: %v", err)
	}

	// Handle shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop both services
	fmt.Println("\nShutting down services...")
	
	if err := readerCmd.Process.Signal(syscall.SIGTERM); err != nil {
		log.Printf("Error stopping reader: %v", err)
	}
	if err := writerCmd.Process.Signal(syscall.SIGTERM); err != nil {
		log.Printf("Error stopping writer: %v", err)
	}

	readerCmd.Wait()
	writerCmd.Wait()
} 