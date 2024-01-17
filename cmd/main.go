package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/czernous/mars-rover/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program> <path-to-input-file>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	rovers, err := ParseInput(file)

	if err != nil {
		log.Fatal("No rover data was found")
	}

	fmt.Printf("Found %v rovers\n", len(rovers))

	for i, r := range rovers {
		fmt.Printf("Processing rover #%v\n", i+1)

		for _, c := range r.Commands {
			if c == 'M' {
				err := r.Advance(&r.Plateau)
				if err != nil {
					fmt.Printf("Rover %d error: %v \n", i, err.Error())
				}
			} else {

				r.Turn(byte(c))
			}

		}
		fmt.Printf("Finished processing rover #%v\n", i+1)
		rp := r.GetPosition()
		fmt.Printf("Rover position is:  %v %v %s\n", rp.X, rp.Y, string(rp.Heading))
	}

	defer file.Close()
}
