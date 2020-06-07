package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func allWordGen(ch chan string, file *os.File) {
	buf := make([]byte, 2048)

	n, err := file.Read(buf)
	for {
		if n == 0 || err == io.EOF {
			return
		}
		temp := strings.Split(string(buf), " ")
		for _, w := range temp {
			ch <- w
		}
		n, err = file.Read(buf)
	}
}

func allwords(filename string) <-chan string {
	readStream := make(chan string)
	go func() {
		file, _ := os.Open(filename)
		defer close(readStream)
		defer file.Close()
		allWordGen(readStream, file)
	}()
	return readStream
}

func main() {
	for w := range allwords("testfile.txt") {
		fmt.Println(w)
	}
}
