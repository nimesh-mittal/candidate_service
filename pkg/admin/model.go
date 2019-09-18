package admin

import (
	"candidate_service/pkg/commons"
)

type Heartbeat struct {
	Status string `json: "status"`
}
type Info struct {
	CPU                interface{}
	Memory             interface{}
	Host               interface{}
	Disk               interface{}
	DBConnectionHealth string
	CacheHealth        string
	KinesisHealth      string
}

type Health struct {
	CPU                interface{}
	Memory             interface{}
	Host               interface{}
	Disk               interface{}
	DBConnectionHealth string
	CacheHealth        string
	KinesisHealth      string
}

type HeartbeatResponse struct {
	Data  Heartbeat        `json: "heartbeat"`
	Error commons.APIError `json: "error"`
}
