package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var logs Logs

// GCPLogFormat ...
type GCPLogFormat struct {
	Message  string            `json:"message"`
	Severity string            `json:"severity"`
	Labels   map[string]string `json:"logging.googleapis.com/labels"`
}

// Log ...
type Log struct {
	Message   string       `json:"message"`
	Timestamp time.Time    `json:"timestamp"`
	GCPLogger GCPLogFormat `json:"google"`
}

type Logs []Log

// Ear ...
type Ear struct {
	ID string `json:"bucketname"`
}

func startHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	gcp := GCPLogger{}

	reader := bufio.NewReader(os.Stdin)
	logs = make(Logs, 0)
	go startHttp()

	for {
		text, _, _ := reader.ReadLine()
		if text != nil {
			l := Log{
				Message:   string(text),
				Timestamp: time.Now(),
			}

			_ = gcp.Read(text, &l)

			logs = append(logs, l)

			fmt.Println(string(text))
		}
	}
}
