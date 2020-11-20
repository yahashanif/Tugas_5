package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	cm "Framework_Hanif_Aulia_Sabri/git/customer/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) FastPayHandler(ctx context.Context, req cm.FastPayRequest) (res cm.FastPayResponse) {
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

	var transId = req.TransId
	var fastPayResponse cm.FastPayResponse

	sql := `SELECT
				TransId,
				IFNULL(Merchant_id,''),
				IFNULL(Merchant,'') Merchant,
				IFNULL(Code,'') Code,
				IFNULL(Name,'') Name,
				IFNULL(Response_Code,'') Response_Code,
				IFNULL(Response_Desc,'') Response_Desc
			FROM trans WHERE TransId = ?`

	result, err := db.Query(sql, transId)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&fastPayResponse.TransId, &fastPayResponse.Merchant_id, &fastPayResponse.Merchant_id,
			&fastPayResponse.Code, &fastPayResponse.Name, &fastPayResponse.Response_Code, &fastPayResponse.Response_Desc)

		if err != nil {
			panic(err.Error())
		}

	}

	if fastPayResponse.TransId != "" {
		res = fastPayResponse
		res.Response = req.Request
		res.Response_Code = strconv.Itoa(http.StatusOK)
		res.Response_Desc = "Sukses ambil data"
	} else {
		res.Response_Code = strconv.Itoa(http.StatusNotFound)
		res.Response_Desc = "Gagal ambil data"
	}

	return

}
