package user_management_client;

import (
	"context"
	"log"
	"time"

	pb "google.golang.org/grpc"
)
const (
	address = "localhost:50051"
)

fun main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]int32)
	new_users["Alice"]=43
	new_users["Bob"]=30
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name:name, Age:age})
		if err != nil {
			log.Fatalf("Could not create user: %v", err)
		}
		log.Printf(`User Details:
		NAME: %s
		AGE: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}