package main

import (
    "log"
    "sync"
    "time"
	"github.com/bitly/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
  	wg.Add(1)

	lastMessageTime := time.Now()
    config := nsq.NewConfig()
    server, _ := nsq.NewConsumer("test_topic", "chan", config)

    //Process incoming messages and update the lastMessageTime accordingly.
    server.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
    	log.Printf("%s", m.Body)
    	log.Printf(" Time since last message: %d", time.Now().Unix() - lastMessageTime.Unix())
    	lastMessageTime = time.Now()
    	return nil
	}), 1)

	err := server.ConnectToNSQD("192.168.56.101:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	// Infinitely loop while messages are being received
	// and stop the consumer when 5 seconds has passed since the last message.
	for time.Now().Unix() - lastMessageTime.Unix() < 5 {}

	server.Stop()
}
