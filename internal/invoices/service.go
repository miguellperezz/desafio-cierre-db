package invoices

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	BulkCreate(invoices *[]domain.Invoices) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) ReadAll() ([]*domain.Invoices, error) {
	return s.r.ReadAll()
}

func (s *service) BulkCreate(invoices *[]domain.Invoices) error {
	for _, invoice := range *invoices {
		_, err := s.r.Create(&invoice)
		if err != nil {
			return err
		}
	}
	return nil
}
