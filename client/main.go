package main

import (
	"context"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	pb "traefik-grpc-tls/server/proto"
)

const (
	address     = "frontend.local:4443"
	defaultName = "world"
)

func main() {
	frontendCert, _ := ioutil.ReadFile("../traefik/cert/frontend.local.cert")

	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(frontendCert)

	credsClient := credentials.NewClientTLSFromCert(roots, "")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credsClient))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}