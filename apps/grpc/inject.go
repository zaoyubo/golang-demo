package grpc

import (
	model "golang-demo/internal/model/hello"
	service "golang-demo/internal/service/hello"
)

func Init() *Api {
	dao := model.NewHelloDao()
	helloService := service.NewHelloService(dao)
	api := NewGrpcAPi(helloService)
	return api
}
