package main

import (
	"EndkaGo/calculatorpb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)



type Server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func findAverage(arr []int64) (avg float64) {
	sum := 0.0
	ln := len(arr)
	for i := 0; i < ln; i++ {
		sum += float64(arr[i])
	}
	avg = sum / float64(ln)
	return
}
func Factors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}
	return
}

//Structure
func (s *Server) PrimeNumberDecomposition(req *calculatorpb.NumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v \n", req)
	number := req.GetNumber()
	n := int(number)
	arr := Factors(n)
	for i := 0; i < len(arr); i++ {
		res := &calculatorpb.CalculatorResponse{Result: fmt.Sprintf("Baby : %v\n", arr[i])}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending greet many times responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *Server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("AverageClient function was invoked with a streaming request\n")
	var result float64
	var arr []int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			result = findAverage(arr)
			return stream.SendAndClose(&calculatorpb.AverageResponse{Result: result})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		arr = append(arr, req.GetNumbers())
	}
}

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:6666")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
