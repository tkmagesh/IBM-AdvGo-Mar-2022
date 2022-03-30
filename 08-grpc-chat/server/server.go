package main

import (
	"context"
	"grpc-chat/proto"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Connection struct {
	id       string
	active   bool
	userName string
	error    chan error
	stream   proto.Broadcast_SignInServer
}

type server struct {
	proto.UnimplementedBroadcastServer
	Connections []*Connection
}

func (s *server) SignIn(con *proto.Connect, stream proto.Broadcast_SignInServer) error {
	user := con.GetUser()
	conn := &Connection{
		id:       user.GetId(),
		userName: user.GetName(),
		active:   true,
		stream:   stream,
		error:    make(chan error),
	}
	s.Connections = append(s.Connections, conn)
	return <-conn.error
}

func (s *server) BroadcastMessage(ctx context.Context, message *proto.Message) (*proto.Close, error) {

	wg := &sync.WaitGroup{}
	for _, conn := range s.Connections {
		wg.Add(1)
		go func(msg *proto.Message, conn *Connection) {
			defer wg.Done()
			if conn.active {
				if err := conn.stream.Send(msg); err != nil {
					conn.active = false
					conn.error <- err
				}
			}
		}(message, conn)
	}
	wg.Wait()
	return &proto.Close{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	s := &server{}
	proto.RegisterBroadcastServer(grpcServer, s)
	grpcServer.Serve(listener)
}
