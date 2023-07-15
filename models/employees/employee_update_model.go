package employees

import (
	"net/http"
	"prakerja/db"
	"prakerja/models"
)

func UpdateEmployees(id int, nama string, email string, telepon string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama = ?, email = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, email, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
