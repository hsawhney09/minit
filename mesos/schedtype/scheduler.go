package schedtype


type Scheduler interface {
	Start() error
	Stop()
}

