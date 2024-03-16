package ws

import "log"

var messageChannel = make(chan *Message)

// Start listening for messages and saving them to the database
func SaveChatsToDatabase() {
	go func() {
		for receivedMsg := range messageChannel {
			// Handle received message
			// e.g., Save message to the database
			log.Println("Message Received in this go routine", receivedMsg)
			// log.Println("Sender", receivedMsg.SenderUsername)
			// log.Println("Reciver", receivedMsg.ReceiverUsername)
		}
	}()
}
