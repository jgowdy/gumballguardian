package gumballguardian

// Transport abstraction for mesh network
type Transport interface {
	Open()         // Client will connect, server will listen
	Close()        // Client will disconnect, server will stop listening and disconnect all clients
	OnConnect()    // When we connect or re-connect
	OnDisconnect() // Handle disconnection
	OnMessage()    // When we receive a new message from the other endpoint
	OnError()      // When we have an error

	SendMessage() // Send a new message to the other endpoint
	SendReply()   // Send a response to a message we sent
}

type Message struct {
	src     string
	dst     string
	ttl     int8
	payload []byte
}

func MessageFromString(src string, dst string, payload string) Message {
	return Message{src, dst, 64, []byte(payload)}
}
