package main

import (
	"fmt"
	"net/http"
	"os"
)

type httpBody interface{
	*Response struct
}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//for Go, resp.body doesn't work. create byteslice with make
	bs := make([]byte, 99999)
}
