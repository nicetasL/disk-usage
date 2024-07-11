package main

import (
	"fmt"
	"os"
	"syscall"
)

// getDiskUsage calculates and displays disk usage statistics for a given path.
func getDiskUsage(path string) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		fmt.Printf("Error fetching disk usage for %s: %v\n", path, err)
		return
	}

	total := stat.Blocks * uint64(stat.Bsize) // Total disk size in bytes
	free := stat.Bfree * uint64(stat.Bsize)   // Free space in bytes
	used := total - free                      // Used space in bytes
	percentUsed := float64(used) / float64(total) * 100

	// Display disk usage statistics
	fmt.Printf("Disk usage of %s:\n", path)
	fmt.Printf("Total: %d GB\n", total/1e9)
	fmt.Printf("Used: %d GB (%.2f%%)\n", used/1e9, percentUsed)
	fmt.Printf("Free: %d GB\n", free/1e9)
}

func main() {
	// Default path is root
	path := "/"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// Check if the specified path exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Error: Path '%s' does not exist.\n", path)
		return
	} else if err != nil {
		fmt.Printf("Error accessing path '%s': %v\n", path, err)
		return
	}

	// Calculate and display disk usage
	getDiskUsage(path)
}
