package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"time"
	"unicode/utf8"

	pb "Chitty-Chat_HW3_V2/chittychatpb"

	"google.golang.org/grpc"
)

func main() {
	const maxMessageLength = 128
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	go ListenMessages(client)
	// // Example: Join chat
	// joinResp, err := client.JoinChat(context.Background(), &pb.Participant{Name: name})
	// if err != nil {
	// 	log.Fatalf("Could not join: %v", err)
	// }

	// log.Println(joinResp.Message)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			break
		}
		if len(text) > maxMessageLength {
			log.Printf("Message too long, max %d characters\n", maxMessageLength)
			continue
		}
		// Validate UTF-8 characters only in message
		if !utf8.ValidString(text) {
			log.Println("Invalid UTF-8 characters")
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

	// // Example: Leave chat
	// leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: name})
	// if err != nil {
	// 	log.Fatalf("Could not leave: %v", err)
	// }
	// log.Println(leaveResp.Message)
}

func ListenMessages(client pb.ChittyChatClient) {
	for true {
		listen, err := client.RecieveMessages(context.Background(), &pb.Empty{})
		if err != nil {
			log.Println("Did not recieve:", err)
		}
		for true {
			var message pb.ChatMessage
			err := listen.RecvMsg(&message)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("recieve error:", err)
			}
			log.Println("New message:", message.Message, message.Timestamp)
		}
		time.Sleep(time.Second)
	}
}
