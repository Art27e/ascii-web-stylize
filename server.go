package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Output struct {
	Orig string
	Res  string
}

var outputAscii string

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("./data/"))))
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/info", info)
	fmt.Printf("Starting server at port 8080\nOpen http://localhost:8080")
	removeLogFile := os.Remove("log.txt")
	if removeLogFile != nil {
		defer os.Remove("log.txt")
	}
	removeDlFile := os.Remove("ascii.txt")
	if removeDlFile != nil {
		defer os.Remove("ascii.txt")
	}
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}
