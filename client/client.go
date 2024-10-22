package main

import (
	"context"
	"log"
	"time"

	pb "chittychatpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)

	// Example: Join chat
	joinResp, err := client.JoinChat(context.Background(), &pb.Participant{Name: "Alice"})
	if err != nil {
		log.Fatalf("Could not join: %v", err)
	}
	log.Println(joinResp.Message)

	// Example: Publish a message
	_, err = client.PublishMessage(context.Background(), &pb.ChatMessage{
		Participant: "Alice",
		Message:     "Hello, everyone!",
		Timestamp:   time.Now().Unix(),
	})
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Example: Leave chat
	leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: "Alice"})
	if err != nil {
		log.Fatalf("Could not leave: %v", err)
	}
	log.Println(leaveResp.Message)
}
