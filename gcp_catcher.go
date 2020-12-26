package main

import "encoding/json"

type GCPLogger struct {
}

func (gcp *GCPLogger) Read(d []byte, l *Log) error {

	var a GCPLogFormat

	err := json.Unmarshal(d, &a)
	if err != nil {
		return err
	}

	l.GCPLogger = a
	l.Message = a.Message
	return nil
}
