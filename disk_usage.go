package main

import (
	"fmt"
	"syscall"
)

func main() {
	var stat syscall.Statfs_t
	path := "/"
	syscall.Statfs(path, &stat)
	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	fmt.Printf("Disk usage of %s:\n", path)
	fmt.Printf("Total: %d GB\n", total/1e9)
	fmt.Printf("Used: %d GB\n", used/1e9)
	fmt.Printf("Free: %d GB\n", free/1e9)
}
