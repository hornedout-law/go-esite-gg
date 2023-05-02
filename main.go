package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func init () {
    projectRoot, _ := os.Getwd()
    _, err := os.ReadDir(projectRoot+ "/blogs")
    if err!=nil {
        fmt.Print(err)
        os.Mkdir(projectRoot + "/blogs", fs.ModePerm)
    }
}

// this is a test handler

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

// create a directory named blogs

func GetContentsHandler(w http.ResponseWriter, r *http.Request) {
	// get blogs from the current dirctory

    projectRoot, _ := os.Getwd()

    resp := json.NewEncoder(w)
    files, err := os.ReadDir(projectRoot + "/blogs")

    for _, file:= range files {
        fmt.Print(file.Name())
    }

	if err != nil {
        fmt.Print(err)
        resp.Encode(err)
	}
	resp.Encode(files)
}

func main() {

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/content", GetContentsHandler)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))

}
