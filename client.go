package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "Chitty-Chat_HW3_V2/chittychatpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const maxMessageLength = 128

	conn, err := grpc.NewClient("nicklas-laptop:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	// Join the chat
	joinResp, err := client.JoinChat(context.Background(), &pb.Participant{Name: name})
	if err != nil {
		log.Fatalf("Could not join: %v", err)
	}
	log.Println(joinResp.Message)

	// Start a goroutine to receive messages from the server
	go func() {
		stream, err := client.BroadcastMessages(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
		}

		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			log.Printf("[Broadcast] %s: %s (Timestamp: %d)", msg.Participant, msg.Message, msg.Timestamp)
		}
	}()

	// Handle sending messages
	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			break
		}
		if len(text) > maxMessageLength {
			log.Printf("Message too long, max %d characters\n", maxMessageLength)
			continue
		}

		_, err = client.PublishMessage(context.Background(), &pb.ChatMessage{
			Participant: name,
			Message:     text,
			Timestamp:   time.Now().Unix(),
		})
		if err != nil {
			log.Fatalf("Could not publish message: %v", err)
		}
	}

	// Leave the chat
	leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: name})
	if err != nil {
		log.Fatalf("Could not leave: %v", err)
	}
	log.Println(leaveResp.Message)
}
