package commons

func OnError(payload interface{}, err error) (interface{}, error) {
	if err != nil{
		return nil, err
	}
	return payload, nil
}
