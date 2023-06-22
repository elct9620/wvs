package controller

import "time"

type System struct{}

func NewSystem() *System {
	return &System{}
}

type HeartbeatArgs struct {
	Timestamp time.Time `json:"timestamp"`
}

type HeartbeatReply struct {
	Timestamp time.Time `json:"timestamp"`
}

func (ctrl *System) Heartbeat(args *HeartbeatArgs, reply *HeartbeatReply) error {
	*reply = HeartbeatReply{
		Timestamp: time.Now(),
	}

	return nil
}
