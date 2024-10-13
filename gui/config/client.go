package config

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/rzaf/p2p-chat/pb"

	"golang.org/x/crypto/nacl/secretbox"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func SendMessageTo(secret []byte, uuid string, userUuid string, m string, ip string, port string, username string) error {
	// conn, err := grpc.NewClient(addr,
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
	// 		dst, err := net.ResolveTCPAddr("tcp", addr)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		src := &net.TCPAddr{
	// 			IP:   net.ParseIP("127.54.0.1"),
	// 			Port: 1212111,
	// 		}
	// 		return net.DialTCP("tcp", src, dst)
	// 	}),
	// )
	conn, err := grpc.NewClient(ip+":"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("connection failed:\n")
		return err
	}
	defer conn.Close()
	// fmt.Printf("connected to server \n")
	c := pb.NewChatServiceClient(conn)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], []byte(m), &nonce, (*[32]byte)(secret))
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("port", Port))
	_, err = c.Message(ctx, &pb.Text{Message: encrypted, UnixMilli: time.Now().UnixMilli(), RoomUuid: uuid, UserUuid: userUuid, Username: username})
	if err != nil {
		fmt.Printf("failed to send message:\n")
		return err
	}
	return nil
}
