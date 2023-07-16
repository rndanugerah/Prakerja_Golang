package employees

import (
	"net/http"
	"prakerja/db"
	"prakerja/models"

	validator "github.com/go-playground/validator/v10"
)

func StoreEmployees(nama string, email string, telepon string) (models.Response, error) {
	var res models.Response

	v := validator.New()

	emp := Employee{
		Nama:    nama,
		Email:   email,
		Telepon: telepon,
	}

	err := v.Struct(emp)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT pegawai (nama, email, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, email, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}
