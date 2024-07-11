package main

import (
	"fmt"
	"os"
	"syscall"
)

func getDiskUsage(path string) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		fmt.Printf("Error fetching disk usage for %s: %v\n", path, err)
		return
	}
	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	fmt.Printf("Disk usage of %s:\n", path)
	fmt.Printf("Total: %d GB\n", total/1e9)
	fmt.Printf("Used: %d GB\n", used/1e9)
	fmt.Printf("Free: %d GB\n", free/1e9)
}

func main() {
	path := "/"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	getDiskUsage(path)
}
