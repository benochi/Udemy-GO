package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type FileReader interface {
	Read(filename string) ([]byte, error)
}

type SimpleFileReader struct{}

func (sfr SimpleFileReader) Read(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a filename as an argument.")
	}

	filename := os.Args[1]
	fileReader := SimpleFileReader{}

	content, err := openFile(filename, fileReader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(content))
}

func openFile(filename string, reader FileReader) ([]byte, error) {
	return reader.Read(filename)
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// // actual video solution.
// func main() {
// 	f, err = os.Open(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	io.Copy(os.Stdout, f)
// }
