package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Need a waitGrroup for all goroutines to be finished
	var wg sync.WaitGroup

	// A channel to collect folder names and their file counts
	results := make(chan [2]interface{})

	// Walk through the home directory and start a goroutine for each folder (1 level deep)
	filepath.WalkDir(homeDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error walking through directory:", err)
			return err
		}

		if d.IsDir() && path != homeDir {
			wg.Add(1)
			go func(folderPath string) {
				defer wg.Done()
				fileCount := countFiles(folderPath)
				results <- [2]interface{}{folderPath, fileCount}
			}(path)
			// Skip subdirectories to ensure 1-level depth
			return filepath.SkipDir
		}

		return nil
	})

	// Close the results channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print the results
	totalFiles := 0
	for result := range results {
		fmt.Printf("Folder: %s, Files: %d\n", result[0], result[1])
		totalFiles += result[1].(int)
	}
	fmt.Printf("Total files: %d\n", totalFiles)
}

func countFiles(folderPath string) int {
	fileCount := 0
	filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error walking through directory:", err)
			return err
		}
		if !d.IsDir() {
			fileCount++
		}
		return nil
	})
	return fileCount
}
