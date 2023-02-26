package main

import (
	"html/template"
	"net/http"
	"os"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("data/index.html")
	check(err)
	switch request.Method {
	case "GET":
		if request.URL.Path != "/" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	case "POST":
		if request.URL.Path != "/ascii-art" {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	}
	inputData := request.FormValue("userinput")
	for i := 0; i < len(inputData); i++ {
		if inputData[i] > 128 {
			writer.WriteHeader(500)
			return
		}
	}
	chooseFont := request.FormValue("fonts")
	fMap, err := createMap(chooseFont + ".txt")
	if request.URL.Path == "/ascii-art" && err != nil {
		writer.WriteHeader(500)
		return
	}
	switch {
	case chooseFont == "standard":
		outputAscii = printAsciiArt(inputData, fMap)
	case chooseFont == "shadow":
		outputAscii = printAsciiArt(inputData, fMap)
	case chooseFont == "thinkertoy":
		outputAscii = printAsciiArt(inputData, fMap)
	}
	writeRes := Output{Orig: inputData, Res: outputAscii}
	if inputData != "" {
		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("log.txt", options, os.FileMode(0o666))
		check(err)
		file.WriteString(outputAscii)
		err = file.Close()
		check(err)
	}
	writer.WriteHeader(200)
	err = html.Execute(writer, writeRes)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
}

func downloadFile(writer http.ResponseWriter, request *http.Request) {
	f, err := os.Create("ascii.txt")
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	defer f.Close()
	f.WriteString(outputAscii)
	writer.Header().Set("Content-Disposition", "attachment; filename=ascii.txt")
	writer.Header().Set("Content-Type", request.Header.Get("Content-Type"))
	http.ServeFile(writer, request, "ascii.txt")
}

func info(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("data/info.html")
	check(err)
	err = html.Execute(writer, nil)
	if err != nil {
		writer.WriteHeader(500)
	}
}