package room

import (
	"context"
	"fmt"
	"github.com/rzaf/p2p-chat/gui/config"
	"github.com/rzaf/p2p-chat/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.ChatServiceServer
}

var (
	lis        net.Listener
	grpcServer *grpc.Server
)

func (s *server) Message(c context.Context, m *pb.Text) (*pb.Empty, error) {
	fmt.Printf("message %v recieved \n", m)
	if m != nil {
		AddMessage(m)
	}
	return &pb.Empty{}, nil
}

func StartServer() {
	addr := config.Addr + ":" + config.Port
	var err error
	lis, err = net.Listen("tcp4", addr)
	if err != nil {
		fmt.Printf("Failed to listen at:%v", addr)
		config.ShowError(err)
		return
	}
	fmt.Printf("serving at: %v\n", addr)
	grpcServer = grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})
	if err = grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to start server")
		config.ShowError(err)
	}
}

func StopServer() {
	if grpcServer != nil {
		grpcServer.Stop()
	}
	if lis != nil {
		lis.Close()
	}
}
