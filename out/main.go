package main

import (
	pb "github.com/some-repo/examples/helloworld"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterHelloWorldServiceService(s.Service("helloworld.HelloWorldService"), &helloWorldServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
