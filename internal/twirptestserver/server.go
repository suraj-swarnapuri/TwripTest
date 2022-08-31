package twirptestserver

import (
	"context"
	"fmt"
	pb "twirptest/rpc/twirptest"
)

type HelloWorldServer struct{
	playerCount int
	
}


func (s *HelloWorldServer) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	s.playerCount++
	if s.playerCount > 5 {
		return nil, fmt.Errorf("too many players")
	}
	msg := fmt.Sprintf("playerCount: %d", s.playerCount)
	return &pb.HelloResp{Text: "Hello " + in.Subject + " "+ msg}, nil
}