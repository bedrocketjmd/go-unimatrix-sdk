package unimatrix

type Response struct {
	parser *Parser
}

func NewResponse(body []byte) (*Response, error) {
	parser, error := NewParser(body)

	if error != nil {
		return nil, error
	}

	return &Response{parser: parser}, nil
}

func (response *Response) Name() (string, error) {
	return response.parser.Name, nil
}

func (response *Response) TypeName() (string, error) {
	return response.parser.TypeName, nil
}

func (response *Response) Ids() ([]string, error) {
	return response.parser.Keys, nil
}

func (response *Response) Resources() ([]Resource, error) {
	return response.parser.Resources, nil
}

func (response *Response) Count() (int, error) {
	return response.parser.Count, nil
}

func (response *Response) UnlimitedCount() (int, error) {
	return response.parser.UnlimitedCount, nil
}

func (response *Response) Offset() (int, error) {
	return response.parser.Offset, nil
}
