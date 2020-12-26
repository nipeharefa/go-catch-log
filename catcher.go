package main

type Catcher interface {
	Read([]byte, *Log) error
}
