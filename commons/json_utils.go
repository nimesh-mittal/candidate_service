package commons

import (
	"candidate_service/models"
	"encoding/json"
)

func ToResponse(payload interface{}, code string, error error) *models.Response {
	if error == nil {
		return &models.Response{payload, &models.APIError{code, EMPTY, error}}
	}
	return &models.Response{payload, &models.APIError{code, error.Error(), error}}
}

func MakeResp(payload interface{}, code string, err error) []byte {
	var res models.Response
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
