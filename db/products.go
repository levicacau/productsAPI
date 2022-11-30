package db

import (
	"database/sql"
	"productsAPI/models"
)

func (db Database) GetAllProducts() (*models.ProductList, error) {
	list := &models.ProductList{}
	rows, err := db.Conn.Query("SELECT codigobarras, precogondola1, descricao, unidade FROM webcompras ORDER BY descricao")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Codigobarras, &product.Precogondola1, &product.Descricao, &product.Unidade)
		if err != nil {
			return list, err
		}
		list.Products = append(list.Products, product)
	}
	return list, nil
}
func (db Database) AddProduct(product *models.Product) error {
	var codigobarras int
	var descricao string
	query := `INSERT INTO webcompras (codigobarras, precogondola1, descricao, unidade) VALUES ($1, $2, $3, $4) RETURNING codigobarras, descricao`
	err := db.Conn.QueryRow(query, product.Codigobarras, product.Precogondola1, product.Descricao, product.Unidade).Scan(&codigobarras, &descricao)
	if err != nil {
		return err
	}
	product.Codigobarras = codigobarras
	product.Descricao = descricao
	return nil
}
func (db Database) GetProductById(productId int) (models.Product, error) {
	product := models.Product{}
	query := `SELECT codigobarras, precogondola1, descricao, unidade FROM webcompras WHERE codigobarras = $1;`
	row := db.Conn.QueryRow(query, productId)
	switch err := row.Scan(&product.Codigobarras, &product.Precogondola1, &product.Descricao, &product.Unidade); err {
	case sql.ErrNoRows:
		return product, ErrNoMatch
	default:
		return product, err
	}
}
func (db Database) DeleteProduct(productId int) error {
	query := `DELETE FROM webcompras WHERE codigobarras = $1;`
	_, err := db.Conn.Exec(query, productId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
func (db Database) UpdateProduct(productId int, productData models.Product) (models.Product, error) {
	product := models.Product{}
	query := `UPDATE webcompras SET precogondola1=$1 descricao=$2, unidade=$3 WHERE codigobarras=$4 RETURNING codigobarras, precogondola1, descricao, unidade;`
	err := db.Conn.QueryRow(query, productData.Precogondola1, productData.Descricao, productData.Unidade, productId).Scan(&product.Codigobarras, &product.Precogondola1, &product.Descricao, &product.Unidade)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, ErrNoMatch
		}
		return product, err
	}
	return product, nil
}
