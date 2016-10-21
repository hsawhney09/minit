package main

import (
	"fmt"

	"github.com/hsawhney09/minit/mesos/proto"
	"github.com/hsawhney09/minit/mesos/scheduler"
	"github.com/gogo/protobuf/proto"
)

func main() {
	fmt.Println("heelo")


	sc := scheduler.Config{
		Framework: &MyFramework{},
		FrameworkInfo: &mesos.FrameworkInfo{
			User: proto.String("harpreet"),
			Name: proto.String("MyFramework"),
		},
		Master: "192.168.23.126:5050",
	}

	s := scheduler.NewMesosScheduler(sc)
	if err :=  s.Start(); err != nil {
		fmt.Println("dadada")
	}
}
