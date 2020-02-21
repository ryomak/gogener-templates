package main

import (
	"context"
	"fmt"
	"os"

	"[[.ModName]]/src/interface/rpc"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}
	defer conn.Close()
	client := rpc.NewRoomHandlerClient(conn)
	coordinate := &rpc.Coordinate{Latitude: 13.22, Longitude: 12.22}
	user := &rpc.User{Name: "test", Coordinate: coordinate}
	req := &rpc.RoomRequest{Name: "test", User: user}
	room, err := client.GetRoom(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "room: %v\n", room)
		fmt.Fprintf(os.Stderr, "room: %v\n", err)
		return
	}
	fmt.Printf("%v", room)
}
