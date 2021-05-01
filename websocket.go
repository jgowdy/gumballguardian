package main

// Works for websocket client and server
type Transport interface {
	Open()        // Client will connect, server will listen
	Close()       // Client will disconnect, server will stop listening and disconnect all clients
	OnConnect()   // When we connect or re-connect
	SendMessage() // Send a new message to the other endpoint
	OnMessage()   // When we receive a new message from the other endpoint
	SendReply()   // Send a response to a message we sent
	OnError()     // When we have an error
}

type WebsocketClient struct {
	port int16
}

type WebsocketServer struct {
	port int16
}

func (ws WebsocketServer) Open() {

}

func (ws WebsocketServer) Close() {

}

func (ws WebsocketServer) OnConnect() {

}

func (ws WebsocketServer) SendMessage() {

}

func (ws WebsocketServer) OnMessage() {

}

func (ws WebsocketServer) SendReply() {

}

func (ws WebsocketServer) OnError() {

}

func (wc WebsocketClient) Open() {

}

func (wc WebsocketClient) Close() {

}

func (wc WebsocketClient) OnConnect() {

}

func (wc WebsocketClient) SendMessage() {

}

func (wc WebsocketClient) OnMessage() {

}

func (wc WebsocketClient) SendReply() {

}

func (wc WebsocketClient) OnError() {

}
