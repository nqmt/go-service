package soap

import (
	"github.com/jmoiron/sqlx"
)

//go:generate mockery -name=ISoapRepo
type ISoapRepo interface {
	GetSoap() ([]*Soap, error)
}

type SoapRepo struct {
	soap *sqlx.DB
}

func NewSoapRepo(sql *sqlx.DB) *SoapRepo {
	return &SoapRepo{
		soap: sql,
	}
}

func (p *SoapRepo) GetSoap() ([]*Soap, error) {
	var output []*Soap

	err := p.soap.Select(&output, "select * from PERSONS2")
	if err != nil {
		return nil, err
	}

	return output, nil
}