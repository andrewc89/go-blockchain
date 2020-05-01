package blockchain

import (
	"fmt"
	"strings"
)

type QueryParameters map[string]string

func (qps *QueryParameters) String() string {
	s := make([]string, 0, len(*qps))
	for k, v := range *qps {
		qp := fmt.Sprintf("%s=%s", k, v)
		s = append(s, qp)
	}
	return strings.Join(s, "&")
}

type Endpoint struct {
	BaseUrl         string
	QueryParameters QueryParameters
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s?%s", endpoint.BaseUrl, endpoint.QueryParameters.String())
}
