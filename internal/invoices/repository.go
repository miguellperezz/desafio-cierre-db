package invoices

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	ReadSalesByInvoiceID(id int) ([]*domain.Sales, error)
	Update(invoices *domain.Invoices) error
	GetProduct(id int) (*domain.Product, error)

}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) ReadSalesByInvoiceID(id int) ([]*domain.Sales, error){
	query := `SELECT id, invoice_id, product_id, quantity FROM sales WHERE invoice_id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sales := make([]*domain.Sales, 0)
	for rows.Next() {
		sale := domain.Sales{}
		err := rows.Scan(&sale.Id, &sale.InvoicesId, &sale.ProductId, &sale.Quantity)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}
	return sales, nil
}

func (r *repository) Update(invoices *domain.Invoices) error{
	query := "UPDATE invoices SET total = ? WHERE id = ? "
	_, err := r.db.Exec(query, &invoices.Total, &invoices.Id)
	if err != nil {
		return err
	}
	return nil
	
}

func (r *repository) GetProduct(id int) (*domain.Product, error){
	query := "SELECT id, price FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)
	product := domain.Product{}
	err := row.Scan(&product.Id, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}


