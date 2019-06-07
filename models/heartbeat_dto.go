package models

type HeartbeatResponse struct {
	Data  Heartbeat `json: "heartbeat"`
	Error APIError  `json: "error"`
}
