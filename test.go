package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

func decompressGzip(input []byte) ([]byte, error) {
	// Create a bytes reader from the input gzip data
	byteReader := bytes.NewReader(input)

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer gzipReader.Close()

	// Read the decompressed data into a buffer
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, gzipReader)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %v", err)
	}

	return buffer.Bytes(), nil
}

func main() {
	// Here we assume you already have the gzip data in a variable called gzipData
	// For example:
	// var gzipData []byte = ...

	// Replace this comment with the actual gzip data
	var gzipData []byte = /* your gzip data */

	// Decompress the gzip data
	decompressedData, err := decompressGzip(gzipData)
	if err != nil {
		log.Fatalf("Error decompressing data: %v", err)
	}

	// Use the decompressed data as needed
	fmt.Printf("Decompressed data: %s\n", decompressedData)
}
