package main

import (
	"fmt"
	"log"
	"time"

	"github.com/saranblock3/goma"
)

func main() {
	// Create the Homa socket with an id (equivalent of TCP port number)
	// This will fail if the id is already in use
	// The socket is connected to the Homa daemon
	homaSocket, err := goma.NewHomaSocket(8080)
	if err != nil {
		log.Fatal("Socket error:", err)
	}
	// Defer closing the socket
	defer homaSocket.Close()

	fmt.Println("SERVER STARTED")
	fmt.Println("")

	for {
		// Read any incoming messages
		content, sourceAddress, destinationAddress, sourceId, err := homaSocket.RecvFrom()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("===================================================")
		fmt.Printf("REQUEST FROM %s:%d TO %s:%d\n", sourceAddress, sourceId, destinationAddress, 8080)
		fmt.Printf("CONTENT: %s\n", content)
		fmt.Printf("TIME: %s\n", time.Now().Format(time.RFC3339))
		fmt.Println("===================================================")
		fmt.Println("")

		// Send a response to the client that just sent a message
		err = homaSocket.SendTo(
			[]byte("Greetings!"),
			destinationAddress, // Swap the source and destination addresses
			sourceAddress,
			sourceId,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
