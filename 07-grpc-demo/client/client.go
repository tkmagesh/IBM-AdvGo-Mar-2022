package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//doRequestResponse(ctx, client)
	//doServerStreaming(ctx, client)
	doClientStreaming(ctx, client)
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
