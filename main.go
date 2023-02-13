package main

import (
	"bufio"
	"fmt"
	"net/netip"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read each line of the file
	for scanner.Scan() {

		ip, err := netip.ParseAddr(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(ip.StringExpanded())
	}
}
