package main

import (
	"context"
	"log"
	"net"

	pb "git.der-waldemar.net/shipping-application/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// ConsignmentRepository - Dummy repository
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create a new Consignment and add it to the ConsignmentRepository
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}

// Service - Implements all methods to satisfy the service defined in the protobuf.
type service struct {
	repo Repository
}

// CreateConsignment - Implementation of the service as defined in protobuf
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// Save consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return 'Response' message
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &ConsignmentRepository{}

	// Setup gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on: %v", err)
	}
	server := grpc.NewServer()

	// Register our service with the gRPC server
	pb.RegisterShippingServiceServer(server, &service{repo})

	// Register reflection service on gRPC server
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
