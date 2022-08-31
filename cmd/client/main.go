package main

import (
	"context"
	"net/http"

	pb "twirptest/rpc/twirptest"
)

func main(){
	client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})
	hello, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "world"})
	if err != nil {
		panic(err)
	}
	println(hello.Text)

}