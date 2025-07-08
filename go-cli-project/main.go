package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcDir := flag.String("src", ".", "The source directory to organize")
	destDir := flag.String("dest", "./organized", "The destination directory for organized files")
	dryRun := flag.Bool("dry-run", false, "Perform a dry run without moving files")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	recursive := flag.Bool("recursive", false, "Recursively organize files in subdirectories")
	flag.Parse()

	if *srcDir == "" || *destDir == "" {
		fmt.Println("Source and destination directories must be specified.")
		return
	}

	if *recursive {
		if *verbose {
			fmt.Printf("Starting recursive scan in %s...\n", *srcDir)
		}
		err := filepath.WalkDir(*srcDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			processErr := processFile(path, *destDir, *dryRun, *verbose)
			if processErr != nil {
				fmt.Printf("Error processing %s: %v\n", path, processErr)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
		}
	} else {
		if *verbose {
			fmt.Printf("Starting non-recursive scan in %s...\n", *srcDir)
		}
		files, err := os.ReadDir(*srcDir)
		if err != nil {
			fmt.Printf("Error reading source directory: %v\n", err)
			return
		}
		for _, file := range files {
			if file.IsDir() {
				if *verbose {
					fmt.Printf("Skipping directory: %s\n", file.Name())
				}
				continue
			}
			sourcePath := filepath.Join(*srcDir, file.Name())
			err := processFile(sourcePath, *destDir, *dryRun, *verbose)
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", file.Name(), err)
			}
		}
	}
	fmt.Println("File organization complete.")
}

func processFile(sourcePath string, destRoot string, dryRun bool, verbose bool) error {
	ext := filepath.Ext(sourcePath)
	if ext == "" {
		return fmt.Errorf("file %q has no extension", sourcePath)
	}

	subdir := strings.TrimPrefix(ext, ".")
	finalDestPath := filepath.Join(destRoot, subdir, filepath.Base(sourcePath))

	if verbose {
		fmt.Printf("Planning to move %q to %q\n", sourcePath, finalDestPath)
	}

	if dryRun {
		return nil
	}

	err := os.MkdirAll(filepath.Dir(finalDestPath), 0755)
	if err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	err = os.Rename(sourcePath, finalDestPath)
	if err != nil {
		return fmt.Errorf("could not move file: %w", err)
	}

	fmt.Printf("Successfully moved %s\n", filepath.Base(sourcePath))
	return nil
}
