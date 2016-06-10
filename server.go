package main

import (
    zmq "github.com/pebbe/zmq4"
)

func main() {
    socket, _ := zmq.NewSocket(zmq.PUB)
    socket.Bind("tcp://*:5555")
    socket.send("Ping.")
}
