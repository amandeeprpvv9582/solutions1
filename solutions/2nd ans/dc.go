package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type DirectoryStats struct {
	Files          int
	Subdirectories int
}

func countFilesAndSubdirectories(path string) (int, int) {
	files := 0
	subdirectories := 0

	filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", filePath, err)
			return nil
		}

		if info.IsDir() {
			subdirectories++
		} else {
			files++
		}

		return nil
	})

	return files, subdirectories
}

func processDirectory(path string, wg *sync.WaitGroup, directoryStats chan<- DirectoryStats) {
	defer wg.Done()

	files, subdirectories := countFilesAndSubdirectories(path)
	directoryStats <- DirectoryStats{Files: files, Subdirectories: subdirectories}
}

func main() {
	rootDir := `C:\Program Files\dotnet` 

	var wg sync.WaitGroup
	directoryStats := make(chan DirectoryStats)

	wg.Add(1)
	go processDirectory(rootDir, &wg, directoryStats)

	totalFiles := 0
	totalSubdirectories := 0

	go func() {
		wg.Wait()
		close(directoryStats)
	}()

	for stats := range directoryStats {
		totalFiles += stats.Files
		totalSubdirectories += stats.Subdirectories
	}

	fmt.Printf("Total Files: %d\n", totalFiles)
	fmt.Printf("Total Subdirectories: %d\n", totalSubdirectories)
}
