package main

import (
	"EndkaGo/calculatorpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func PrimeComposition(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()
	req := &calculatorpb.NumberRequest{
		Number: 120,
	}

	stream, err := c.PrimeNumberDecomposition(ctx, req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from PrimeComposition RPC %v", err)
		}
		log.Printf("response from GreetManyTimes:%v \n", res.GetResult())
	}

}

func getAverage(c calculatorpb.CalculatorServiceClient) {

	requests := []*calculatorpb.NumbersRequest{
		{
			Numbers: 1,
		},
		{
			Numbers: 2,
		},
		{
			Numbers: 3,
		},
		{
			Numbers: 4,
		},
	}

	ctx := context.Background()
	stream, err := c.ComputeAverage(ctx)
	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending : %v\n", req)
		stream.Send(req)
		time.Sleep(345 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverage Response: %v\n", res)
}

func main() {
	fmt.Println("Hi baby im u client")

	conn, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(" cant connect to 6666: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	PrimeComposition(c)
	getAverage(c)
}
