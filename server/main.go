package main

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "traefik-grpc-tls/server/proto"
)

const (
	port = ":5300"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	backendCert, _ := ioutil.ReadFile("/go/src/traefik-grpc-tls/traefik/cert/backend.local.cert")
	backendKey, _ := ioutil.ReadFile("/go/src/traefik-grpc-tls/traefik/cert/backend.local.key")

	cert, err := tls.X509KeyPair(backendCert, backendKey)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
	}

	creds := credentials.NewServerTLSFromCert(&cert)

	serverOption := grpc.Creds(creds)

	s := grpc.NewServer(serverOption)

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

