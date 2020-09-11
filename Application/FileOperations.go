//Reading and writting to a file
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	WriteToFile("test.txt", "The Witcher")

	ReadFromFile("test.txt")

}

func ReadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error in reading file")
		return
	}
	defer file.Close()
	buf := make([]byte, 50)
	if _, err := io.ReadFull(file, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}

	io.WriteString(os.Stdout, string(buf))
	fmt.Println()

}

func WriteToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}

	defer file.Close()
	_, err = io.WriteString(file, data)
	if err != nil {
		return
	}

	file.Sync()

}
