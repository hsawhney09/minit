package scheduler

import (
	"github.com/hsawhney09/minit/mesos/framework"
	"github.com/hsawhney09/minit/mesos/proto"
)

type Config struct {
	Framework     framework.Framework
	FrameworkInfo *mesos.FrameworkInfo
	Master        string
}
