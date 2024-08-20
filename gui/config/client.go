package config

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/rzaf/p2p-chat/pb"
	"io"
	"time"

	"golang.org/x/crypto/nacl/secretbox"
	"google.golang.org/grpc"
)

func SendMessageTo(secret []byte, uuid string, userUuid string, m string, addr string, username string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connection failed:\n")
		return err
	}
	defer conn.Close()
	fmt.Printf("connected to server \n")
	c := pb.NewChatServiceClient(conn)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], []byte(m), &nonce, (*[32]byte)(secret))
	_, err = c.Message(context.Background(), &pb.Text{Message: encrypted, UnixMilli: time.Now().UnixMilli(), RoomUuid: uuid, UserUuid: userUuid, Username: username})
	if err != nil {
		fmt.Printf("failed to send message:\n")
		return err
	}
	return nil
}
