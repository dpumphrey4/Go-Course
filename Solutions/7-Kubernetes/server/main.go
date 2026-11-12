package main

import (
	"context"
	"log"
	"net"
	"sync/atomic"

	pb "server/gophertalk"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGopherChatServer

	chirpCount uint64
	gnawCount  uint64
}

func (s *server) SendChirp(
	ctx context.Context,
	req *pb.Chirp,
) (*pb.Chitter, error) {

	total := atomic.AddUint64(&s.chirpCount, 1)

	log.Printf(
		"[CHIRP #%d] sender=%s message=%q",
		total,
		req.Sender,
		req.Message,
	)

	return &pb.Chitter{
		Responder: "TunnelMaster",
		Message:   "Heard your chirp: " + req.Message,
	}, nil
}

func (s *server) SendNibble(
	ctx context.Context,
	req *pb.Nibble,
) (*pb.GnawAck, error) {

	total := atomic.AddUint64(&s.gnawCount, 1)

	log.Printf(
		"[GNAW #%d] sender=%s content=%q",
		total,
		req.Sender,
		req.Content,
	)

	return &pb.GnawAck{
		Responder: "GnawKeeper",
		Received:  true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGopherChatServer(
		grpcServer,
		&server{},
	)

	log.Println("🐹 GopherChat listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
