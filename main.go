package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open .gz file
	file, err := os.Open("/var/tmp/bufio_benchmarking/xaa30.csv.gz")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create gzip.Reader from opened file
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("Error creating gzip reader: %v\n", err)
		return
	}
	defer gzipReader.Close()

	// Read uncompressed data using a chosen buffer
	buffer := make([]byte, 1024)
	totalBytesRead :=0
	for {
		n, err := gzipReader.Read(buffer) // Read into the buffer
		if n > 0 {
			//fmt.Print(string(buffer[:n]))
                        totalBytesRead += n
		}
		if err == io.EOF { // Check for end of file
			break
		}
		if err != nil { // Handle other errors
			fmt.Printf("Error reading from gzip reader: %v\n", err)
			return
		}
	}
	fmt.Printf("totalBytesRead: %d", totalBytesRead)
}
