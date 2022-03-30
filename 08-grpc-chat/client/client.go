package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"grpc-chat/proto"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var client proto.BroadcastClient

func main() {
	timestamp := time.Now()
	name := flag.String("N", "Anon", "The name of the user")
	flag.Parse()
	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client = proto.NewBroadcastClient(conn)
	user := &proto.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}
	connect(user)

	//get input from stdin and send to the server
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := &proto.Message{
				Id:      user.Id,
				Content: scanner.Text(),
			}
			if _, err := client.BroadcastMessage(context.Background(), msg); err != nil {
				log.Fatalln(err)
				break
			}
		}
	}()
	wg.Wait()

}

func connect(user *proto.User) {
	stream, err := client.SignIn(context.Background(), &proto.Connect{User: user, Active: true})
	if err != nil {
		log.Fatalln(err)
	}
	go func(strm proto.Broadcast_SignInClient) {
		for {
			msg, err := strm.Recv()
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(msg.GetContent())
		}
	}(stream)
}
