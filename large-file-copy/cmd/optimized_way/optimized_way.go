package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/zydhanlinnar11/go-contents/large-file-copy/pkg/tools"
)

func main() {
	// Start HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--- Before operations ---")
		tools.PrintMemStats("Initial")

		// Benchmark download time
		downloadStart := time.Now()

		fmt.Println("\n--- Streaming file ---")
		body, err := tools.DownloadBigFile()
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()

		if _, err := io.Copy(w, body); err != nil {
			log.Fatal(err)
		}

		duration := time.Since(downloadStart)
		fmt.Printf("Streaming time: %v\n", duration)

		fmt.Println("\n--- After operations ---")
		tools.PrintMemStats("After")
	})

	// Simple way to run the server, not recommended for production
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
