package query

import "github.com/xluohome/phonedata/phonedatatool"

type Querier struct{}

func NewQuerier(phoneDataFilePath string) phonedatatool.Querier {
	return new(Querier)
}

func (q *Querier) QueryNumber(number string) (*phonedatatool.QueryResult, error) {
	return nil, nil
}
