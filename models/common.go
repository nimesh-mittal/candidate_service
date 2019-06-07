package models

type Response struct {
	Data  interface{}
	Error *APIError
}

type APIError struct {
	Code     string
	Message  string
	Metadata interface{}
}

type FlowCtx struct {
	TrackingID string
}
