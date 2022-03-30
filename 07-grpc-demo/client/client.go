package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	/* certFile := "ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalln(sslErr)
	} */

	insecureOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", insecureOptions)

	/* secureOptions := grpc.WithTransportCredentials(creds)
	clientConn, err := grpc.Dial("localhost:50051", secureOptions) */

	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	doRequestResponse(ctx, client)
	//doServerStreaming(ctx, client)
	//doClientStreaming(ctx, client)
	//doBidirectionalStreaming(ctx, client)
	//doRequestResponseWithTimeout(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Result = ", res.GetResult())
}

func doRequestResponseWithTimeout(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	res, err := client.Add(timeoutCtx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timeout error")
			} else {
				log.Fatalln(err)
			}
		}
		log.Fatalln(err)
	}
	fmt.Println("Add Result = ", res.GetResult())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("Thats all folks")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No = %d\n", res.GetPrimeNo())
	}
}

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	nos := []int32{3, 1, 4, 2, 5, 9, 6, 8, 7}
	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		fmt.Println("Avg Request : Sending ", no)
		time.Sleep(500 * time.Millisecond)
		req := &proto.AverageRequest{
			No: no,
		}
		clientStream.Send(req)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Response from server : Count = %d, Average = %d\n", res.GetCount(), res.GetAverage())
}

func doBidirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	personNames := []proto.PersonName{
		proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		proto.PersonName{FirstName: "Ganesh", LastName: "Easwaran"},
		proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}

	clientStream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	doneCh := make(chan struct{})
	go func() {
		for {
			res, err := clientStream.Recv()
			if err == io.EOF {
				close(doneCh)
				break
			}
			if err != nil {
				close(doneCh)
				log.Fatalln(err)
			}
			fmt.Printf("Response from Server : %s\n", res.GetGreetMessage())
		}
	}()

	for _, personName := range personNames {
		time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		fmt.Printf("Sending : %s, %s\n", personName.GetFirstName(), personName.GetLastName())
		clientStream.Send(req)
	}
	clientStream.CloseSend()
	<-doneCh
}
