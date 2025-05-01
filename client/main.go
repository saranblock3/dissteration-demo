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
	homaSocket, err := goma.NewHomaSocket(7070)
	if err != nil {
		log.Fatal(err)
	}
	defer homaSocket.Close()

	fmt.Println("CLIENT STARTED")
	fmt.Println("")

	// Write to a remote application by specifying the content, address, and id
	err = homaSocket.SendTo(
		[]byte("Hello, World!"), // Content
		"10.10.1.1",             // Source address
		"10.10.1.2",             // Destination address
		8080,                    // Destination ID
	)
	if err != nil {
		log.Fatal(err)
	}

	// Read any incoming messages
	content, sourceAddress, destinationAddress, sourceId, err := homaSocket.RecvFrom()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===================================================")
	fmt.Printf("RESPONSE FROM %s:%d TO %s:%d\n", sourceAddress, sourceId, destinationAddress, 7070)
	fmt.Printf("CONTENT: %s\n", content)
	fmt.Printf("TIME: %s\n", time.Now().Format(time.RFC3339))
	fmt.Println("===================================================")
	fmt.Println("")
}
