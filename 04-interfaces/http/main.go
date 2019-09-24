package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	res, err := http.Get("https://swapi.co/api/planets/1/?format=json")
	if err != nil {
		fmt.Println("an error happens", err)
		os.Exit(1)
	}

	// Manually read data
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	wr := logWriter{}
	io.Copy(wr, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	data := string(bs)
	fmt.Println(data)
	return len(bs), nil
}
