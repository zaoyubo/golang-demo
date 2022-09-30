package main

import (
	"log"

	chassis "github.com/go-chassis/go-chassis/v2"
	"golang-demo/apps/grpc"
	_ "golang-demo/initialize"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server_test/
func main() {
	api := grpc.Init()
	api.Register()
	if err := chassis.Init(); err != nil {
		log.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
