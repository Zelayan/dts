package main

import (
	"context"
	"github.com/Zelayan/dts/internal/service"
	pb "github.com/Zelayan/dts/proto-gen/v1/dts"

)

func main() {

	collertorAddress := ":3001"
	ctx := context.Background()

	ch := make(chan *pb.Span, 10)

	
	// TODO: server
	sender, err := service.NewSender(ctx, collertorAddress, ch)
	if err != nil {
		panic(err)
	}
	go sender.BatchSendG()
	sender.Run()
}
