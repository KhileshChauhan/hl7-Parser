package hl7

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// LoadMessage takes the message passed and returns the list of segments in []byte
func LoadMessage(message []byte) [][]byte {
	message = bytes.TrimSpace(message)
	var segments [][]byte
	scanner := bufio.NewScanner(bytes.NewReader(message))
	for scanner.Scan() {
		line := scanner.Bytes()
		line = bytes.TrimSpace(line)
		segments = append(segments, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return segments
}
