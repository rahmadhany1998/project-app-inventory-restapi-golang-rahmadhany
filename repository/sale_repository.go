package repository

import (
	"database/sql"
	"project-app-inventory-restapi-golang-rahmadhany/model"
)

type SaleRepository interface {
	GetAll() ([]model.Sale, error)
	GetByID(id int) (*model.Sale, error)
	Create(sale model.Sale) (int, error)
	GetReportSummaryByDate(start, end string) (*model.Report, error)
}

type saleRepository struct {
	db *sql.DB
}

func NewSaleRepository(db *sql.DB) SaleRepository {
	return &saleRepository{db}
}

func (r *saleRepository) GetAll() ([]model.Sale, error) {
	rows, err := r.db.Query("SELECT id, product_id, item_sold, total_bill, date_sale FROM sales ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []model.Sale
	for rows.Next() {
		var u model.Sale
		if err := rows.Scan(&u.ID, &u.ProductID, &u.ItemSold, &u.TotalBill, &u.DateSale); err != nil {
			return nil, err
		}
		sales = append(sales, u)
	}
	return sales, nil
}

func (r *saleRepository) GetByID(id int) (*model.Sale, error) {
	row := r.db.QueryRow("SELECT id, product_id, item_sold, total_bill, date_sale FROM sales WHERE id = $1", id)
	var u model.Sale
	if err := row.Scan(&u.ID, &u.ProductID, &u.ItemSold, &u.TotalBill, &u.DateSale); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *saleRepository) Create(sale model.Sale) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO sales(product_id, item_sold, total_bill, date_sale) VALUES($1, $2, $3, $4) RETURNING id",
		sale.ProductID, sale.ItemSold, sale.TotalBill, sale.DateSale).Scan(&id)
	return id, err
}

func (r *saleRepository) GetReportSummaryByDate(start, end string) (*model.Report, error) {
	row := r.db.QueryRow("SELECT (SELECT COUNT(*) FROM products) AS total_product, COALESCE(SUM(s.item_sold), 0) AS total_sale, COALESCE(SUM((p.selling_price - p.retail_price) * s.item_sold), 0) AS total_revenue FROM sales s JOIN products p ON s.product_id = p.id WHERE s.date_sale BETWEEN $1 AND $2", start, end)
	var u model.Report
	if err := row.Scan(&u.TotalProduct, &u.TotalSale, &u.TotalRevenue); err != nil {
		return nil, err
	}
	return &u, nil
}
