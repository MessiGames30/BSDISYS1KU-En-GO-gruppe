package main

import (
	pb "Chitty-Chat_HW3_V2/chittychatpb" // Import the generated protobuf package
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChittyChatServer
	participants map[string]int64
	messages     []*pb.ChatMessage
	clients      []pb.ChittyChat_RecieveMessagesServer
	mu           sync.Mutex
	logicalClock int64
}

// Start the server
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChittyChatServer(grpcServer, &server{
		// participants: make(map[string]int64),
		messages: []*pb.ChatMessage{},
	})

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// PublishMessage - allows clients to send a message
func (s *server) PublishMessage(ctx context.Context, msg *pb.ChatMessage) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Increment the logical clock for each new message
	s.logicalClock++

	// Broadcast message to all participants
	broadcast := &pb.ChatMessage{
		Participant: msg.Participant,
		Message:     msg.Message,
		Timestamp:   s.logicalClock,
	}
	s.messages = append(s.messages, broadcast)
	for _, stream := range s.clients {
		stream.Send(broadcast)
	}

	log.Printf("Message from %s: %s (Timestamp: %d)", msg.Participant, msg.Message, s.logicalClock)
	return &pb.Empty{}, nil
}

// RecieveMessages - stream messages to connected clients
func (s *server) RecieveMessages(_ *pb.Empty, stream pb.ChittyChat_RecieveMessagesServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	log.Println("client connected")

	for _, msg := range s.messages {
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}
	s.clients = append(s.clients, stream)

	for {
		time.Sleep(time.Second)
	}
}

// // JoinChat - client joins and gets a welcome message with a timestamp
// func (s *server) JoinChat(ctx context.Context, p *pb.Participant) (*pb.JoinLeaveResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	// Increment logical clock on join
// 	s.logicalClock++
// 	s.participants[p.Name] = s.logicalClock

// 	message := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", p.Name, s.logicalClock)
// 	log.Println(message)

// 	return &pb.JoinLeaveResponse{
// 		Message:   message,
// 		Timestamp: s.logicalClock,
// 	}, nil
// }

// // LeaveChat - client leaves and gets a leave message
// func (s *server) LeaveChat(ctx context.Context, p *pb.Participant) (*pb.JoinLeaveResponse, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	s.logicalClock++
// 	delete(s.participants, p.Name)

// 	message := fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time %d", p.Name, s.logicalClock)
// 	log.Println(message)

// 	return &pb.JoinLeaveResponse{
// 		Message:   message,
// 		Timestamp: s.logicalClock,
// 	}, nil
// }
