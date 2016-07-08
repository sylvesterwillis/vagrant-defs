package main

import (
	"log"
	"github.com/bitly/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	prod, _ := nsq.NewProducer("192.168.56.101:4150", config)

	err := prod.Publish("test_topic", []byte("Hello I have another message."))
	if err != nil {
		log.Panic("Could not connect")
	}

	prod.Stop()
}