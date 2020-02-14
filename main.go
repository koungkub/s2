package main

import (
	"context"
	"fmt"

	pb "github.com/koungkub/s2/proto"
	micro "github.com/micro/go-micro/v2"
)

type Park struct {
}

func (p *Park) Get(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Name = "hello"
	return nil
}

func main() {

	s := micro.NewService(
		micro.Name("koung-server2"),
	)

	s.Init()

	pb.RegisterParkHandler(s.Server(), new(Park))

	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
