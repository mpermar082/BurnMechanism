// cmd/burnmechanism/main.go
package main

import (
    "flag"
    "log"
    "os"

    "burnmechanism/internal/burnmechanism"
)

// Main function is the entry point of the application
func main() {
    // Define command-line flags
    verbose := flag.Bool("verbose", false, "Enable verbose logging")
    input := flag.String("input", "", "Input file path")
    output := flag.String("output", "", "Output file path")
    flag.Parse()

    // Create a new application instance
    app := burnmechanism.NewApp(*verbose)

    // Run the application with the provided input and output paths
    if err := app.Run(*input, *output); err != nil {
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}