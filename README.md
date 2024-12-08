
# BSDISYS1KU-En-GO-gruppe
**Run the program**

Open terminal session and cd to project
```
project
│   README.md
│   server.go
│   client.go
│   go.mod
│   go.sum
│   chitty_chat.proto
└───chittychatpb
│   │   chitty_chat.pb.go
│   │   chitty_chat_grpc.pb.go
└───└─────────────────────────
```

Run the clientside part of the program
```
go run client.go
```
Run the serverside part of the program
```
go run server.go
```

| System Requirements |  |
|--|--|
| R1 | x |
| R2 | x |
| R3 | x |
| R4 | x |
| R5 | x |
| R6 | x |
| R7 | x |
| R8 | x |

-   R1: Chitty-Chat is a distributed service, that enables its clients to chat. The service is using gRPC for communication. You have to design the API, including gRPC methods and data types.
-   R2: Clients in Chitty-Chat can **Publish** a valid chat message at any time they wish. A valid message is a string of UTF-8 encoded text with a maximum length of 128 characters. A client **publishes** a message by making a gRPC call to Chitty-Chat.
-   R3: The Chitty-Chat service has to **broadcast** every published message, together with the current  **logical timestamp**, to all participants in the system, by using gRPC. It is an implementation decision left to the students, whether a Vector Clock or a Lamport timestamp is sent.
-   R4: When a client receives a broadcasted message, it has to write the message and the current logical timestamp to the log
-   R5: Chat clients can join at any time.
-   R6: A "Participant X joined Chitty-Chat at Lamport time L" message is broadcast to all Participants when client X joins, including the new Participant.
-   R7: Chat clients can drop out at any time.
-   R8: A "Participant X left Chitty-Chat at Lamport time L" message is broadcast to all remaining Participants when Participant X leaves.
