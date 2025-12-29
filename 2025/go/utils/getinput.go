package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// ReadInput fetches the Advent of Code input for the given day and year.
// It caches the input in a temp file to avoid repeated requests.
// Requires AOC_SESSION environment variable to be set with your session cookie.
func ReadInput(day int) []string {
	return ReadInputForYear(day, 2025)
}

// ReadInputForYear fetches input for a specific day and year.
func ReadInputForYear(day, year int) []string {
	tempDir := os.TempDir()
	cacheFile := filepath.Join(tempDir, fmt.Sprintf("aoc_%d_day%02d.txt", year, day))

	// Check if cached file exists
	if _, err := os.Stat(cacheFile); err == nil {
		return readLinesFromFile(cacheFile)
	}

	// Fetch from Advent of Code
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		panic("AOC_SESSION environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create request: %v", err))
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	req.Header.Set("User-Agent", "github.com/manning0218/adventOfCode by manning0218@gmail.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("failed to fetch input: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("failed to fetch input: status %d", resp.StatusCode))
	}

	// Create temp file and write response
	file, err := os.Create(cacheFile)
	if err != nil {
		panic(fmt.Sprintf("failed to create cache file: %v", err))
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(fmt.Sprintf("failed to write cache file: %v", err))
	}

	return readLinesFromFile(cacheFile)
}

// readLinesFromFile reads all lines from a file and returns them as a slice.
func readLinesFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open file: %v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("failed to read file: %v", err))
	}

	return lines
}
