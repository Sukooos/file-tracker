package main

import (
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "tracker-service/internal/proto/tracker"
)

type trackerServer struct {
    pb.UnimplementedTrackerServiceServer
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterTrackerServiceServer(s, &trackerServer{})

    fmt.Println("🚀 Tracker Service running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
