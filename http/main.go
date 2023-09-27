package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//for Go, resp.body doesn't work. create byteslice with make
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	//os.Stdout is a rwriter function repalces by lw
	io.Copy(lw, resp.Body)
}

// implementing the writer interface, becuase writer has a function called Write.
func (logWriter) Write(bs []byte) (int, error) {
	//interfaces don't enforce correct code so you can mess this up.
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	//return int and err, but make sure valuies are correct per documentation.
	return len(bs), nil
}
