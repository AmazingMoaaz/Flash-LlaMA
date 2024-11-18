package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"source-code-review/config" // Import config package for ASCII art and colored logging
	"source-code-review/internal/markdown"
	"source-code-review/internal/scanner"
	"syscall"

	"github.com/jpoz/groq" // Import Groq Go SDK
)

func main() {
	// Show ASCII art at startup
	config.ShowASCII()

	// Handle graceful shutdown signals (SIGINT, SIGTERM)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Command line arguments for file or directory and whether to use Groq
	filePtr := flag.String("file", "", "File to scan for vulnerabilities")
	dirPtr := flag.String("dir", "", "Directory to scan for vulnerabilities")
	saveDir := flag.String("save", ".", "Directory to save the results (default is current directory)") // Custom save directory
	flag.Parse()

	// Check if the save directory exists
	if _, err := os.Stat(*saveDir); os.IsNotExist(err) {
		config.Warn(fmt.Sprintf("Save directory does not exist: %s", *saveDir))
		os.Exit(1)
	}

	// Signal handling routine in a separate goroutine
	go func() {
		// Wait for the signal
		sig := <-sigs
		// When a signal is received, attempt graceful shutdown
		config.Info(fmt.Sprintf("Received signal: %s. Attempting graceful shutdown.", sig))
		os.Exit(0)
	}()

	var filesToScan []string
	if *filePtr != "" {
		filesToScan = append(filesToScan, *filePtr)
	} else if *dirPtr != "" {
		files, err := scanner.ScanDirectory(*dirPtr)
		if err != nil {
			config.Warn(fmt.Sprintf("Failed to scan directory: %v", err))
			os.Exit(1)
		}
		filesToScan = files
	} else {
		config.Warn("No file or directory specified")
		os.Exit(1)
	}

	// Initialize the Groq client with the API key from the config
	client := groq.NewClient(groq.WithAPIKey("YOUR_API_KEY")) // Replace YOUR_API_KEY with your Groq API key

	// Iterate over the files and process each one
	for _, file := range filesToScan {
		config.Info(fmt.Sprintf("Scanning %s with Groq AI", file))

		// Read the file content
		content, err := scanner.ScanFile(file)
		if err != nil {
			config.Warn(fmt.Sprintf("Failed to read file: %v", err))
			os.Exit(1)
		}

		// Prepare the message for Groq chat completion
		messages := []groq.Message{
			{Role: "user", Content: content},
		}

		// Request AI response from Groq
		response, err := client.CreateChatCompletion(groq.CompletionCreateParams{
			Model:    "llama-3.1-70b-versatile", // Use the appropriate model for your use case
			Messages: messages,
		})

		if err != nil {
			log.Fatalf("Error creating chat completion: %v", err)
		}

		// Save the result to a markdown file in the specified save directory
		config.Info(fmt.Sprintf("Saving results for %s", file))
		baseName := filepath.Base(file)                       // Get the base file name (without directories)
		resultFile := filepath.Join(*saveDir, baseName+".md") // Save the result in the custom directory
		err = markdown.SaveMarkdown(resultFile, response.Choices[0].Message.Content)
		if err != nil {
			config.Warn(fmt.Sprintf("Failed to save markdown for %s: %v", file, err))
			os.Exit(1)
		}
	}

	config.Info("Scan complete.")
}
