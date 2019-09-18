package commons

import (
	"encoding/json"
)

func ToResponse(payload interface{}, code string, error error) *Response {
	if error == nil {
		return &Response{payload, &APIError{code, EMPTY, error}}
	}
	return &Response{payload, &APIError{code, error.Error(), error}}
}

func MakeResp(payload interface{}, code string, err error) []byte {
	var res Response
	if err != nil {
		res = *ToResponse(nil, code, err)
	} else {
		res = *ToResponse(payload, code, nil)
	}
	data, err1 := json.Marshal(res)

	if err1 != nil {
		res = *ToResponse(nil, code, err1)
		data, _ = json.Marshal(res)
	}

	return data
}
