package scheduler

import (
	"github.com/hsawhney09/minit/mesos/framework"
	"github.com/hsawhney09/minit/mesos/proto"
)

type MesosScheduler struct {
	framework     framework.Framework
	frameworkInfo *mesos.FrameworkInfo
	master        string
}

func NewMesosScheduler(config Config) *MesosScheduler {
	return &MesosScheduler{
		master:        config.Master,
		frameworkInfo: config.FrameworkInfo,
		framework:     config.Framework,
	}
}

// Scheduler interface methods
func (p *MesosScheduler) Start() error {

	return nil
}

func (p *MesosScheduler) Stop() error {
	return nil
}
