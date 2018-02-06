package unimatrix_sdk

func Read( url string ) ( UnimatrixObject, error ) {
	response, err := Request( url, "GET" )

	if err != nil {
		return UnimatrixObject{}, err
	}

	return response, nil
}