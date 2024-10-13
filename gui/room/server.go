package room

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/rzaf/p2p-chat/gui/config"
	"github.com/rzaf/p2p-chat/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.ChatServiceServer
}

var (
	lis        net.Listener
	grpcServer *grpc.Server
)

// Get IP from GRPC context
func getIP(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// fmt.Print("%v\n", md)
		port := md.Get("port")[0]
		p, _ := peer.FromContext(ctx)
		ip := strings.Split(p.Addr.String(), ":")[0]
		return ip + ":" + port
	}
	return ""
}

func (s *server) Message(c context.Context, m *pb.Text) (*pb.Empty, error) {
	ip := ""
	port := ""
	if md, ok := metadata.FromIncomingContext(c); ok {
		fmt.Print("%v\n", md)
		port = md.Get("port")[0]
		p, _ := peer.FromContext(c)
		ip = strings.Split(p.Addr.String(), ":")[0]
	}
	fmt.Printf("message recieved from `ip:%s Port:%s` \n", ip, port)
	if m != nil {
		AddMessage(m, ip, port)
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
