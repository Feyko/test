package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("This program takes one path to a log file as an argument.")
		os.Exit(1)
	}
	path := os.Args[1]
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read the file %v: %v", path, err)
	}
	re := regexp.MustCompile(`Warning: OSS: EOS_SessionModification_SetHostAddress\(EOS.\w{32}(\w{32}):7777\)`)
	id := re.FindSubmatch(b)
	if len(id) == 0 {
		fmt.Println("The log doesn't not contain a session ID")
		os.Exit(2)
	}
	fmt.Println(string(id[1]))
}
