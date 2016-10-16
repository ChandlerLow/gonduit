package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

func (c *Conn) RemarkupProcess(
	req requests.RemarkupProcessRequest,
) (*responses.RemarkupProcessResponse, error) {
	var res responses.RemarkupProcessResponse

	if err := c.Call("remarkup.process", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
