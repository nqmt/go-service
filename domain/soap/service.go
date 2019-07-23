package soap

type IService interface {
	GetSoap() ([]*Soap, error)
}

type Service struct {
	soapRepo ISoapRepo
}

func NewService(soapRepo ISoapRepo) *Service {
	return &Service{
		soapRepo: soapRepo,
	}
}

func (s *Service) GetSoap() ([]*Soap, error) {
	return s.soapRepo.GetSoap()
}
