package task

import (
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

type Task struct {
	ID           uuid.UUID
	Name         string
	State        State
	Image        string
	Memory       int
	Disk         int
	ExposedPorts nat.PortSet
}
