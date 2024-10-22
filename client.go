package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	pb "Chitty-Chat_HW3_V2/chittychatpb"

	"google.golang.org/grpc"
)

func main() {
	const maxMessageLength = 128
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	// Example: Join chat
	joinResp, err := client.JoinChat(context.Background(), &pb.Participant{Name: name})
	if err != nil {
		log.Fatalf("Could not join: %v", err)
	}
	log.Println(joinResp.Message)

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

	// Example: Leave chat
	leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: name})
	if err != nil {
		log.Fatalf("Could not leave: %v", err)
	}
	log.Println(leaveResp.Message)
}
