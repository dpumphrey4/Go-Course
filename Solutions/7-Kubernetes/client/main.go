package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "client/gophertalk"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Send a normal chirp
func sendChirp(client pb.GopherChatClient, sender string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := pb.Chirp{
		Sender:    sender,
		Message:   "Chirp at " + time.Now().Format("15:04:05"),
		Timestamp: time.Now().Unix(),
	}

	resp, err := client.SendChirp(ctx, &message)
	if err != nil {
		log.Printf("[%s] Chirp error: %v", sender, err)
		return
	}

	log.Printf("[%s] Chirp response: %s", sender, resp.Message)
}

// Send a playful gnaw
func sendGnaw(client pb.GopherChatClient, sender string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nibble := pb.Nibble{
		Sender:  sender,
		Content: "Gnaw at " + time.Now().Format("15:04:05"),
	}

	resp, err := client.SendNibble(ctx, &nibble)
	if err != nil {
		log.Printf("[%s] Gnaw error: %v", sender, err)
		return
	}

	log.Printf("[%s] Gnaw ack: received=%v by %s", sender, resp.Received, resp.Responder)
}

func main() {
	// Connect to the server
	conn, err := grpc.NewClient(
		"my-server-service:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGopherChatClient(conn)

	gopherNames := []string{"Gophy", "Gus", "Ginny", "Griff"}

	batchSize := 10000
	var wg sync.WaitGroup
	messageCount := 0

	log.Println("Starting infinite gopher chatter with gnaws...")

	for {
		for _, name := range gopherNames {
			// Randomly decide to send a chirp or a gnaw
			wg.Add(1)
			messageCount++
			if rand.Intn(2) == 0 {
				go sendChirp(client, name, &wg)
			} else {
				go sendGnaw(client, name, &wg)
			}

			// Batch pause
			if messageCount%batchSize == 0 {
				wg.Wait()
				log.Printf("Sent %d messages, pausing for 500ms...", messageCount)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}
