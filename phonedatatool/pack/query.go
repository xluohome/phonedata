package pack

import "github.com/xluohome/phonedata/phonedatatool"

type Querier struct {
}

func NewQuerier() *Querier {
	return &Querier{}
}

func (q *Querier) Query(phoneDataBuf []byte, number string) (*phonedatatool.QueryResult, error) {
	return nil, nil
}
