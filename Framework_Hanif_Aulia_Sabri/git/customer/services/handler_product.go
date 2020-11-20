package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "Framework_Hanif_Aulia_Sabri/git/customer/common"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB
// var err error

func (PaymentService) ProductHandler(ctx context.Context, req cm.Products) (res cm.Products) {
	var db *sql.DB
	var err error
	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.ProductID = req.ProductID

	var product cm.Products

	sql := `SELECT
				ProductID,
				IFNULL(ProductName,''),
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitsInStock
			FROM products WHERE ProductID = ?`

	result, err := db.Query(sql, req.ProductID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&product.ProductID, &product.ProductName, &product.SupplierID,
			&product.CategoryID, &product.QuantityPerUnit, &product.UnitPrice, &product.UnitsInStock)

		if err != nil {
			panic(err.Error())
		}

	}

	res = product

	return
}
