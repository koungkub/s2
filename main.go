package main

import (
	"context"
	"fmt"

	pb "github.com/koungkub/s2/proto"
	micro "github.com/micro/go-micro/v2"
)

type Park struct {
}

func (g *Greeter) Get(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Name = "hello"
	return nil
}

func main() {

	s := micro.NewService(
		micro.Name("koung-server"),
	)

	s.Init()

	pb.RegisterGreeterHandler(s.Server(), new(Greeter))

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
