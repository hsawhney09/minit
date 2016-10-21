package main

import (
	"fmt"
	"github.com/hsawhney09/minit/mesos/proto"
	scheduler "github.com/hsawhney09/minit/mesos/schedtype"
)

type MyFramework struct {
}

func (p *MyFramework) Registered(scheduler.Scheduler, *mesos.FrameworkID, *mesos.MasterInfo) {
	fmt.Println("Registered")
}

func (p *MyFramework) Reregistered(scheduler.Scheduler, *mesos.MasterInfo) {
	fmt.Println("Reregistered")
}

func (p *MyFramework) Disconnected(scheduler.Scheduler) {
	fmt.Println("Disconnected")
}

func (p *MyFramework) ResourceOffers(scheduler.Scheduler, []*mesos.Offer) {
	fmt.Println("ResourceOffers")
}

func (p *MyFramework) OfferRescinded(scheduler.Scheduler, *mesos.OfferID) {
	fmt.Println("OfferRescinded")
}

func (p *MyFramework) StatusUpdate(scheduler.Scheduler, *mesos.TaskStatus) {
	fmt.Println("StatusUpdate")
}

func (p *MyFramework) FrameworkMessage(scheduler.Scheduler, *mesos.ExecutorID, *mesos.SlaveID, string) {
	fmt.Println("FrameworkMessage")
}

func (p *MyFramework) SlaveLost(scheduler.Scheduler, *mesos.SlaveID) {
	fmt.Println("SlaveLost")
}

func (p *MyFramework) ExecutorLost(scheduler.Scheduler, *mesos.ExecutorID, *mesos.SlaveID, int) {
	fmt.Println("ExecutorLost")
}

func (p *MyFramework) Error(scheduler.Scheduler, string) {
	fmt.Println("Error")
}
