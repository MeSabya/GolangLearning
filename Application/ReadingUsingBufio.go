package main

//This is a program to read a string word by word

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "sabya sachi nayak"
	//Read the above string word by word
	scanner := bufio.NewScanner(strings.NewReader(s))
	//Comment the line below and try it again, It will print the whole string as it is
	scanner.Split(bufio.ScanWords) //Also available bufio.ScanBytes bufio.ScanLines bufio.ScanRunes

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
