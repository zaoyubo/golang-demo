package grpc

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/go-chassis/v2/core/server"
	pb "golang-demo/protocol/grpc/go"
)

type Api struct {
	helloService pb.HelloServiceServer
}

func NewGrpcAPi(service pb.HelloServiceServer) *Api {
	return &Api{
		helloService: service,
	}
}

func (api *Api) Register() {
	chassis.RegisterSchema("grpc", api.helloService, server.WithRPCServiceDesc(&pb.HelloService_serviceDesc))
}
