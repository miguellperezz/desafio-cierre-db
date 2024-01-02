package invoices

import ("github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
		
)

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	BulkCreate(invoices *[]domain.Invoices) error
	UpdateTotalInvoices() error
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

func (s *service) UpdateTotalInvoices() error {
	invoices, err := s.r.ReadAll()
	if err != nil {
		return err
	}
	for _, invoice := range invoices {
		sales, err := s.r.ReadSalesByInvoiceID(invoice.Id)
		if err != nil {
			return err
		}
		var total float64
		for _, sale := range sales {
			
			//productRepo := NewRepository(s.r)
			//productService := NewService(s.r)
			//product, err :=  GetProduct(sale.ProductId)
			product, err := s.r.GetProduct(sale.ProductId)
			if err != nil {
				return err
			}

			total += product.Price * float64(sale.Quantity)
		}
		invoice.Total = total
		err = s.r.Update(invoice)
		if err != nil {
			return err
		}
	}
	return nil
}
