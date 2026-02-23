package main

import (
	"bufio"
	"fmt"
	"os"
)


func count(file *os.File, split bufio.SplitFunc) (int) {
	scanner := bufio.NewScanner(file)
	var count int = 0

	scanner.Split(split)

	for scanner.Scan() {
		count++
	}

	file.Seek(0, 0)
	return count
}


func wc() {

	command := os.Args[1]
	
	if os.Args[0] != "wc" {
		fmt.Println("Invalid input")
		return
	}
	
	switch command {
	case "help":
		fmt.Print("Usage: wc [command]\n\n")
		fmt.Println("Commands:")
		fmt.Println("  help [command]         Display this help message")
		fmt.Println("  count <file> [options] Count the number of lines in the input")
		fmt.Println("\nFor more detailed documentation, use 'wc help [command]'")
	
	case "count":
		fmt.Println("Counting lines in file")
		file, err := os.Open(os.Args[2])
	
		if err != nil {
			fmt.Println("Error opening file")
			return
		}
	
		words := count(file, bufio.ScanWords)
		lines := count(file, bufio.ScanLines)
		bytes := count(file, bufio.ScanBytes)
		chars := count(file, bufio.ScanRunes)

		fmt.Printf("Words: %d\n", words)
		fmt.Printf("Lines: %d\n", lines)
		fmt.Printf("Bytes: %d\n", bytes)
		fmt.Printf("Chars: %d\n", chars)
	
		defer file.Close()
	
	default:
		fmt.Println("Invalid command")
	
	}
}


func main() {
	os.Args = []string{"wc", "count", "p1_wc_clone.go"}
	wc()
}

