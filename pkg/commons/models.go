package commons

type Response struct {
	Data  interface{}
	Error *APIError
}

type APIError struct {
	Code     string
	Message  string
	Metadata interface{}
}

type FlowContext struct {
	TrackingID string
}
