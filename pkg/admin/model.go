package admin

import (
	"candidate_service/pkg/commons"
)

// heart beat status of the service
type Heartbeat struct {
	Status string `json: "status"`
}

// health info
type Info struct {
	CPU                interface{}
	Memory             interface{}
	Host               interface{}
	Disk               interface{}
	DBConnectionHealth string
	CacheHealth        string
	KinesisHealth      string
}

// health of the service
type Health struct {
	CPU                interface{}
	Memory             interface{}
	Host               interface{}
	Disk               interface{}
	DBConnectionHealth string
	CacheHealth        string
	KinesisHealth      string
}

// heart beat response of the service
type HeartbeatResponse struct {
	Data  Heartbeat        `json: "heartbeat"`
	Error commons.APIError `json: "error"`
}
